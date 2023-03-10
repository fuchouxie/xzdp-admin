// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTbUser = "tb_user"

// TbUser mapped from table <tb_user>
type TbUser struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键
	Phone     string         `gorm:"column:phone;not null" json:"phone"`                                     // 手机号码
	Password  string         `gorm:"column:password" json:"password"`                                        // 密码，加密存储
	NickName  string         `gorm:"column:nick_name" json:"nick_name"`                                      // 昵称，默认是用户id
	Icon      string         `gorm:"column:icon" json:"icon"`                                                // 人物头像
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UserInfo  TbUserInfo     `gorm:"foreignKey:user_id" json:"user_info"`
}

// TableName TbUser's table name
func (*TbUser) TableName() string {
	return TableNameTbUser
}
