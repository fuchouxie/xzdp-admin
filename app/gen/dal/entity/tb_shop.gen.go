// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameTbShop = "tb_shop"

// TbShop mapped from table <tb_shop>
type TbShop struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`               // 主键
	Name       string    `gorm:"column:name;not null" json:"name"`                                // 商铺名称
	TypeID     int64     `gorm:"column:type_id;not null" json:"type_id"`                          // 商铺类型的id
	Images     string    `gorm:"column:images;not null" json:"images"`                            // 商铺图片，多个图片以','隔开
	Area       string    `gorm:"column:area" json:"area"`                                         // 商圈，例如陆家嘴
	Address    string    `gorm:"column:address;not null" json:"address"`                          // 地址
	X          float64   `gorm:"column:x;not null" json:"x"`                                      // 经度
	Y          float64   `gorm:"column:y;not null" json:"y"`                                      // 维度
	AvgPrice   int64     `gorm:"column:avg_price" json:"avg_price"`                               // 均价，取整数
	Sold       int32     `gorm:"column:sold;not null" json:"sold"`                                // 销量
	Comments   int32     `gorm:"column:comments;not null" json:"comments"`                        // 评论数量
	Score      int32     `gorm:"column:score;not null" json:"score"`                              // 评分，1~5分，乘10保存，避免小数
	OpenHours  string    `gorm:"column:open_hours" json:"open_hours"`                             // 营业时间，例如 10:00-22:00
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
}

// TableName TbShop's table name
func (*TbShop) TableName() string {
	return TableNameTbShop
}
