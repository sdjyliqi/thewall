package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"net/url"
	"regexp"
	"strconv"
)

var (
	redisClient *redis.Client
)

func safeRedisAddr(addr string) string {
	return regexp.MustCompile(`://([^:]+):(.*)@`).ReplaceAllString(addr, "://$1:****@")
}

// redis://db:password@host:port
func parseRedisAddr(addr string) (host string, password string, db int) {
	db = 0
	u, err := url.Parse(addr)
	if err != nil {
		host = addr
	} else {
		host = u.Host
		db64, _ := strconv.ParseInt(u.User.Username(), 0, 32)

		db = int(db64)
		password, _ = u.User.Password()
	}
	glog.V(8).Infof("parse redis URI. addr = %s, host = %s, db = %d", safeRedisAddr(addr), host, db)
	return
}

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
		glog.Infof("[init] Initialize redis client failed,please check the addr:%+v,password:%+v,err:%+v", addr, password, err)
	}
	return redisClient, err
}

//GetRedisClient ...创建到redis的连接
func GetRedisClient() *redis.Client {
	return redisClient
}
