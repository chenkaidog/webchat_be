package dao

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"webchat_be/biz/db/mysql"
	"webchat_be/biz/model/po"
)

type AccountDao struct {
	mysql.DbConn
}

func NewAccountDao(tx ...*gorm.DB) *AccountDao {
	return &AccountDao{
		DbConn: mysql.NewDbConn(tx...),
	}
}

func (dao *AccountDao) QueryByUsername(ctx context.Context, username string) (*po.Account, error) {
	var result *po.Account
	if err := dao.WithContext(ctx).
		Where("username = ?", username).
		Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by username errs: %v", err)
		return nil, err
	}

	return result, nil
}

func (dao *AccountDao) QueryByAccountId(ctx context.Context, accountId string) (*po.Account, error) {
	var result *po.Account
	if err := dao.WithContext(ctx).
		Where("account_id = ?", accountId).
		Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by account_id errs: %v", err)
		return nil, err
	}

	return result, nil
}
