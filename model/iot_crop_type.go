package models

import (
	"email-center/utils"
	"errors"
	"github.com/golang/glog"
	"time"
)

var IotCropEx IotCropType

type IotCropType struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Name       string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code       string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	Madv       string    `json:"madv" xorm:"comment('madv') DECIMAL(65,30)"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotCropType) TableName() string {
	return "iot_crop_type"
}

//GetAllItems  ...获取全量数据
func (t IotCropType) GetAllItems() ([]*IotCropType, error) {
	var items []*IotCropType
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsByPage  ...获取全量数据
func (t IotCropType) GetItemsByPage(pageID int) ([]*IotCropType, error) {
	if pageID < 0 {
		return nil, errors.New("invalid-request")
	}
	var items []*IotCropType
	pageCount := 100
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemByID ...根据ID获取对应某条记录
func (t IotCropType) GetItemByID(id int64) (*IotCropType, error) {
	var item *IotCropType
	_, err := utils.GetMysqlClient().ID(id).Get(item)
	if err != nil {
		glog.Errorf("The the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, err
	}
	return item, nil
}

//UpdateItemByID ... 根据数据ID，更新该条数据数据
func (t IotCropType) UpdateItemByID(item *IotCropType) error {
	cols := []string{"name", "code"}
	_, err := utils.GetMysqlClient().Id(item.Id).Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", item, t.TableName(), err)
		return err
	}
	return nil
}
