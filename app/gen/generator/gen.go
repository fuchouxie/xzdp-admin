package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"xzdp-admin/system/utils/config"
)

var g *gen.Generator

func main() {
	initGen()
	shop := g.GenerateModel("tb_shop")
	blog := g.GenerateModel("tb_blog")
	user := g.GenerateModel("tb_user")
	userInfo := g.GenerateModel("tb_user_info")
	voucher := g.GenerateModel("tb_voucher")
	voucherOrder := g.GenerateModel("tb_voucher_order")

	adminUser := g.GenerateModel("admin_user")
	adminRole := g.GenerateModel("admin_role")
	adminUserShop := g.GenerateModel("admin_user_shop")
	g.ApplyBasic(shop)
	g.ApplyBasic(blog)
	g.ApplyBasic(user)
	g.ApplyBasic(userInfo)
	g.ApplyBasic(voucher)
	g.ApplyBasic(voucherOrder)
	g.ApplyBasic(adminUser)
	g.ApplyBasic(adminRole)
	g.ApplyBasic(adminUserShop)
	g.Execute()
}

func initGen() {
	g = gen.NewGenerator(gen.Config{
		OutPath:      "../dal/query",
		ModelPkgPath: "../dal/entity",
	})
	//host
	host := config.GetConfigString("database.host")
	//port
	port := config.GetConfigString("database.port")
	//database
	database := config.GetConfigString("database.database")
	//username
	username := config.GetConfigString("database.username")
	//password
	password := config.GetConfigString("database.password")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	tmp := mysql.Open(dsn)
	db, _ := gorm.Open(tmp)
	g.UseDB(db)
}
