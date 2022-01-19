package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thewall/errs"
	"thewall/model"
	"thewall/utils"
	"time"
)

type PlantingField struct {
	Id         int `json:"id"` //土地的id
	CropTypeId int `json:"crop_type_id" "`
	DoDate     int `json:"do_date"`
	UserID     int `json:"user_id"` //增加地
}

//FieldPlanting ... 增加农场
func FieldPlanting(c *gin.Context) {
	item := PlantingField{}
	err := c.BindJSON(&item)
	if err != nil || item.DoDate == 0 || item.CropTypeId <= 0 {
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
	//继续判断该土地是否处于结束种植状态或者没有找到记录
	addItem := &model.IotPlant{
		FieldId:      item.Id,
		CropTypeId:   item.CropTypeId,
		PlantingDate: time.Now(),
		StateId:      int(utils.FieldPlanting),
		WriteUid:     item.UserID,
		WriteDate:    time.Now(),
	}
	dbNode, addErr := model.PlantModel.Planting(addItem)
	if addErr != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": chkErr.Code, "msg": chkErr.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": dbNode})
	return
}
