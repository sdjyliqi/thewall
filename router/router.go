package router

import (
	"email-center/handle"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/uc/hello", handle.HelloWord) //获取一次邮件比例
	r.POST("/crop/items", handle.GetCropAllItems)
	r.GET("/crop/page", handle.GetCropItemsByPage) //
}
