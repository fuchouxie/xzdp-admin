package dto

import "xzdp-admin/system/core/myGrom"

type BlogListReq struct {
	ShopID int64  `json:"shop_id" form:"shop_id"` // 商户id
	UserID int64  `json:"user_id" form:"user_id"` // 用户id
	Title  string `json:"title" form:"title"`     // 标题
	myGrom.PageLimit
}

type BlogListRes struct {
	List  []BlogOutModel `json:"list"`
	Total int64          `json:"total"`
}

type BlogOutModel struct {
	ID        int64  `json:"id"`      // 主键
	ShopID    int64  `json:"shop_id"` // 商户id
	UserID    int64  `json:"user_id"` // 用户id
	Name      string `json:"shop_name"`
	NickName  string `json:"user_name"`
	Content   string `json:"content"`
	Title     string `json:"title"`      // 标题
	Liked     int32  `json:"liked"`      // 点赞数量
	Comments  int32  `json:"comments"`   // 评论数量
	CreatedAt string `json:"created_at"` // 创建时间
}

type GetBlogInfoReq struct {
	ID int32 `json:"id" form:"id" required:"true"`
}

type RemoveBlogReq struct {
	ID int32 `json:"id" form:"id" required:"true"`
}

type BatchRemoveBlogReq struct {
	IDS string `json:"ids" form:"ids" required:"true"`
}
