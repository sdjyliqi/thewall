package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thewall/errs"
	"thewall/model"
	"thewall/utils"
	"time"
)

type PlantingField struct {
	id         int     `json:"id"`       //某条记录，种植时候，该值为0
	FieldID    int     `json:"field_id"` //土地的id
	CropTypeId int     `json:"crop_type_id" "`
	UserID     int     `json:"user_id"` //增加地
	Amount     float32 `json:"amount"`  //增加地
}

//FieldPlanting ... 增加农场
func FieldPlanting(c *gin.Context) {
	item := PlantingField{}
	err := c.BindJSON(&item)
	if err != nil || item.CropTypeId <= 0 || item.FieldID <= 0 {
		fmt.Println("=============AAAAAAAAAAAA===============")
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//首先判断userid 是否合法
	existed, chkErr := model.UCModel.ChkUserExisted(item.UserID)
	if chkErr != errs.Succ {
		fmt.Println("==============bbbbbbbbbbbbb================")
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	//如果某用户不存在，直接异常返回即可
	if !existed {
		fmt.Println("==============ccc===============")
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrUCNoUser.Code, "msg": errs.ErrUCNoUser.MessageEN, "data": nil})
	}
	fieldInfo, errEX := model.FieldModel.GetItemByID(item.FieldID)
	if errEX != errs.Succ {
		fmt.Println("=============eeee==============")
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	//如果根据id查询不到土地信息，或者土地的状态不是空置状态，直接返回异常。
	if fieldInfo != nil && fieldInfo.StateNowId != int(utils.FieldIdle) {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrPeriodNoPlanting.Code, "msg": errs.ErrPeriodNoPlanting.MessageEN, "data": nil})
		return
	}
	//继续判断该土地是否处于结束种植状态或者没有找到记录
	addItem := &model.IotPlant{
		FieldId:      item.FieldID,
		CropTypeId:   item.CropTypeId,
		PlantingDate: time.Now(),
		StateId:      int(utils.FieldPlanting),
		WriteUid:     item.UserID,
		WriteDate:    time.Now(),
	}
	dbNode, errEX := model.PlantModel.Planting(addItem)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": dbNode})
	return
}

//FieldHarvest ... 收割
func FieldHarvest(c *gin.Context) {
	item := PlantingField{}
	err := c.BindJSON(&item)
	if err != nil || item.FieldID <= 0 || item.UserID <= 0 {
		fmt.Println("aaaaaaaaaaaaa")
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//首先判断userid 是否合法
	existed, chkErr := model.UCModel.ChkUserExisted(item.UserID)
	if chkErr != errs.Succ {
		fmt.Println("BBBBBBBBBBBBBBB")
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	//如果某用户不存在，直接异常返回即可
	if !existed {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrUCNoUser.Code, "msg": errs.ErrUCNoUser.MessageEN, "data": nil})
	}
	//如果根据id查询不到土地信息，或者土地的状态不是空置状态，直接返回异常。
	fieldInfo, errEX := model.FieldModel.GetItemByID(item.FieldID)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	if fieldInfo == nil {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	fmt.Println("++++++++++++++++", fieldInfo.StateNowId)
	if fieldInfo.StateNowId != int(utils.FieldPlanting) {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrPeriodNoHarvest.Code, "msg": errs.ErrPeriodNoHarvest.MessageEN, "data": nil})
		return
	}
	//继续判断该土地是否处于结束种植状态或者没有找到记录
	addItem := &model.IotPlant{
		FieldId:     item.FieldID,
		HarvestDate: time.Now(),
		StateId:     int(utils.FieldHarvest),
		WriteDate:   time.Now(),
	}
	_, errEX = model.PlantModel.Harvest(addItem)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

//FieldWeigh ... 收割
func FieldWeigh(c *gin.Context) {
	item := PlantingField{}
	err := c.BindJSON(&item)
	if err != nil || item.UserID <= 0 || item.FieldID <= 0 || item.Amount < 0 {
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

	//如果根据id查询不到土地信息，或者土地的状态不是空置状态，直接返回异常。
	fieldInfo, errEX := model.FieldModel.GetItemByID(item.FieldID)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	if fieldInfo == nil {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	fmt.Println("++++++++++++++++", fieldInfo.StateNowId)
	if fieldInfo.StateNowId != int(utils.FieldHarvest) {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrPeriodNoWeigh.Code, "msg": errs.ErrPeriodNoWeigh.MessageEN, "data": nil})
		return
	}
	//继续判断该土地是否处于结束种植状态或者没有找到记录
	uItem := &model.IotPlant{
		FieldId:   item.FieldID,
		StateId:   int(utils.FieldWeight),
		Amount:    item.Amount,
		WeighDate: time.Now(),
		WriteDate: time.Now(),
	}
	_, errEX = model.PlantModel.Weigh(uItem)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

//FieldEnded ... 终止
func FieldEnded(c *gin.Context) {
	item := PlantingField{}
	err := c.BindJSON(&item)
	if err != nil || item.UserID <= 0 || item.FieldID <= 0 {
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
	//如果根据id查询不到土地信息，或者土地的状态不是空置状态，直接返回异常。
	fieldInfo, errEX := model.FieldModel.GetItemByID(item.FieldID)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	if fieldInfo == nil {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}

	//继续判断该土地是否处于结束种植状态或者没有找到记录
	editItem := &model.IotPlant{
		FieldId:   item.FieldID,
		StateId:   int(utils.FieldIdle),
		WriteDate: time.Now(),
	}
	_, errEX = model.PlantModel.Ended(editItem)
	if errEX != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}
