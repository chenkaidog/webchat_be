package dto

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,max=64"`
	Password string `form:"password" json:"password" binding:"required,max=128"`
}

type LoginResp struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	Email     string `json:"email"`
}

type AccountInfoQueryResp struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}

type LogoutReq struct {
}

type LogoutResp struct {
}

type PasswordUpdateReq struct {
	Password    string `form:"password" json:"password" binding:"required,max=128"`
	PasswordNew string `form:"password_new" json:"password_new" binding:"required,min=8,max=128,alphanumunicode"`
}

type PasswordUpdateResp struct {
}

type ForgetPasswordReq struct {
	Username string `form:"username" json:"username" binding:"required,max=64"`
}

type ForgetPasswordResp struct{}

type ResetPasswordReq struct {
	Password   string `form:"password" json:"password" binding:"required,max=128"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required,max=10"`
}

type ResetPasswordResp struct{}

type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,max=64"`
	Email    string `form:"email" json:"email" binding:"required,max=64,email"`
	Password string `form:"password" json:"password" binding:"required,max=128,alphanumunicode"`
}

type RegisterResp struct{}

type RegisterVerifyReq struct {
	Email      string `form:"email" json:"email" binding:"required,max=64,email"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required,max=10"`
}

type RegisterVerifyResp struct{}
