package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"webchat_be/biz/dao"
	"webchat_be/biz/db/mysql"
	"webchat_be/biz/model/domain"
	"webchat_be/biz/model/errs"
	"webchat_be/biz/model/po"
	"webchat_be/biz/util/encode"
	"webchat_be/biz/util/random"
)

type LoginRequest struct {
	SessID   string
	Username string
	Password string
	IP       string
	Device   string
}

type LoginResponse struct {
	Username  string
	AccountId string
	Status    string
	Email     string
}

func AccountLogin(ctx context.Context, req *LoginRequest) (resp LoginResponse, loginResult errs.Error) {
	txErr := mysql.GetDbConn().Transaction(func(tx *gorm.DB) error {
		accountInfo, err := dao.NewAccountDao(tx).QueryByUsernameForUpdate(ctx, req.Username)
		if err != nil {
			return err
		}
		if accountInfo == nil {
			hlog.CtxInfof(ctx, "username not exists: %s", req.Username)
			loginResult = errs.AccountNotExistError
			_ = dao.NewLoginRecordDao(tx).Create(ctx, &po.LoginRecord{
				AccountID: "",
				Status:    domain.LoginRecordFailed,
				IP:        req.IP,
				Device:    req.Device,
			})
			return nil
		}
		resp.AccountId = accountInfo.AccountID
		resp.Status = accountInfo.Status
		resp.Email = accountInfo.Email
		resp.Username = accountInfo.Username

		if encode.EncodePassword(accountInfo.Salt, req.Password) != accountInfo.Password {
			hlog.CtxInfof(ctx, "password incorrect: %s", req.Username)
			loginResult = errs.PasswordIncorrect
			_ = dao.NewLoginRecordDao(tx).Create(ctx, &po.LoginRecord{
				AccountID: resp.AccountId,
				Status:    domain.LoginRecordFailed,
				IP:        req.IP,
				Device:    req.Device,
			})
			return nil
		}

		_ = dao.NewLoginRecordDao(tx).Create(ctx, &po.LoginRecord{
			AccountID: resp.AccountId,
			Status:    domain.LoginRecordSuccess,
			IP:        req.IP,
			Device:    req.Device,
		})

		return appendLoginAccount(ctx, resp.AccountId, req.SessID)
	})
	if txErr != nil {
		hlog.CtxErrorf(ctx, "txErr: %v", txErr)
		return LoginResponse{}, errs.ServerError
	}

	return
}

type PasswordUpdateRequest struct {
	AccountId   string
	Password    string
	PasswordNew string
}

func AccountUpdatePassword(ctx context.Context, req *PasswordUpdateRequest) (updateResult errs.Error) {
	txErr := mysql.GetDbConn().Transaction(func(tx *gorm.DB) error {
		accountDao := dao.NewAccountDao(tx)
		accountInfo, err := accountDao.QueryByAccountIdForUpdate(ctx, req.AccountId)
		if err != nil || accountInfo == nil {
			return err
		}
		if encode.EncodePassword(accountInfo.Salt, req.Password) != accountInfo.Password {
			hlog.CtxInfof(ctx, "password incorrect")
			updateResult = errs.PasswordIncorrect
			return nil
		}

		encodedPwd := encode.EncodePassword(accountInfo.Salt, req.PasswordNew)
		return accountDao.UpdatePassword(ctx, req.AccountId, encodedPwd, random.RandStr(32))
	})
	if txErr != nil {
		hlog.CtxErrorf(ctx, "txErr: %v", txErr)
		return errs.ServerError
	}

	return
}
