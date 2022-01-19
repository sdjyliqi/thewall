package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"os"
	"thewall/conf"
	"thewall/handle"
	"thewall/router"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常，请确定用户权限")
	}
	log.SetOutput(logFile)
	go handle.LoadTranslateDic()
}

func main() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.RunPort))
}
