package router

import (
	"github.com/gin-gonic/gin"
	"thewall/handle"
	"thewall/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Logger())
	r.GET("/ping", handle.Ping)                 //测活接口,1
	r.POST("/uc/login", handle.UCLogin)         //利用邮件和密码登录
	r.POST("/uc/register", handle.UCRegister)   //用户注册,包括邮件、密码、昵称、密码、code验证码五个维度数据
	r.POST("/uc/reset", handle.UCResetPassword) //充值密码，包括用户注册邮件、新密码、验证码三个维度数据
	r.GET("/uc/code", handle.UCGetCode)         //获取验证码

	r.GET("/crop/type/items", handle.GetCropTypeAllItems) //获取农作物类型
	//r.GET("/crop/type/page", handle.GetCropTypeItemsByPage) //分页获取农作物类型

	//个人理解用不着了。
	//r.GET("/device/items", handle.GetDeviceAllItems)
	//r.GET("/device/page", handle.GetDeviceItemsByPage)

	r.POST("/field/add", handle.FieldAdd)         //某用户增加农场
	r.POST("/field/edit", handle.FieldEdit)       //修改农场信息
	r.POST("/field/del", handle.FieldDel)         //某用户删除农场信息
	r.GET("/field/items", handle.FieldGetItems)   //查询某个用户的所属农场
	r.POST("/field/plant", handle.FieldPlanting)  //开始种植
	r.POST("/field/harvest", handle.FieldHarvest) //开始收获
	r.POST("/field/weigh", handle.FieldWeigh)     //开始称重
	r.POST("/field/end", handle.FieldEnded)       //终止流程

	r.GET("/field/lines", handle.FieldProbeLines) //查看某土地上的所有probe K线图
	//todo 查找该土地的所有种植的历史记录
	//tood 称重

	r.GET("/sensor/itemsByPage", handle.GetSensorItemsByPage)       //后台分页获取Sensor列表
	r.POST("/sensor/add", handle.AddSensor)                         //后台添加Sensor
	r.GET("/sensor/itemsByField", handle.GetSensorItemsByField)     //获取绑定Field的传感器列表
	r.GET("/sensor/itemsByGateway", handle.GetSensorItemsByGateway) //获取绑定Gateway的传感器列表
	r.GET("/sensor/itemsByUser", handle.GetSensorItemsByUser)       //获取用户绑定的Sensor列表
	r.GET("/sensor/item", handle.GetSensorItem)                     //获取Sensor信息
	r.POST("/sensor/bindByUser", handle.BindSensorByUser)           //APP用户绑定Sensor
	r.POST("/sensor/unbindByUser", handle.UnbindSensorByUser)       //APP用户解绑Sensor
	r.POST("/sensor/editByUser", handle.EditSensorByUser)           //APP用户修改Sensor信息
	r.POST("/sensor/gather", handle.GatherData)                     //收集一次上报数据
	r.GET("/sensor/line", handle.GetLineItems)
	//todo r.GET("/sensor/line", handle.GetLineItems) 数据导出生产csv
	//todo yanghao   add sensor 方法
	//雀巢提供api，本服务把数据给数据送过去。

	//查询某个区间的数据

	r.GET("/gateway/itemsByPage", handle.GetGatewayItemsByPage) //后台分页获取Gateway列表
	r.GET("/gateway/item", handle.GetGatewayItem)               //后台获取Gateway信息
	r.POST("/gateway/add", handle.AddGateway)                   //后台添加Gateway
	r.POST("/gateway/edit", handle.EditGateway)                 //后台修改Gateway
	r.GET("/gateway/itemsByUser", handle.GetGatewayItemsByUser) //APP用户获取绑定的Gateway列表
	r.GET("/gateway/itemByUser", handle.GetGatewayItemByUser)   //APP用户获取Gateway信息
	r.POST("/gateway/bindByUser", handle.BindGatewayByUser)     //APP用户绑定Gateway
	r.POST("/gateway/unbindByUser", handle.UnbindGatewayByUser) //APP用户解绑Gateway
	r.POST("/gateway/editByUser", handle.EditGatewayByUser)     //APP用户修改Gateway
}
