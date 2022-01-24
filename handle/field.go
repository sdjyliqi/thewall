package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
		StateNowId:    int(utils.FieldIdle),
		CreateUid:     item.UserID,
		CreateDate:    time.Now(),
		WriteDate:     time.Now(),
	}
	dbNode, addErr := model.FieldModel.AddFieldByUser(addItem)
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
		Id:         item.Id,
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
	errEx := model.FieldModel.EditField(addItem)
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
	errEx := model.FieldModel.DelField(fid, uid)
	if errEx != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

//FieldGetItems ... 获取验证码
func FieldGetItems(c *gin.Context) {
	//定义app土地列表中的土地信息结构体
	type fieldView struct {
		Id        int     `json:"id" `
		Name      string  `json:"name"`
		SoilType  string  `json:"soil_type"`
		CropType  string  `json:"crop_type"`
		Status    string  `json:"status"`
		Longitude float32 `json:"longitude"`
		Latitude  float32 `json:"latitude"`
		Threshold string  `json:"threshold"`
	}
	var viewItems []*fieldView
	userID, _ := c.GetQuery("user_id")
	//判断一下userid是否为空
	if userID == "" {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//获取用户id
	id := utils.Convert2Int(userID)
	//根据用户id 查找所属田地信息
	items, errEx := model.FieldModel.GetItemsByUser(id)
	if errEx != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": errEx.Code, "msg": errEx.MessageEN, "data": nil})
		return
	}
	for _, v := range items {
		node := &fieldView{
			Id:        v.IotField.Id,
			Name:      v.IotField.Name,
			SoilType:  GetSoilTypeByID(v.IotField.SoilTypeId),
			CropType:  GetCropTypeByID(v.CropTypeNowId),
			Status:    "ok", //todo 后续需要计算得出改制
			Longitude: v.IotField.Longitude,
			Latitude:  v.IotField.Latitude,
			Threshold: GetReferenceNotice(v.SoilTypeId, v.CropTypeNowId),
		}
		viewItems = append(viewItems, node)
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": viewItems})
	return
}

//FieldProbeLines ... 获取该田地的所有探针列表，默认是当天的数据
func FieldProbeLines(c *gin.Context) {
	type probeLine struct {
		Name         string            `json:"name"`
		Code         string            `json:"code"`
		ProbeType    string            `json:"probe_type"`
		Depth        int               `json:"depth"`
		LastReceived string            `json:"last_received"`
		Kline        []*model.IotValue `json:"kline"`
	}
	type FieldLines struct {
		FieldID         int          `json:"field_id" `
		Name            string       `json:"name"`
		CropType        string       `json:"crop_type"`
		SoilType        string       `json:"soil_type"`
		LastReceiveData string       `json:"last_receive_data"`
		Longitude       float32      `json:"longitude"`
		Latitude        float32      `json:"latitude"`
		ProbeLines      []*probeLine `json:"probe_lines"`
	}

	var klines []*probeLine

	probeValueMap := map[string][]*model.IotValue{}
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
	//
	fmt.Println(uid, fid)

	fieldItem, errEX := model.FieldModel.GetItemByID(fid)
	if errEX != errs.Succ || fieldItem == nil {
		c.JSON(http.StatusOK, gin.H{"code": errEX.Code, "msg": errEX.MessageEN, "data": nil})
		return
	}
	probeItems, errEX := model.ProbeModel.GetProbesByFieldID(fid)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": errEX.Code, "msg": errEX.MessageEN, "data": nil})
		return
	}

	var probeCodes []string
	for _, v := range probeItems {
		fmt.Println("AAAAAAAAAAAAAA", v)
		fmt.Println("v.IotProbe.Code:", v.IotProbe.Code)

		klineNode := &probeLine{
			Name:         strings.Replace(v.IotProbe.Code, "-", "/", 1),
			Code:         v.IotProbe.Code,
			ProbeType:    GetProbeTypeByID(v.IotProbe.ProbeTypeId),
			Depth:        v.IotProbe.Depth,
			LastReceived: v.IotProbe.LastReceived.Format(utils.DayCommonFormat),
			Kline:        nil,
		}
		klines = append(klines, klineNode)
		probeCodes = append(probeCodes, v.IotProbe.Code)
	}
	//获取所有探针的固定时间内的value数据
	fmt.Println("AAAAAAAAAAAAA", probeCodes)
	valueItems, errEX := model.IotValueModel.GetItemsByCodes(probeCodes, 0, time.Now().Unix())
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": errEX.Code, "msg": errEX.MessageEN, "data": nil})
		return
	}
	for _, v := range valueItems {
		probeCode := v.Code
		lineItems, ok := probeValueMap[probeCode]
		if !ok {
			probeValueMap[probeCode] = []*model.IotValue{v}
			continue
		}
		lineItems = append(lineItems, v)
		probeValueMap[probeCode] = lineItems
	}

	fmt.Println("AAAAAAAAAAAAAAAAAAAA", probeValueMap)

	for k, v := range klines {
		klineData, ok := probeValueMap[v.Code]
		if !ok {
			continue
		}
		klines[k].Kline = klineData
	}
	fieldLinesView := FieldLines{
		FieldID:         fid,
		Name:            fieldItem.Name,
		SoilType:        GetSoilTypeByID(fieldItem.SoilTypeId),
		CropType:        GetCropTypeByID(fieldItem.CropTypeNowId),
		LastReceiveData: fieldItem.WriteDate.Format(utils.DayCommonFormat),
		Longitude:       fieldItem.Longitude,
		Latitude:        fieldItem.Latitude,
		ProbeLines:      klines,
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": fieldLinesView})
	return
}
