package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	"net/http"
	"webchat_be/biz/dao"
	"webchat_be/biz/handler/service"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/model/dto"
	"webchat_be/biz/model/errs"
	"webchat_be/biz/util/encode"
	"webchat_be/biz/util/origin"
	"webchat_be/biz/util/random"
)

// Login 用户登录接口
//
//	@Tags			account
//	@Summary		用户登录接口
//	@Description	用户登录接口
//	@Accept			json
//	@Produce		json
//	@Param			req		body		dto.LoginReq	true	"login request body"
//	@Success		200		{object}	dto.CommonResp{data=dto.LoginResp}
//	@Header			200		{string}	X-CSRF-TOKEN	"csrf token"
//	@Header			200		{string}	set-cookie		"cookie"
//	@Failure		400,500	{object}	dto.CommonResp
//	@Router			/api/v1/account/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var loginReq dto.LoginReq
	if stdErr := c.BindAndValidate(&loginReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}
	session := sessions.DefaultMany(c, consts.SessionNameAccount)

	bizResp, bizErr := service.AccountLogin(ctx, &service.LoginRequest{
		SessID:   session.ID(),
		Username: loginReq.Username,
		Password: loginReq.Password,
		IP:       origin.GetIp(c),
		Device:   origin.GetDevice(c),
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	session.Set(consts.SessionKeyAccountId, bizResp.AccountId)
	session.Set(consts.SessionKeyLoginIP, origin.GetIp(c))
	session.Set(consts.SessionKeyDevice, origin.GetDevice(c))

	csrfToken := random.RandStr(64)
	csrfSalt := random.RandStr(64)
	session.Set(consts.SessionKeyCsrfSalt, csrfSalt)
	session.Set(consts.SessionKeyCsrfToken, encode.EncodePassword(csrfSalt, csrfToken))

	if err := session.Save(); err != nil {
		hlog.CtxInfof(ctx, "session save fail, %v", err)
		dto.FailResp(c, errs.ServerError)
		return
	}

	c.Header(consts.HeaderKeyCsrfToken, csrfToken)
	dto.SuccessResp(c, &dto.LoginResp{
		AccountID: bizResp.AccountId,
		Username:  bizResp.Username,
		Email:     bizResp.Email,
		Status:    bizResp.Status,
	})
	return
}

// GetAccountInfo 用户信息查询接口
//
//	@Tags			account
//	@Summary		用户信息查询接口
//	@Description	用户信息查询接口
//	@Accept			json
//	@Produce		json
//	@Param			X-CSRF-TOKEN	header		string	true	"csrf token"
//	@Success		200				{object}	dto.CommonResp{data=dto.AccountInfoQueryResp}
//	@Header			200				{string}	set-cookie	"cookie"
//	@Failure		400,500			{object}	dto.CommonResp
//	@Router			/api/v1/account/info [GET]
func GetAccountInfo(ctx context.Context, c *app.RequestContext) {
	accountId := ctx.Value(consts.SessionKeyAccountId).(string)
	accountInfo, err := dao.NewAccountDao().QueryByAccountId(ctx, accountId)
	if err != nil {
		dto.FailResp(c, errs.ServerError)
		return
	}
	if accountInfo == nil {
		hlog.CtxInfof(ctx, "account_id not exists: %s", accountId)
		dto.FailResp(c, errs.AccountNotExistError)
		return
	}

	dto.SuccessResp(c, &dto.AccountInfoQueryResp{
		AccountID: accountInfo.AccountID,
		Username:  accountInfo.Username,
		Status:    accountInfo.Status,
		Email:     accountInfo.Email,
	})
	return
}

// Logout 用户登出接口
//
//	@Tags			account
//	@Summary		用户登出接口
//	@Description	用户登出接口
//	@Accept			json
//	@Produce		json
//	@Param			req				body		dto.LogoutReq	true	"logout request body"
//	@Param			X-CSRF-TOKEN	header		string			true	"csrf token"
//	@Success		200				{object}	dto.CommonResp{data=dto.LogoutResp}
//	@Header			200				{string}	set-cookie	"cookie"
//	@Failure		400,500			{object}	dto.CommonResp
//	@Router			/api/v1/account/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var logoutReq dto.LogoutReq
	if stdErr := c.BindAndValidate(&logoutReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}

	session := sessions.DefaultMany(c, consts.SessionNameAccount)
	session.Options(sessions.Options{MaxAge: -1})
	if err := session.Save(); err != nil {
		hlog.CtxInfof(ctx, "session save fail, %v", err)
		dto.FailResp(c, errs.ServerError)
		return
	}

	accountId := ctx.Value(consts.SessionKeyAccountId).(string)
	_ = service.RemoveSession(ctx, accountId, session.ID())

	dto.SuccessResp(c, &dto.LogoutResp{})
}

