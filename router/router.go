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

	r.POST("/field/add", handle.FieldAdd)       //某用户增加农场
	r.POST("/field/edit", handle.FieldEdit)     //修改农场信息
	r.POST("/field/del", handle.FieldDel)       //某用户删除农场信息
	r.GET("/field/items", handle.FieldGetItems) //查询某个用户的所属农场

	//todo
	r.GET("/sensor/itemsbyfield", handle.GetSensorItemsByField) //获取这个农场的senser 列表，写完了
	r.GET("/sensor/delbyfield", handle.FieldGetItems)           //某个农场结束绑定的sensor
	r.GET("/sensor/itemsbyuser", handle.GetSensorItemsByUser)   //当前用户的senser列表，写完了

	r.GET("/gateway/items", handle.GetGatewayAllItems)      //获取网关列表
	r.POST("/gateway/add", handle.AddGateway)               //添加网关
	r.POST("/gateway/edit", handle.EditGateway)             //修改网关
	r.POST("/gateway/bindByUser", handle.BindGatewayByUser) //网关绑定用户
}
