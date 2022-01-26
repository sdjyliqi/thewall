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
		StateNowId:    int(utils.FieldPlanting),
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
	fieldItem := &IotField{
		Id:         item.FieldId,
		StateNowId: int(utils.FieldHarvest),
		WriteDate:  time.Now(),
	}
	cols := []string{"write_date", "state_now_id"}
	errEX := FieldModel.EditFieldByID(fieldItem, cols)
	if errEX != errs.Succ {
		return nil, errEX
	}
	cols = []string{"harvest_date", "state_id", "write_date"}
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

//Weigh  ...开始称重
func (t IotPlant) Weigh(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	fieldItem := &IotField{
		Id:         item.FieldId,
		StateNowId: int(utils.FieldWeight),
		WriteDate:  time.Now(),
	}
	cols := []string{"write_date", "state_now_id"}
	errEX := FieldModel.EditFieldByID(fieldItem, cols)
	if errEX != errs.Succ {
		return nil, errEX
	}
	cols = []string{"amount", "weigh_date", "state_id", "write_date"}
	_, err := utils.GetMysqlClient().
		Where("field_id=?", item.FieldId).
		And("state_id!=?", int(utils.FieldFinish)).
		Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBUpdate
	}
	return item, errs.Succ
}

//Ended  ...终止
func (t IotPlant) Ended(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	//更新土地的当前状态
	fieldItem := &IotField{
		Id:         item.FieldId,
		StateNowId: int(utils.FieldFinish),
		WriteDate:  time.Now(),
	}
	cols := []string{"write_date", "state_now_id"}
	errEX := FieldModel.EditFieldByID(fieldItem, cols)
	if errEX != errs.Succ {
		return nil, errEX
	}
	//更详当前其次农作物的状态
	cols = []string{"state_id", "write_date"}
	uItem := &IotField{
		Id:         item.FieldId,
		StateNowId: int(utils.FieldFinish),
		WriteDate:  time.Now(),
	}
	errEX = FieldModel.EditFieldByID(uItem, cols)
	if errEX != errs.Succ {
		return nil, errEX
	}
	cols = []string{"state_id", "write_date"}
	_, err := utils.GetMysqlClient().
		Where("field_id=?", item.FieldId).
		And("state_id!=?", int(utils.FieldIdle)).
		Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBUpdate
	}
	return item, errs.Succ
}

//GetHistoryPlant  ...获取
func (t IotPlant) GetHistoryPlant(fieldID int) ([]*IotPlant, errs.ErrInfo) {
	var items []*IotPlant
	err := utils.GetMysqlClient().
		Where("field_id=?", fieldID).
		And("state_id =?", int(utils.FieldIdle)).
		OrderBy("planting_date").
		Find(&items)
	if err != nil {
		glog.Errorf("Find the items by field %+v from table %s failed,err:%+v", fieldID, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}
