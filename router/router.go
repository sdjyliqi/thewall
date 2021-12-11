package router

import (
	"email-center/handle"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/uc/login", handle.UCLogin)       //利用邮件和密码登录
	r.POST("/uc/register", handle.UCRegister) //用户注册邮件和密码

	r.GET("/uc/hello", handle.HelloWord) //获取一次邮件比例
	r.GET("/crop/items", handle.GetCropAllItems)
	r.GET("/crop/page", handle.GetCropItemsByPage) //

	r.GET("/device/items", handle.GetDeviceAllItems)
	r.GET("/device/page", handle.GetDeviceItemsByPage)
}
