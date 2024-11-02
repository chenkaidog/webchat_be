package dao

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"time"
	"webchat_be/biz/db/mysql"
	"webchat_be/biz/model/po"
)

type LoginRecordDao struct {
	conn *mysql.DbConn
}

func NewLoginRecordDao(tx ...*gorm.DB) *LoginRecordDao {
	return &LoginRecordDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (dao *LoginRecordDao) Create(ctx context.Context, record *po.LoginRecord) error {
	return dao.conn.WithContext(ctx).Create(record).Error
}

// QueryByIP 查询最近一段时间的登录记录
func (dao *LoginRecordDao) QueryByIP(ctx context.Context, ip string, startTime time.Time) ([]*po.LoginRecord, error) {
	var resultList []*po.LoginRecord
	err := dao.conn.WithContext(ctx).
		Where("ip", ip).
		Where("created_at>?", startTime).
		Order("created_at DESC").
		Find(&resultList).
		Error
	if err != nil {
		hlog.CtxErrorf(ctx, "LoginRecordDao.QueryByIP err: %v", err)
		return nil, err
	}
	return resultList, nil
}
