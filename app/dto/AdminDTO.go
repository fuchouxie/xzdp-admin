package dto

type LoginReq struct {
	Username string `json:"username" form:"username" required:"true"`
	Password string `json:"password" form:"password" required:"true" `
}
