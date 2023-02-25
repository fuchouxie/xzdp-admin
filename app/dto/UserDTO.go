package dto

import "xzdp-admin/system/core/myGrom"

type UserListReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	Phone    string `form:"phone" json:"phone"`
	Level    int32  `json:"level" form:"level"`
	City     string `json:"city" form:"city"`
	myGrom.PageLimit
}

type UserListRes struct {
	List  []UserOutModel `json:"list,omitempty"`
	Total int64          `json:"total"`
}

type UserOutModel struct {
	ID        int64  `json:"id,omitempty"`           // 主键
	Phone     string `json:"phone,omitempty"`        // 手机号码
	NickName  string `json:"nick_name,omitempty"`    // 昵称
	City      string `json:"city,omitempty"`         // 城市名称
	Level     int32  `json:"level,omitempty"`        // 会员级别，0~9级,0代表未开通会员
	CreatedAt string `json:"created_time,omitempty"` //创建时间
}
