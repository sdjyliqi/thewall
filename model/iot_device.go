package model

import (
	"errors"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var DeviceModel IotDevice

type IotDevice struct {
	Id          int       `json:"id" xorm:"not null pk INT(11)"`
	UserId      int       `json:"user_id" xorm:"comment('User') index INT(11)"`
	PartnerId   int       `json:"partner_id" xorm:"comment('Partner') index INT(11)"`
	CompanyId   int       `json:"company_id" xorm:"comment('Company') index INT(11)"`
	GatewayId   int       `json:"gateway_id" xorm:"comment('Gateway') index INT(11)"`
	ProjectId   int       `json:"project_id" xorm:"comment('Project') index INT(11)"`
	Name        string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code        string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	Latitude    string    `json:"latitude" xorm:"comment('Latitude') DECIMAL(65,30)"`
	Longitude   string    `json:"longitude" xorm:"comment('Longitude') DECIMAL(65,30)"`
	Interval    int       `json:"interval" xorm:"comment('Interval') INT(11)"`
	LastEtlTime time.Time `json:"last_etl_time" xorm:"comment('Last Check Time') DATETIME"`
	CreateUid   int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate  time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid    int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate   time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotDevice) TableName() string {
	return "iot_device"
}

//GetAllItems  ...获取全量数据
func (t IotDevice) GetAllItems() ([]*IotDevice, error) {
	var items []*IotDevice
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsByPage  ...分页获取全量数据
func (t IotDevice) GetItemsByPage(pageID int) ([]*IotDevice, error) {
	if pageID < 0 {
		return nil, errors.New("invalid-request")
	}
	var items []*IotDevice
	pageCount := 100
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemByID ...根据ID获取对应某条记录
func (t IotDevice) GetItemByID(id int64) (*IotDevice, error) {
	var item *IotDevice
	_, err := utils.GetMysqlClient().ID(id).Get(item)
	if err != nil {
		glog.Errorf("The the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, err
	}
	return item, nil
}

//UpdateItemByID ... 根据数据ID，更新该条数据数据
func (t IotDevice) UpdateItemByID(item *IotDevice) (int64, error) {
	rows, err := utils.GetMysqlClient().Id(item.Id).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", item, t.TableName(), err)
		return 0, err
	}
	return rows, nil
}

//AddItem ... 添加一条数据
func (t IotDevice) AddItem(item *IotDevice) (bool, errs.ErrInfo) {
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}
