package ctx_util

import (
	"context"
	"webchat_be/biz/model/consts"
)

func GetEmail(ctx context.Context) string {
	email, _ := ctx.Value(consts.SessionKeyEmail).(string)
	return email
}

func GetDevice(ctx context.Context) string {
	device, _ := ctx.Value(consts.SessionKeyDevice).(string)
	return device
}

func GetOriginIP(ctx context.Context) string {
	ip, _ := ctx.Value(consts.SessionKeyLoginIP).(string)
	return ip
}

func GetSessionID(ctx context.Context) string {
	sessId, _ := ctx.Value(consts.SessionKeySessID).(string)
	return sessId
}

func GetAccountId(ctx context.Context) string {
	accountId, _ := ctx.Value(consts.SessionKeyAccountId).(string)
	return accountId
}
