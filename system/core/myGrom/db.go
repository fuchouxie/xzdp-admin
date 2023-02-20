package myGrom

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"xzdp-admin/system/utils/config"
)

var Db *gorm.DB

// 初始化
func init() {
	//获取连接数据库配置项
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
	var err error
	Db, err = gorm.Open(tmp, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("orm myGrom connect error")
	} else {
	}
}
