package dto

type LoginReq struct {
	Username string `json:"username" form:"username" required:"true"`
	Password string `json:"password" form:"password" required:"true" `
}

type RegisterReq struct {
	Username string `json:"username" form:"username" required:"true"`
	Password string `json:"password" form:"password" required:"true" `
	Phone    string `json:"phone" form:"phone"`
	Name     string `json:"name" form:"name"`
	Remark   string `json:"remark" form:"remark"`
	RoleID   int32  `json:"role_id" form:"roleID"`
}
