package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thewall/errs"
	"thewall/model"
	"thewall/utils"
	"time"
)

type AddField struct {
	Id         int     `json:"id"` //新增的时候，该字段为空
	Name       string  `json:"name" `
	NameCn     string  `json:"name_cn" `
	SoilTypeId int     `json:"soil_type_id" ` //土地类型
	Country    string  `json:"country" `      //国家
	Longitude  float32 `json:"longitude" `    //经度
	Latitude   float32 `json:"latitude" `     //维度
	Area       float32 `json:"area" `         //面积
	UserID     int     `json:"user_id"`       //增加地
	Sensors    []int   `json:"sensors"`       //增加地
}

//FieldAdd ... 增加农场
func FieldAdd(c *gin.Context) {
	item := AddField{}
	err := c.BindJSON(&item)
	if err != nil || (item.UserID < 1 || item.Name == "" || item.Latitude < 0 || item.Longitude < 0 || item.SoilTypeId < 0) {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//首先判断userid 是否合法
	existed, chkErr := model.UCModel.ChkUserExisted(item.UserID)
	if chkErr != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	//如果某用户不存在，直接异常返回即可
	if !existed {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrUCNoUser.Code, "msg": errs.ErrUCNoUser.MessageEN, "data": nil})
	}
	addItem := &model.IotField{
		Name:          item.Name,
		NameCn:        item.NameCn,
		UserId:        item.UserID,
		Longitude:     item.Longitude,
		Latitude:      item.Latitude,
		Area:          item.Area,
		SoilTypeId:    item.SoilTypeId,
		CropTypeNowId: 0,
		StateNowId:    0,
		CreateUid:     item.UserID,
		CreateDate:    time.Now(),
		WriteDate:     time.Now(),
	}
	dbNode, addErr := model.IotFieldEx.AddFieldByUser(addItem)
	if addErr != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	errEX := model.SensorModel.SensorBindFiled(item.Sensors, dbNode.Id, item.UserID)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

//FieldEdit ...修改农场信息
func FieldEdit(c *gin.Context) {
	item := AddField{}
	err := c.BindJSON(&item)
	if err != nil || (item.Id < 0 || item.UserID < 1 || item.Name == "" || item.Latitude < 0 || item.Longitude < 0 || item.SoilTypeId < 0) {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//首先判断userid 是否合法
	existed, chkErr := model.UCModel.ChkUserExisted(item.UserID)
	if chkErr != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	//如果某用户不存在，直接异常返回即可
	if !existed {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrUCNoUser.Code, "msg": errs.ErrUCNoUser.MessageEN, "data": nil})
	}
	addItem := &model.IotField{
		Name:       item.Name,
		NameCn:     item.NameCn,
		UserId:     item.UserID,
		Longitude:  item.Longitude,
		Latitude:   item.Latitude,
		Area:       item.Area,
		SoilTypeId: item.SoilTypeId,
		CreateUid:  item.UserID,
		WriteDate:  time.Now(),
	}
	errEx := model.IotFieldEx.EditField(addItem)
	if errEx != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	errEX := model.SensorModel.SensorBindFiled(item.Sensors, item.Id, item.UserID)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

//FieldDel ... 获取验证码
func FieldDel(c *gin.Context) {
	strUID, _ := c.GetQuery("user_id")
	strFID, _ := c.GetQuery("field_id")
	//判断一下userid是否为空
	if strUID == "" || strFID == "" {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//获取用户id
	uid := utils.Convert2Int(strUID)
	fid := utils.Convert2Int(strFID)
	errEx := model.IotFieldEx.DelField(fid, uid)
	if errEx != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

//FieldGetItems ... 获取验证码
func FieldGetItems(c *gin.Context) {
	userID, _ := c.GetQuery("user_id")
	//判断一下userid是否为空
	if userID == "" {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//获取用户id
	id := utils.Convert2Int(userID)
	//根据用户id 查找所属田地信息
	items, errEx := model.IotFieldEx.GetItemsByUser(id)
	if errEx != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}
