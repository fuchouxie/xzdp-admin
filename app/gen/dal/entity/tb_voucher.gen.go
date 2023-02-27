// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTbVoucher = "tb_voucher"

// TbVoucher mapped from table <tb_voucher>
type TbVoucher struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键
	ShopID      int64          `gorm:"column:shop_id" json:"shop_id"`                     // 商铺id
	Title       string         `gorm:"column:title;not null" json:"title"`                // 代金券标题
	SubTitle    string         `gorm:"column:sub_title" json:"sub_title"`                 // 副标题
	Rules       string         `gorm:"column:rules" json:"rules"`                         // 使用规则
	PayValue    int64          `gorm:"column:pay_value;not null" json:"pay_value"`        // 支付金额，单位是分。例如200代表2元
	ActualValue int64          `gorm:"column:actual_value;not null" json:"actual_value"`  // 抵扣金额，单位是分。例如200代表2元
	Type        int32          `gorm:"column:type;not null;default:1" json:"type"`        // 1,普通券；2,秒杀券
	Status      int32          `gorm:"column:status;not null;default:1" json:"status"`    // 1,上架; 2,下架; 3,过期
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`      // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`      // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName TbVoucher's table name
func (*TbVoucher) TableName() string {
	return TableNameTbVoucher
}
