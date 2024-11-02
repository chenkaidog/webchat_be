package dao

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"webchat_be/biz/db/mysql"
	"webchat_be/biz/model/po"
)

type AccountDao struct {
	conn *mysql.DbConn
}

func NewAccountDao(tx ...*gorm.DB) *AccountDao {
	return &AccountDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (dao *AccountDao) Create(ctx context.Context, accountInfo *po.Account) error {
	return dao.conn.WithContext(ctx).Create(accountInfo).Error
}

func (dao *AccountDao) QueryByUsernameForUpdate(ctx context.Context, username string) (*po.Account, error) {
	var result *po.Account
	if err := dao.conn.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("username", username).
		Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by username errs: %v", err)
		return nil, err
	}

	return result, nil
}

func (dao *AccountDao) QueryByAccountIdForUpdate(ctx context.Context, accountId string) (*po.Account, error) {
	var result *po.Account
	if err := dao.conn.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("account_id", accountId).
		Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by account_id errs: %v", err)
		return nil, err
	}

	return result, nil
}

func (dao *AccountDao) QueryByAccountId(ctx context.Context, accountId string) (*po.Account, error) {
	var result *po.Account
	if err := dao.conn.WithContext(ctx).
		Where("account_id", accountId).
		Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by account_id errs: %v", err)
		return nil, err
	}

	return result, nil
}

func (dao *AccountDao) UpdatePassword(ctx context.Context, accountId, password, salt string) error {
	err := dao.conn.WithContext(ctx).
		Where("account_id", accountId).
		Updates(map[string]string{
			"password": password,
			"salt":     salt,
		}).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "update account errs: %v", err)
		return err
	}
	return nil
}
