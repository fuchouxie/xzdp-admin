package myGrom

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
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

	logFile, err := os.OpenFile("./logs/db.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("open logFile error : " + err.Error())
	}
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)

	Db, err = gorm.Open(tmp, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("orm myGrom connect error")
	}
}

type PageLimit struct {
	Page     int `form:"page,default=1" json:"page,default=1"`
	PageSize int `form:"page_size,default=10" json:"page_size,default=10"`
	OffSet   int
}

func Paginate(p PageLimit) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.Page
		if page == 0 {
			page = 1
		}

		pageSize := p.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
