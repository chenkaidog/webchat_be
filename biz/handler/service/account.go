package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"regexp"
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
	accountId, bizErr := getAccountId(ctx, req.Username)
	if bizErr != nil {
		if errs.ErrorEqual(bizErr, errs.AccountNotExist) {
			_ = dao.NewLoginRecordDao().Create(ctx, &po.LoginRecord{
				AccountID: "",
				Status:    domain.LoginRecordFailed,
				IP:        req.IP,
				Device:    req.Device,
			})
		}
		return LoginResponse{}, bizErr
	}
	txErr := mysql.GetDbConn().Transaction(func(tx *gorm.DB) error {
		accountInfo, err := dao.NewAccountDao(tx).QueryByAccountIdForUpdate(ctx, accountId)
		if err != nil {
			return err
		}
		if accountInfo == nil {
			hlog.CtxInfof(ctx, "unexpected err, account_id not exist: %s", req.Username)
			loginResult = errs.ServerError
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

// 通过判断用户输入的是用户名还是邮箱，然后获取account_id进行登录
func getAccountId(ctx context.Context, username string) (string, errs.Error) {
	const emailPattern = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	matched, err := regexp.MatchString(emailPattern, username)
	if err != nil {
		hlog.CtxErrorf(ctx, "regexp err: %v", err)
		return "", errs.ServerError
	}

	var account *po.Account
	if matched {
		account, err = dao.NewAccountDao().QueryByEmail(ctx, username)
	} else {
		account, err = dao.NewAccountDao().QueryByUsername(ctx, username)
	}
	if err != nil {
		hlog.CtxErrorf(ctx, "query account err: %v", err)
		return "", errs.ServerError
	}
	if account == nil {
		return "", errs.AccountNotExist
	}
	return account.AccountID, nil
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

		salt := random.RandStr(32)
		encodedPwd := encode.EncodePassword(salt, req.PasswordNew)
		if err := accountDao.UpdatePassword(ctx, req.AccountId, encodedPwd, salt); err != nil {
			return err
		}

		return RemoveAllSession(ctx, req.AccountId)
	})
	if txErr != nil {
		hlog.CtxErrorf(ctx, "txErr: %v", txErr)
		return errs.ServerError
	}

	return
}
