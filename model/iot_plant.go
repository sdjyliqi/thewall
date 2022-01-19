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
	WeightDate   time.Time `json:"weight_date" xorm:"DATE"`
	HarvestDate  time.Time `json:"harvest_date" xorm:"comment('End Date') DATE"`
	Amount       float32   `json:"amount" xorm:"comment('Amount') FLOAT(11,2)"`
	StateId      int       `json:"state_id" xorm:"comment('种植周期阶段') INT(11)"`
	WriteUid     int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate    time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotPlant) TableName() string {
	return "iot_plant"
}

//Planting  ...开始种植
func (t IotPlant) Planting(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBInsert
	}
	return item, errs.Succ
}

//Harvest  ...开始种植
func (t IotPlant) Harvest(item *IotPlant) (*IotPlant, errs.ErrInfo) {
	cols := []string{"harvest_date", "state_id", "write_date"}
	_, err := utils.GetMysqlClient().ID(item.Id).Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return nil, errs.ErrDBInsert
	}
	return item, errs.Succ
}
