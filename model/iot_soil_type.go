package model

import (
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var IotSoilTypeModel IotSoilType

type IotSoilType struct {
	Id                  int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name                string    `json:"name" xorm:"comment('Name') VARCHAR(32)"`
	NameCn              string    `json:"name_cn" xorm:"comment('中文名') VARCHAR(32)"`
	Code                string    `json:"code" xorm:"comment('Code') VARCHAR(32)"`
	FieldCapacity       float32   `json:"field_capacity" xorm:"comment('Field Capacity') FLOAT(11,4)"`
	TotalAvailableWater float32   `json:"total_available_water" xorm:"comment('Total Available Water') FLOAT(11,4)"`
	Pwp                 float32   `json:"pwp" xorm:"comment('PWP') FLOAT(11,4)"`
	CreateUid           int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate          time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid            int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate           time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotSoilType) TableName() string {
	return "iot_soil_type"
}

//GetAllItems  ...获取全量数据
func (t IotSoilType) GetAllItems() ([]*IotSoilType, error) {
	var items []*IotSoilType
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//UpdateItemByID ... 根据数据ID，更新该条数据数据
func (t IotSoilType) UpdateItemByID(item *IotSoilType) error {
	cols := []string{"name", "code", "field_capacity", "total_available_water", "pwp", "write_uid", "write_date"}
	_, err := utils.GetMysqlClient().ID(item.Id).Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", item, t.TableName(), err)
		return err
	}
	return nil
}

//AddItem ... 添加一条数据
func (t IotSoilType) AddItem(item *IotSoilType) (bool, errs.ErrInfo) {
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}
