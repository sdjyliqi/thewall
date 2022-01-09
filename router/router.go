package router

import (
	"github.com/gin-gonic/gin"
	"thewall/handle"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", handle.Ping)                 //测活接口
	r.POST("/uc/login", handle.UCLogin)         //利用邮件和密码登录
	r.POST("/uc/register", handle.UCRegister)   //用户注册,包括邮件、密码、昵称、密码、code验证码五个维度数据
	r.POST("/uc/reset", handle.UCResetPassword) //充值密码，包括用户注册邮件、新密码、验证码三个维度数据
	r.GET("/uc/code", handle.UCGetCode)         //获取验证码

	r.GET("/crop/type/items", handle.GetCropTypeAllItems)   //获取农作物类型
	r.GET("/crop/type/page", handle.GetCropTypeItemsByPage) //分页获取农作物类型

	r.GET("/device/items", handle.GetDeviceAllItems)
	r.GET("/device/page", handle.GetDeviceItemsByPage)

	r.POST("/field/add", handle.FieldAdd)       //某用户增加地
	r.POST("/field/edit", handle.FieldEdit)     //修改田地信息
	r.POST("/field/del", handle.FieldDel)       //某用户删除田地信息
	r.GET("/field/items", handle.FieldGetItems) //查询某个用户的所属土地

}
