package model

import (
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var PlantModel IotPlant

type IotPlant struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	FieldId      int       `json:"field_id" xorm:"comment('field_id') INT(11)"`
	CropTypeId   int       `json:"crop_type_id" xorm:"comment('crop_type_id') INT(11)"`
	PlantingDate time.Time `json:"planting_date" xorm:"comment('Start Date') DATE"`
	WeighDate    time.Time `json:"weigh_date" xorm:"DATE"`
	HarvestDate  time.Time `json:"harvest_date" xorm:"comment('End Date') DATE"`
	Amount       float32   `json:"amount" xorm:"comment('Amount') FLOAT(11,2)"`
	StateId      int       `json:"state_id" xorm:"comment('种植周期阶段') INT(11)"`
	WriteUid     int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate    time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotPlant) TableName() string {
	return "iot_plant"
}

//Planting  ...开始种植,记得在田地中修改当前农作物属性
func (t IotPlant) Planting(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBInsert
	}
	fieldItem := &IotField{
		Id:            item.FieldId,
		CropTypeNowId: item.CropTypeId,
		WriteDate:     time.Now(),
	}
	cols := []string{"crop_type_now_id", "write_date"}
	errEX := FieldModel.EditFieldByID(fieldItem, cols)
	if errEX != errs.Succ {
		return nil, errEX
	}
	return item, errs.Succ
}

//Harvest  ...开始种植
func (t IotPlant) Harvest(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	cols := []string{"harvest_date", "state_id", "write_date"}
	_, err := utils.GetMysqlClient().
		Where("field_id=?", item.FieldId).
		And("state_id=?", int(utils.FieldPlanting)).
		Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBInsert
	}
	return item, errs.Succ
}

//Weigh  ...开始称重
func (t IotPlant) Weigh(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	cols := []string{"weigh_date", "state_id", "write_date"}
	_, err := utils.GetMysqlClient().
		Where("field_id=?", item.FieldId).
		And("state_id=?", int(utils.FieldHarvest)).
		Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBInsert
	}
	return item, errs.Succ
}

//Ended  ...终止
func (t IotPlant) Ended(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	cols := []string{"crop_type_now_id", "write_date"}
	fieldItem := &IotField{
		Id:            item.FieldId,
		CropTypeNowId: utils.NoPlantsCropType,
		WriteDate:     time.Now(),
	}
	errEX := FieldModel.EditFieldByID(fieldItem, cols)
	if errEX != errs.Succ {
		return nil, errEX
	}
	cols = []string{"state_id", "write_date"}
	_, err := utils.GetMysqlClient().
		Where("field_id=?", item.FieldId).
		And("state_id!=?", int(utils.FieldFinish)).
		Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBInsert
	}
	return item, errs.Succ
}
