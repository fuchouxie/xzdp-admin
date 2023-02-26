package dto

import "xzdp-admin/system/core/myGrom"

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
	RoleID   int32  `json:"role_id" form:"role_id"`
}

type AdminListReq struct {
	ID       int32  `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	Name     string `json:"name" form:"name"`
	RoleID   int32  `json:"role_id" form:"role_id"`
	myGrom.PageLimit
}

type AdminListRes struct {
	List  []AdminListOutModel `json:"list"`
	Total int64               `json:"total"`
}

type AdminListOutModel struct {
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
	RoleID   int32  `json:"role_id" form:"role_id"`
}

type ChangePasswordReq struct {
	Username    string `json:"username" form:"username" required:"true"`
	OldPassword string `json:"old_password" form:"old_password" required:"true"`
	NewPassword string `json:"new_password" form:"new_password" required:"true"`
}

type RemoveAdminReq struct {
	ID int32 `json:"id" form:"id"  required:"true"`
}

type BatchRemoveAdminReq struct {
	IDS string `json:"ids" form:"ids"  required:"true"`
}
