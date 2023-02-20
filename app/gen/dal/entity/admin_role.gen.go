// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAdminRole = "admin_role"

// AdminRole mapped from table <admin_role>
type AdminRole struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`     // 角色名称
	Remark    string         `gorm:"column:remark;not null" json:"remark"` // 备注
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName AdminRole's table name
func (*AdminRole) TableName() string {
	return TableNameAdminRole
}