// UpdatePassword 用户修改密码接口
//
//	@Tags			account
//	@Summary		用户修改密码接口
//	@Description	用户修改密码接口
//	@Accept			json
//	@Produce		json
//	@Param			req				body		dto.PasswordUpdateReq	true	"password update request body"
//	@Param			X-CSRF-TOKEN	header		string					true	"csrf token"
//	@Success		200				{object}	dto.CommonResp{data=dto.PasswordUpdateResp}
//	@Header			200				{string}	set-cookie	"cookie"
//	@Failure		400,500			{object}	dto.CommonResp
//	@Router			/api/v1/account/update_password [POST]
func UpdatePassword(ctx context.Context, c *app.RequestContext) {
	var updateReq dto.PasswordUpdateReq
	if stdErr := c.BindAndValidate(&updateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}

	accountId := ctx.Value(consts.SessionKeyAccountId).(string)

	if bizErr := service.AccountUpdatePassword(ctx, &service.PasswordUpdateRequest{
		AccountId:   accountId,
		Password:    updateReq.Password,
		PasswordNew: updateReq.PasswordNew,
	}); bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}

	session := sessions.DefaultMany(c, consts.SessionNameAccount)
	session.Options(sessions.Options{MaxAge: -1})
	if err := session.Save(); err != nil {
		hlog.CtxInfof(ctx, "session save fail, %v", err)
		dto.FailResp(c, errs.ServerError)
		return
	}

	_ = service.RemoveSession(ctx, accountId, session.ID())
	dto.SuccessResp(c, &dto.PasswordUpdateResp{})
}

// ForgetPassword 用户忘记密码接口
//
//	@Tags			account
//	@Summary		用户忘记密码接口
//	@Description	用户忘记密码接口，请求获取验证码进行重置
//	@Accept			json
//	@Produce		json
//	@Param			req		body		dto.ForgetPasswordReq	true	"password forget request body"
//	@Success		200		{object}	dto.CommonResp{data=dto.ForgetPasswordResp}
//	@Header			200		{string}	set-cookie	"cookie"
//	@Failure		400,500	{object}	dto.CommonResp
//	@Router			/api/v1/account/forget_password [POST]
func ForgetPassword(ctx context.Context, c *app.RequestContext) {
	var logoutReq dto.ForgetPasswordReq
	if stdErr := c.BindAndValidate(&logoutReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}

	dto.SuccessResp(c, &dto.ForgetPasswordResp{})
	return
}

// ResetPassword 用户重置密码接口
//
//	@Tags			account
//	@Summary		用户修改密码接口
//	@Description	用户修改密码接口
//	@Accept			json
//	@Produce		json
//	@Param			req		body		dto.ResetPasswordReq	true	"password reset request body"
//	@Success		200		{object}	dto.CommonResp{data=dto.ResetPasswordResp}
//	@Header			200		{string}	set-cookie	"cookie"
//	@Failure		400,500	{object}	dto.CommonResp
//	@Router			/api/v1/account/reset_password [POST]
func ResetPassword(ctx context.Context, c *app.RequestContext) {
	var logoutReq dto.ResetPasswordReq
	if stdErr := c.BindAndValidate(&logoutReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}

	dto.SuccessResp(c, &dto.ResetPasswordResp{})
	return
}

// Register 用户注册接口
//
//	@Tags			account
//	@Summary		用户注册接口
//	@Description	用户注册接口，请求后获取验证码，然后才能创建
//	@Accept			json
//	@Produce		json
//	@Param			req		body		dto.RegisterReq	true	"register request body"
//	@Success		200		{object}	dto.CommonResp{data=dto.RegisterResp}
//	@Header			200		{string}	set-cookie	"cookie"
//	@Failure		400,500	{object}	dto.CommonResp
//	@Router			/api/v1/account/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var registerReq dto.RegisterReq
	if stdErr := c.BindAndValidate(&registerReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}

	salt := random.RandStr(16)
	password := encode.EncodePassword(salt, registerReq.Password)
	hlog.CtxInfof(ctx, "salt: %+v\npassword: %+v\nencode: %+v",
		[]byte(salt), []byte(registerReq.Password), []byte(password))

	dto.SuccessResp(c, &dto.RegisterResp{})
	return
}

// RegisterVerify 用户注册验证接口
//
//	@Tags			account
//	@Summary		用户注册验证接口
//	@Description	用户注册验证接口
//	@Accept			json
//	@Produce		json
//	@Param			req		body		dto.RegisterVerifyReq	true	"register request body"
//	@Success		200		{object}	dto.CommonResp{data=dto.RegisterVerifyResp}
//	@Failure		400,500	{object}	dto.CommonResp
//	@Router			/api/v1/account/register_verify [POST]
func RegisterVerify(ctx context.Context, c *app.RequestContext) {
	var registerReq dto.RegisterVerifyReq
	if stdErr := c.BindAndValidate(&registerReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("body invalid", http.StatusBadRequest)
		return
	}

	dto.SuccessResp(c, &dto.RegisterVerifyResp{})
	return
}
