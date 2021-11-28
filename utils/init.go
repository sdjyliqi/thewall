package utils

import "email-center/conf"

//init() ...完成工具类相关的初始化相关事项
func init() {
	InitMySQL(conf.DefaultConfig.DBMysql, true)
}
