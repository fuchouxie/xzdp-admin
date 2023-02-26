package dto

import "xzdp-admin/system/core/myGrom"

type ShopListReq struct {
	ID      int64  `json:"id,omitempty" form:"id"`           // 主键
	Name    string `json:"name,omitempty" form:"name"`       // 商铺名称
	TypeID  int64  `json:"type_id,omitempty" form:"type_id"` // 商铺类型的id
	Area    string `json:"area,omitempty" form:"area"`       // 商圈，例如陆家嘴
	Address string `json:"address,omitempty" form:"address"` // 地址
	Score   int32  `json:"score,omitempty" form:"score"`     // 评分，1~5分，乘10保存，避免小数
	myGrom.PageLimit
}

type ShopListRes struct {
	List  []ShopListOutModel `json:"list"`
	Total int64              `json:"total"`
}

type ShopListOutModel struct {
	ID        int64   `json:"id,omitempty"`         // 主键
	Name      string  `json:"name,omitempty"`       // 商铺名称
	TypeID    int64   `json:"type_id,omitempty"`    // 商铺类型的id
	Images    string  `json:"images,omitempty"`     // 商铺图片，多个图片以','隔开
	Area      string  `json:"area,omitempty"`       // 商圈，例如陆家嘴
	Address   string  `json:"address,omitempty"`    // 地址
	X         float64 `json:"x,omitempty"`          // 经度
	Y         float64 `json:"y,omitempty"`          // 维度
	AvgPrice  int64   `json:"avg_price,omitempty"`  // 均价，取整数
	Sold      int32   `json:"sold,omitempty"`       // 销量
	Comments  int32   `json:"comments,omitempty"`   // 评论数量
	Score     int32   `json:"score,omitempty"`      // 评分，1~5分，乘10保存，避免小数
	OpenHours string  `json:"open_hours,omitempty"` // 营业时间，例如 10:00-22:00
}

type UpdateShopReq struct {
	ID        int64   `form:"id" json:"id,omitempty"`
	Name      string  `form:"name" json:"name,omitempty"`             // 商铺名称
	TypeID    int64   `form:"type_id" json:"type_id,omitempty"`       // 商铺类型的id
	Images    string  `form:"images" json:"images,omitempty"`         // 商铺图片，多个图片以','隔开
	Area      string  `form:"area" json:"area,omitempty"`             // 商圈，例如陆家嘴
	Address   string  `form:"address" json:"address,omitempty"`       // 地址
	X         float64 `form:"x" json:"x,omitempty"`                   // 经度
	Y         float64 `form:"y" json:"y,omitempty"`                   // 维度
	AvgPrice  int64   `form:"avg_price" json:"avg_price,omitempty"`   // 均价，取整数
	Sold      int32   `form:"sold" json:"sold,omitempty"`             // 销量
	Comments  int32   `form:"comments" json:"comments,omitempty"`     // 评论数量
	Score     int32   `form:"score" json:"score,omitempty"`           // 评分，1~5分，乘10保存，避免小数
	OpenHours string  `form:"open_hours" json:"open_hours,omitempty"` // 营业时间，例如 10:00-22:00
}

type CreateShopReq struct {
	Name      string  `form:"name" json:"name,omitempty"`             // 商铺名称
	TypeID    int64   `form:"type_id" json:"type_id,omitempty"`       // 商铺类型的id
	Images    string  `form:"images" json:"images,omitempty"`         // 商铺图片，多个图片以','隔开
	Area      string  `form:"area" json:"area,omitempty"`             // 商圈，例如陆家嘴
	Address   string  `form:"address" json:"address,omitempty"`       // 地址
	X         float64 `form:"x" json:"x,omitempty"`                   // 经度
	Y         float64 `form:"y" json:"y,omitempty"`                   // 维度
	AvgPrice  int64   `form:"avg_price" json:"avg_price,omitempty"`   // 均价，取整数
	Sold      int32   `form:"sold" json:"sold,omitempty"`             // 销量
	Comments  int32   `form:"comments" json:"comments,omitempty"`     // 评论数量
	Score     int32   `form:"score" json:"score,omitempty"`           // 评分，1~5分，乘10保存，避免小数
	OpenHours string  `form:"open_hours" json:"open_hours,omitempty"` // 营业时间，例如 10:00-22:00
}

type RemoveShopReq struct {
	ID int64 `form:"id" json:"id,omitempty"`
}

type BatchRemoveShopReq struct {
	IDS string `json:"ids" form:"ids"`
}
