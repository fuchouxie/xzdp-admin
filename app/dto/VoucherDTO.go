package dto

import "xzdp-admin/system/core/myGrom"

type VoucherListReq struct {
	ShopID   int64  `json:"shop_id" form:"shop_id"` // 商铺id
	ShopName string `json:"shop_name" form:"shop_name"`
	Status   int32  `json:"status" form:"status"` // 1,上架; 2,下架; 3,过期
	Title    string `json:"title" form:"title"`   // 代金券标题
	Type     int32  `json:"type" form:"type"`     //1.普通；2.秒杀
	myGrom.PageLimit
}

type VoucherListRes struct {
	List  []VoucherOutModel `json:"list"`
	Total int64             `json:"total"`
}

type VoucherOutModel struct {
	ID          int64  `json:"id,omitempty"`
	ShopID      int64  `json:"shop_id,omitempty"` // 商铺id
	Name        string `json:"shop_name"`
	Title       string `json:"title,omitempty"`        // 代金券标题
	SubTitle    string `json:"sub_title,omitempty"`    // 副标题
	Rules       string `json:"rules,omitempty"`        // 使用规则
	PayValue    int64  `json:"pay_value,omitempty"`    // 支付金额，单位是分。例如200代表2元
	ActualValue int64  `json:"actual_value,omitempty"` // 抵扣金额，单位是分。例如200代表2元
	Type        int32  `json:"type"`                   // 0,普通券；1,秒杀券
	Status      int32  `json:"status,omitempty"`       // 1,上架; 2,下架; 3,过期
	CreatedAt   string `json:"created_time"`           //创建时间
}

type UpdateVoucherReq struct {
	ID          int64  `json:"id" form:"id"`
	Title       string `json:"title,omitempty" form:"title"`               // 代金券标题
	SubTitle    string `json:"sub_title,omitempty" form:"sub_title"`       // 副标题
	Rules       string `json:"rules,omitempty" form:"rules"`               // 使用规则
	PayValue    int64  `json:"pay_value,omitempty" form:"pay_value"`       // 支付金额，单位是分。例如200代表2元
	ActualValue int64  `json:"actual_value,omitempty" form:"actual_value"` // 抵扣金额，单位是分。例如200代表2元
	Type        int32  `json:"type" form:"type"`                           // 0,普通券；1,秒杀券
	Status      int32  `json:"status,omitempty" form:"status"`             // 1,上架; 2,下架; 3,过期
}

type CreateVoucherReq struct {
	ShopID      int64  `json:"shop_id,omitempty" form:"shop_id"`           // 商铺id
	Title       string `json:"title,omitempty" form:"title"`               // 代金券标题
	SubTitle    string `json:"sub_title,omitempty" form:"sub_title"`       // 副标题
	Rules       string `json:"rules,omitempty" form:"rules"`               // 使用规则
	PayValue    int64  `json:"pay_value,omitempty" form:"pay_value"`       // 支付金额，单位是分。例如200代表2元
	ActualValue int64  `json:"actual_value,omitempty" form:"actual_value"` // 抵扣金额，单位是分。例如200代表2元
	Type        int32  `json:"type,omitempty" form:"type"`                 // 0,普通券；1,秒杀券
	Status      int32  `json:"status,omitempty" form:"status"`             // 1,上架; 2,下架; 3,过期
}

type GetVoucherInfoReq struct {
	ID int64 `json:"id" form:"id"`
}

type RemoveVoucherReq struct {
	ID int64 `json:"id" form:"id"`
}

type BatchRemoveVoucherReq struct {
	IDS string `json:"ids" form:"ids"`
}
