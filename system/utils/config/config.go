package config

import "github.com/spf13/viper"

// var configFiles = "/app/config/config"
var config *viper.Viper

// 读取配置文件
func init() {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigFile("C:\\Users\\14426\\GolandProjects\\xzdp-admin\\app\\config\\config.yaml")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
}
func GetConfigString(key string) string {
	return config.GetString(key)
}
