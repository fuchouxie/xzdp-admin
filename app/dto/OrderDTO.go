package dto

import (
	"xzdp-admin/system/core/myGrom"
)

type OrderListReq struct {
	UserID   int64  `form:"user_id" json:"user_id,omitempty"`
	UserName string `form:"nick_name" json:"nick_name,omitempty"`
	PayType  int32  `form:"pay_type" json:"pay_type,omitempty"`
	Status   int32  `form:"status" json:"status,omitempty"`
	myGrom.PageLimit
}

type OrderListRes struct {
	List  []OrderOutModel `json:"list"`
	Total int64           `json:"total"`
}

type OrderOutModel struct {
	ID         int64  `json:"id"`      // 主键
	UserID     int64  `json:"user_id"` // 下单的用户id
	NickName   string `json:"nick_name"`
	VoucherID  int64  `json:"voucher_id"` // 购买的代金券id
	Title      string `json:"voucher_title"`
	PayType    int32  `json:"pay_type"`     // 支付方式 1：余额支付；2：支付宝；3：微信
	Status     int32  `json:"status"`       // 订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款
	CreatedAt  string `json:"created_time"` // 下单时间
	PayTime    string `json:"pay_time"`     // 支付时间
	UseTime    string `json:"use_time"`     // 核销时间
	RefundTime string `json:"refund_time"`  // 退款时间
}

type GetOrderInfoReq struct {
	ID int64 `json:"id" form:"id" required:"true"`
}
