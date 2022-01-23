package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thewall/errs"
	"thewall/model"
	"thewall/utils"
)

type GatewayDto struct {
	Id        int     `json:"id"`
	UserId    int     `json:"user_id"`
	FieldId   int     `json:"field_id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	SensorIds []int   `json:"sensor_ids"`
}

//GetGatewayItemsByUser ... APP获取用户绑定的Gateway列表数据
func GetGatewayItemsByUser(c *gin.Context) {
	strUserId, _ := c.GetQuery("user_id")
	userId := utils.Convert2Int(strUserId)
	if userId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	items, err := model.GatewayModel.GetItemsByUser(userId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetGatewayItemsByPage ... 后台分页获取Gateway全量数据
func GetGatewayItemsByPage(c *gin.Context) {
	strPage, _ := c.GetQuery("page")
	if strPage == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	pageId := utils.Convert2Int(strPage)
	if pageId < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	items, err := model.GatewayModel.GetItemsByPage(pageId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetGatewayItem ... 后台获取Gateway信息
func GetGatewayItem(c *gin.Context) {
	strId, _ := c.GetQuery("id")
	if strId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	id := utils.Convert2Int(strId)
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	item, err := model.GatewayModel.GetItemByID(id)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": item})
	return
}

//GetGatewayItemByUser ... APP获取Gateway信息
func GetGatewayItemByUser(c *gin.Context) {
	strId, _ := c.GetQuery("id")
	strUserId, _ := c.GetQuery("user_id")
	if strId == "" || strUserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	id := utils.Convert2Int(strId)
	userId := utils.Convert2Int(strUserId)
	if id <= 0 || userId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	item, err := model.GatewayModel.GetItemByUser(id, userId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": item})
	return
}

//AddGateway ... 后台添加一条Gateway数据
func AddGateway(c *gin.Context) {
	item := model.IotGateway{}
	bindErr := c.BindJSON(&item)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	ok, err := model.GatewayModel.AddItem(&item)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//EditGateway ... 后台编辑一条Gateway数据
func EditGateway(c *gin.Context) {
	itemDto := GatewayDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Id <= 0 || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	item := model.IotGateway{
		Id:        itemDto.Id,
		Longitude: itemDto.Longitude,
		Latitude:  itemDto.Latitude,
		WriteUid:  itemDto.UserId,
	}
	_, err := model.GatewayModel.UpdateItemByID(&item)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	//Sensor绑定Gateway
	if len(itemDto.SensorIds) > 0 {
		errEX := model.SensorModel.SensorBindGateway(itemDto.SensorIds, itemDto.Id, itemDto.UserId)
		if errEX != errs.Succ {
			c.JSON(http.StatusInternalServerError, gin.H{"code": errEX.Code, "msg": errEX.MessageEN, "data": nil})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": true})
	return
}

//EditGatewayByUser ... APP编辑一条Gateway数据
func EditGatewayByUser(c *gin.Context) {
	itemDto := GatewayDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Id <= 0 || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	item := model.IotGateway{
		Id:        itemDto.Id,
		UserId:    itemDto.UserId,
		Longitude: itemDto.Longitude,
		Latitude:  itemDto.Latitude,
	}
	_, err := model.GatewayModel.UpdateItemByUser(&item)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	//Sensor绑定Gateway
	if len(itemDto.SensorIds) > 0 {
		errEX := model.SensorModel.SensorBindGateway(itemDto.SensorIds, itemDto.Id, itemDto.UserId)
		if errEX != errs.Succ {
			c.JSON(http.StatusInternalServerError, gin.H{"code": errEX.Code, "msg": errEX.MessageEN, "data": nil})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": true})
	return
}

//BindGatewayByUser ... APP用户绑定Gateway
func BindGatewayByUser(c *gin.Context) {
	itemDto := GatewayDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Code == "" || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	ok, err := model.GatewayModel.BindItemByUser(itemDto.Code, itemDto.UserId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//UnbindGatewayByUser ... APP用户解绑
func UnbindGatewayByUser(c *gin.Context) {
	itemDto := GatewayDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Id <= 0 || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	ok, err := model.GatewayModel.UnbindItemByUser(itemDto.Id, itemDto.UserId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}
