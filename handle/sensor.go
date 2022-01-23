package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"thewall/errs"
	"thewall/model"
	"thewall/utils"
	"time"
)

type SensorDto struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	FieldId   int     `json:"field_id"`
	UserId    int     `json:"user_id"`
	GatewayId int     `json:"gateway_id"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type SensorItemDto struct {
	Id          int                `json:"id"`
	UserId      int                `json:"user_id"`
	Name        string             `json:"name"`
	Code        string             `json:"code"`
	FieldName   string             `json:"field_name"`
	GatewayCode string             `json:"gateway_code"`
	Longitude   float32            `json:"longitude"`
	Latitude    float32            `json:"latitude"`
	Depths      []*model.ProbeItem `json:"depths"`
}

type SensorGather struct {
	SensorID     int    `json:"sensor_id"`
	EtlTimeStamp string `json:"etl_timestamp"`
	Value        int    `json:"value"`
}

//GetSensorAllItems ... 获取Sensor全量数据
func GetSensorAllItems(c *gin.Context) {
	items, err := model.SensorModel.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetSensorItemsByPage ... 分页获取Sensor全量数据
func GetSensorItemsByPage(c *gin.Context) {
	strPage, _ := c.GetQuery("page")
	if strPage == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	pageId, err := strconv.Atoi(strPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	items, err := model.SensorModel.GetItemsByPage(pageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//AddSensor ... 添加一条Sensor数据
func AddSensor(c *gin.Context) {
	item := model.IotSensor{}
	bindErr := c.BindJSON(&item)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if item.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if item.Name == "" {
		item.Name = item.Code
	}
	ok, err := model.SensorModel.AddItem(&item)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//GetSensorItemsByField ... 获取绑定Field的传感器列表
func GetSensorItemsByField(c *gin.Context) {
	strField, _ := c.GetQuery("field_id")
	if strField == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	FID := utils.Convert2Int(strField)
	items, errEx := model.SensorModel.GetItemsByField(FID)
	if errEx != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetSensorItemsByGateway ... 获取绑定Gateway的传感器列表
func GetSensorItemsByGateway(c *gin.Context) {
	strGatewayId, _ := c.GetQuery("gateway_id")
	if strGatewayId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	gatewayId := utils.Convert2Int(strGatewayId)
	if gatewayId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	items, errEx := model.SensorModel.GetItemsByGateway(gatewayId)
	if errEx != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetSensorItemsByUser ... 获取某用户的传感器列表
func GetSensorItemsByUser(c *gin.Context) {
	strUID, _ := c.GetQuery("user_id")
	if strUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	UID := utils.Convert2Int(strUID)
	items, errEx := model.SensorModel.GetItemsByUser(UID)
	if errEx != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetSensorItem ... 获取Sensor信息
func GetSensorItem(c *gin.Context) {
	strId, _ := c.GetQuery("sensor_id")
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
	items, err := model.SensorModel.GetItemsByID(id, userId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	var item *SensorItemDto = nil
	if len(items) > 0 {
		item = &SensorItemDto{
			Id:          items[0].Id,
			Name:        items[0].Name,
			Code:        items[0].Code,
			FieldName:   items[0].FieldName,
			GatewayCode: items[0].GatewayCode,
			Longitude:   items[0].Longitude,
			Latitude:    items[0].Latitude,
		}
		for _, v := range items {
			itemProbe := model.ProbeItem{
				Code:  v.ProbeCode,
				Depth: v.ProbeDepth,
			}
			item.Depths = append(item.Depths, &itemProbe)
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": item})
	return
}

//BindSensorByUser ... APP用户绑定Sensor
func BindSensorByUser(c *gin.Context) {
	itemDto := SensorDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Code == "" || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	ok, err := model.SensorModel.BindItemByUser(itemDto.Code, itemDto.UserId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//UnbindSensorByUser ... APP用户解绑
func UnbindSensorByUser(c *gin.Context) {
	itemDto := SensorDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Id <= 0 || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	ok, err := model.SensorModel.UnbindItemByUser(itemDto.Id, itemDto.UserId)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//EditSensorByUser ... APP编辑一条Sensor数据
func EditSensorByUser(c *gin.Context) {
	itemDto := SensorItemDto{}
	bindErr := c.BindJSON(&itemDto)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if itemDto.Id <= 0 || itemDto.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	item := model.IotSensor{
		Id:        itemDto.Id,
		UserId:    itemDto.UserId,
		Longitude: itemDto.Longitude,
		Latitude:  itemDto.Latitude,
	}
	ok, err := model.SensorModel.UpdateItemByUser(&item)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	//更新探针depth
	for _, v := range itemDto.Depths {
		itemProbe := model.IotProbe{
			Code:  v.Code,
			Depth: v.Depth,
		}
		ok, err = model.ProbeModel.UpdateItem(&itemProbe)
		if err != errs.Succ {
			c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//GatherData ... 添加一个传感器采集的数据
//todo  probe的类型也在此接收上传，坐标
func GatherData(c *gin.Context) {
	item := SensorGather{}
	bindErr := c.BindJSON(&item)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	if item.SensorID <= 0 || item.EtlTimeStamp == "" || item.Value <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	sensorItem, errEX := model.SensorModel.GetItemID(item.SensorID)
	if errEX != errs.Succ || sensorItem == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errEX.Code, "msg": errEX.MessageEN, "data": nil})
		return
	}
	value := model.IotValue{
		EtlTimestamp: utils.Convert2Int(item.EtlTimeStamp),
		FieldId:      sensorItem.FieldId,
		Value:        item.Value,
		CreateUid:    0,
		CreateDate:   time.Now(),
		WriteUid:     0,
		WriteDate:    time.Now(),
	}
	ok, err := model.IotValueModel.AddItem(&value)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//GetLineItems ... 获取Sensor信息
func GetLineItems(c *gin.Context) {
	strId, _ := c.GetQuery("sensor_id")
	strStart, _ := c.GetQuery("start")
	strEnd, _ := c.GetQuery("end")
	if strId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	sensorID := utils.Convert2Int(strId)
	startTS, stopTS := utils.Convert2Int64(strStart), utils.Convert2Int64(strEnd)
	items, err := model.IotValueModel.GetLineItems(sensorID, startTS, stopTS)
	if err != errs.Succ {
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}
