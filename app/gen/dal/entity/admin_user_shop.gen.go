// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameAdminUserShop = "admin_user_shop"

// AdminUserShop mapped from table <admin_user_shop>
type AdminUserShop struct {
	AdminUserID int32 `gorm:"column:admin_user_id;not null" json:"admin_user_id"` // 后台用户id
	ShopID      int32 `gorm:"column:shop_id;not null" json:"shop_id"`             // 商店id
}

// TableName AdminUserShop's table name
func (*AdminUserShop) TableName() string {
	return TableNameAdminUserShop
}
