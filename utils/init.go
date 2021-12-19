package utils

import "thewall/conf"

//init() ...完成工具类相关的初始化相关事项
func init() {
	InitMySQL(conf.DefaultConfig.DBMysql, true)
	InitRedis(conf.DefaultConfig.RedisAddr, conf.DefaultConfig.RedisPassword)
}
