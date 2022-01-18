package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"thewall/conf"
	"thewall/handle"
	"thewall/router"
)

func init() {
	var ymlPath string
	flag.StringVar(&ymlPath, "c", "", "configuration file")
	flag.Parse()
	if ymlPath == "" {
		log.Println("You must input path of the yml ....")
	}
	handle.LoadTranslateDic()
}

func main() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.RunPort))
}
