package model

import (
	"errors"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var SensorTypeModel IotSensorType

type IotSensorType struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Name       string    `json:"name" xorm:"comment('Name') VARCHAR(32)"`
	Code       string    `json:"code" xorm:"comment('Code') VARCHAR(32)"`
	NameCn     string    `json:"name_cn" xorm:"comment('中文名') VARCHAR(32)"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotSensorType) TableName() string {
	return "iot_sensor_type"
}

//GetAllItems  ...获取全量数据
func (t IotSensorType) GetAllItems() ([]*IotSensorType, error) {
	var items []*IotSensorType
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsByPage  ...分页获取全量数据
func (t IotSensorType) GetItemsByPage(pageID int) ([]*IotSensorType, error) {
	if pageID < 0 {
		return nil, errors.New("invalid-request")
	}
	var items []*IotSensorType
	pageCount := 100
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemByID ...根据ID获取对应某条记录
func (t IotSensorType) GetItemByID(id int64) (*IotSensorType, error) {
	var item *IotSensorType
	_, err := utils.GetMysqlClient().ID(id).Get(item)
	if err != nil {
		glog.Errorf("The the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, err
	}
	return item, nil
}

//UpdateItemByID ... 根据数据ID，更新该条数据数据
func (t IotSensorType) UpdateItemByID(item *IotSensorType) (int64, error) {
	rows, err := utils.GetMysqlClient().Id(item.Id).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", item, t.TableName(), err)
		return 0, err
	}
	return rows, nil
}

//AddItem ... 添加一条数据
func (t IotSensorType) AddItem(item *IotSensorType) (bool, errs.ErrInfo) {
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}
