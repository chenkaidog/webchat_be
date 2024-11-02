package dao

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"webchat_be/biz/db/mysql"
	"webchat_be/biz/model/po"
)

type ModelDao struct {
	conn *mysql.DbConn
}

func NewModelDao(tx ...*gorm.DB) *ModelDao {
	return &ModelDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (dao *ModelDao) QueryByModelId(ctx context.Context, modelId string) (*po.Model, error) {
	var result *po.Model
	err := dao.conn.WithContext(ctx).
		Where("model_id = ?", modelId).
		Take(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by model id err: %v", err)
		return nil, err
	}

	return result, nil
}

func (dao *ModelDao) QueryByAccountId(ctx context.Context, accountId string) ([]*po.Model, error) {
	var result []*po.Model
	err := dao.conn.WithContext(ctx).
		Model(&po.Model{}).
		Joins("join account_model on model.model_id = account_model.model_id").
		Where("account_model.account_id = ?", accountId).
		Where("account_model.deleted_at IS NULL").
		Where("model.deleted_at IS NULL").
		Scan(&result).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "query by account id error: %v", err)
		return nil, err
	}

	return result, nil
}
