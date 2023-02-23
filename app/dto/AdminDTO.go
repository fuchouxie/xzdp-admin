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

type AdminListReq struct {
	ID       int32  `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	Name     string `json:"name" form:"name"`
	RoleID   int32  `json:"role_id" form:"role_id"`
}

type AdminListRes struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Remark   string `json:"remark"`
	RoleID   int32  `json:"role_id"`
}

type UpdateAdminReq struct {
	ID       int32  `json:"id" form:"id" required:"true"`
	Username string `json:"username" form:"username" `
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Name     string `json:"name" form:"name"`
	Remark   string `json:"remark" form:"remark"`
	RoleID   int32  `json:"role_id" form:"roleID"`
}

type ChangePasswordReq struct {
	Username    string `json:"username" form:"username" required:"true"`
	OldPassword string `json:"old_password" form:"old_password" required:"true"`
	NewPassword string `json:"new_password" form:"new_password" required:"true"`
}
