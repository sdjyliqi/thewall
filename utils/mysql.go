package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"sync"
)

var mysqlOnce sync.Once
var msqlEngine *xorm.Engine

//InitMySQL ...初始化mysql连接
func InitMySQL(addr string, showSQL bool) (*xorm.Engine, error) {
	var err error
	mysqlOnce.Do(func() {
		msqlEngine, err = xorm.NewEngine("mysql", addr)
		msqlEngine.ShowSQL(showSQL)
		if err != nil {
			glog.Fatalf("Initialize mysql client failed,err:%+v,please check the addr:%s,", err, addr)
		}
	})
	return msqlEngine, err
}

//GetMysqlClient ...获取mysql客户端连接的方法
func GetMysqlClient() *xorm.Engine {
	return msqlEngine
}
