package model

import (
	"fmt"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var IotFieldEx IotField

//IotField ... 土地基本信息
type IotField struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name          string    `json:"name" xorm:"comment('Name') VARCHAR(64)"`
	NameCn        string    `json:"name_cn" xorm:"comment('中文名') VARCHAR(64)"`
	UserId        int       `json:"user_id" xorm:"comment('User') index INT(11)"`
	AddressId     int       `json:"address_id" xorm:"comment('Address') INT(11)"`
	Longitude     float32   `json:"longitude" xorm:"FLOAT"`
	Latitude      float32   `json:"latitude" xorm:"FLOAT"`
	Area          float32   `json:"area" xorm:"FLOAT"`
	SoilTypeId    int       `json:"soil_type_id" xorm:"comment('soil_type_id') INT(11)"`
	CropTypeNowId int       `json:"crop_type_now_id" xorm:"comment('crop_type_now_id') INT(11)"`
	StateNowId    int       `json:"state_now_id" xorm:"comment('state_now_id') INT(11)"`
	CreateUid     int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate    time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid      int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate     time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

//FieldExtend ...先临时做一个多表查询
type FieldExtend struct {
	IotField `xorm:"extends"`
	//IotCropType `xorm:"extends"`
	IotSoilType `xorm:"extends"`
}

func (t IotField) TableName() string {
	return "iot_field"
}

//AddFieldByUser ... 用户增加土地
func (t IotField) AddFieldByUser(item *IotField) errs.ErrInfo {
	//Todo，后续需要去重，条件需要沟通
	item.CreateDate = time.Now()
	item.WriteDate = time.Now()
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return errs.ErrDBInsert
	}
	return errs.Succ
}

//EditField ... 用户增加土地
func (t IotField) EditField(item *IotField) errs.ErrInfo {
	//Todo，后续需要去重，条件需要沟通
	cols := []string{"name", "name_cn", "longitude", "latitude", "area", "soil_type_id", "write_date"}
	item.CreateDate = time.Now()
	item.WriteDate = time.Now()
	_, err := utils.GetMysqlClient().ID(item.Id).Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Insert the item %+v to table %s failed,err:%+v", *item, t.TableName(), err)
		return errs.ErrDBInsert
	}
	return errs.Succ
}

//DelField ... 删除土地信息
func (t IotField) DelField(fieldID, userID int) errs.ErrInfo {
	condition := fmt.Sprintf("user_id=%d", userID)
	_, err := utils.GetMysqlClient().ID(fieldID).Where(condition).Delete(IotField{})
	if err != nil {
		glog.Errorf("Delete the item by id %d and userID %+v to table %s failed,err:%+v", fieldID, userID, t.TableName(), err)
		return errs.ErrDBDel
	}
	return errs.Succ
}

func (t IotField) GetItemsByUser(userID int) ([]*FieldExtend, errs.ErrInfo) {
	var items []*FieldExtend
	join := fmt.Sprintf("%s.soil_type_id=%s.id", t.TableName(), IotSoilTypeModel.TableName())
	condition := fmt.Sprintf("%s.user_id=%d", t.TableName(), userID)
	err := utils.GetMysqlClient().Table(t.TableName()).Join("LEFT", IotSoilTypeModel, join).Where(condition).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}
