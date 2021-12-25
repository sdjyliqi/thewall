package model

import (
	"errors"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var ProjectModel IotProject

type IotProject struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Name       string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code       string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	UserId     int       `json:"user_id" xorm:"comment('User') index INT(11)"`
	PartnerId  int       `json:"partner_id" xorm:"comment('Partner') index INT(11)"`
	CompanyId  int       `json:"company_id" xorm:"comment('Company') index INT(11)"`
	AddressId  int       `json:"address_id" xorm:"comment('Address') index INT(11)"`
	SoilTypeId int       `json:"soil_type_id" xorm:"comment('soil_type_id') index INT(11)"`
	CropTypeId int       `json:"crop_type_id" xorm:"comment('crop_type_id') index INT(11)"`
	StartDate  time.Time `json:"start_date" xorm:"comment('Start Date') DATE"`
	EndDate    time.Time `json:"end_date" xorm:"comment('End Date') DATE"`
	Amount     string    `json:"amount" xorm:"comment('Amount') DECIMAL(65,30)"`
	State      string    `json:"state" xorm:"not null comment('Status') LONGTEXT"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotProject) TableName() string {
	return "iot_project"
}

//GetAllItems  ...获取全量数据
func (t IotProject) GetAllItems() ([]*IotProject, error) {
	var items []*IotProject
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsByPage  ...分页获取全量数据
func (t IotProject) GetItemsByPage(pageID int) ([]*IotProject, error) {
	if pageID < 0 {
		return nil, errors.New("invalid-request")
	}
	var items []*IotProject
	pageCount := 100
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemByID ...根据ID获取对应某条记录
func (t IotProject) GetItemByID(id int64) (*IotProject, error) {
	var item *IotProject
	_, err := utils.GetMysqlClient().ID(id).Get(item)
	if err != nil {
		glog.Errorf("The the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, err
	}
	return item, nil
}

//UpdateItemByID ... 根据数据ID，更新该条数据数据
func (t IotProject) UpdateItemByID(item *IotProject) (int64, error) {
	rows, err := utils.GetMysqlClient().Id(item.Id).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", item, t.TableName(), err)
		return 0, err
	}
	return rows, nil
}

//AddItem ... 添加一条数据
func (t IotProject) AddItem(item *IotProject) (bool, errs.ErrInfo) {
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}
