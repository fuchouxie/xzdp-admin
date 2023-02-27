// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTbVoucherOrder = "tb_voucher_order"

// TbVoucherOrder mapped from table <tb_voucher_order>
type TbVoucherOrder struct {
	ID         int64          `gorm:"column:id;primaryKey" json:"id"`                     // 主键
	UserID     int64          `gorm:"column:user_id;not null" json:"user_id"`             // 下单的用户id
	VoucherID  int64          `gorm:"column:voucher_id;not null" json:"voucher_id"`       // 购买的代金券id
	PayType    int32          `gorm:"column:pay_type;not null;default:1" json:"pay_type"` // 支付方式 1：余额支付；2：支付宝；3：微信
	Status     int32          `gorm:"column:status;not null;default:1" json:"status"`     // 订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`       // 下单时间
	PayTime    time.Time      `gorm:"column:pay_time" json:"pay_time"`                    // 支付时间
	UseTime    time.Time      `gorm:"column:use_time" json:"use_time"`                    // 核销时间
	RefundTime time.Time      `gorm:"column:refund_time" json:"refund_time"`              // 退款时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`       // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName TbVoucherOrder's table name
func (*TbVoucherOrder) TableName() string {
	return TableNameTbVoucherOrder
}
