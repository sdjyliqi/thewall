package model

import (
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

type IotSoilType struct {
	Id                  int       `json:"id" xorm:"not null pk INT(11)"`
	Name                string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code                string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	FieldCapacity       string    `json:"field_capacity" xorm:"comment('Field Capacity') DECIMAL(65,30)"`
	TotalAvailableWater string    `json:"total_available_water" xorm:"comment('Total Available Water') DECIMAL(65,30)"`
	Pwp                 string    `json:"pwp" xorm:"comment('PWP') DECIMAL(65,30)"`
	CreateUid           int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate          time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid            int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
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
