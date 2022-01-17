package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	redisClient *redis.Client
)

//InitRedis ...初始化redis，记得要先解密
func InitRedis(addr string, password string) (*redis.Client, error) {
	fmt.Println("Init redis ", addr, password)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		glog.Infof("[init] Initialize redis client failed,please check the addr:%+v,err:%+v", addr, err)
	}
	return redisClient, err
}

//GetRedisClient ...创建到redis的连接
func GetRedisClient() *redis.Client {
	return redisClient
}
