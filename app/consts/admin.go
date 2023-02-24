package consts

const (
	SuperAdmin = 1
	ShopAdmin  = 2
)

var EnumAdminType = map[int32]string{
	SuperAdmin: "超级管理员",
	ShopAdmin:  "运营管理员",
}
