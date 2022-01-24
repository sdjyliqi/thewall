package model

import (
	"fmt"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var GatewayModel IotGateway

type IotGateway struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId     int       `json:"user_id" xorm:"comment('User') INT(11)"`
	FieldId    int       `json:"field_id" xorm:"comment('field_id') INT(11)"`
	Name       string    `json:"name" xorm:"comment('Name') VARCHAR(32)"`
	Code       string    `json:"code" xorm:"comment('2000年以后的16进制数') VARCHAR(16)"`
	Longitude  float32   `json:"longitude" xorm:"FLOAT"`
	Latitude   float32   `json:"latitude" xorm:"FLOAT"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotGateway) TableName() string {
	return "iot_gateway"
}

//GetItemsByPage  ...后台分页获取全量数据
func (t IotGateway) GetItemsByPage(pageID int) ([]*IotGateway, errs.ErrInfo) {
	if pageID < 0 {
		return nil, errs.ErrBadRequest
	}
	var items []*IotGateway
	pageCount := 10
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("Get the items from %s failed,err:%+v", t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

//GetItemByID ...后台根据ID获取对应某条记录
func (t IotGateway) GetItemByID(id int) (*IotGateway, errs.ErrInfo) {
	if id <= 0 {
		return nil, errs.ErrBadRequest
	}
	item := &IotGateway{
		Id: id,
	}
	_, err := utils.GetMysqlClient().Get(item)
	if err != nil {
		glog.Errorf("Get the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return item, errs.Succ
}

//GetItemByUser ...APP根据userID获取对应某条记录
func (t IotGateway) GetItemByUser(id, userID int) (*IotGateway, errs.ErrInfo) {
	if userID <= 0 {
		return nil, errs.ErrBadRequest
	}
	item := new(IotGateway)
	condition := fmt.Sprintf("user_id=%d", userID)
	_, err := utils.GetMysqlClient().ID(id).And(condition).Get(item)
	if err != nil {
		glog.Errorf("Get the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return item, errs.Succ
}

//UpdateItemByID ... 后台根据数据ID，更新该条数据数据
func (t IotGateway) UpdateItemByID(item *IotGateway) (bool, errs.ErrInfo) {
	if item.Id <= 0 || item.WriteUid <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"write_uid", "write_date"}
	updateItem := &IotGateway{
		WriteUid:  item.WriteUid,
		WriteDate: time.Now(),
	}
	if item.Longitude > 0 {
		cols = append(cols, "longitude")
		updateItem.Longitude = item.Longitude
	}
	if item.Latitude > 0 {
		cols = append(cols, "latitude")
		updateItem.Latitude = item.Latitude
	}
	rows, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).Update(updateItem)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}

//UpdateItemByUser ... APP根据数据userID，更新该条数据数据
func (t IotGateway) UpdateItemByUser(item *IotGateway) (bool, errs.ErrInfo) {
	if item.Id <= 0 || item.UserId <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"write_uid", "write_date"}
	updateItem := &IotGateway{
		WriteUid:  item.UserId,
		WriteDate: time.Now(),
	}
	if item.Longitude > 0 {
		cols = append(cols, "longitude")
		updateItem.Longitude = item.Longitude
	}
	if item.Latitude > 0 {
		cols = append(cols, "latitude")
		updateItem.Latitude = item.Latitude
	}
	condition := fmt.Sprintf("user_id=%d", item.UserId)
	rows, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).And(condition).Update(updateItem)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}

//BindItemByUser ... APP绑定用户ID
func (t IotGateway) BindItemByUser(code string, userID int) (bool, errs.ErrInfo) {
	if code == "" || userID <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"user_id", "write_uid", "write_date"}
	updateItem := &IotGateway{
		UserId:    userID,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	condition := fmt.Sprintf("code='%s'", code)
	rows, err := utils.GetMysqlClient().Cols(cols...).Where(condition).Update(updateItem)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}

//UnbindItemByUser ... APP用户解绑
func (t IotGateway) UnbindItemByUser(id, userID int) (bool, errs.ErrInfo) {
	if id <= 0 || userID <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"user_id", "longitude", "latitude", "write_uid", "write_date"}
	updateItem := &IotGateway{
		UserId:    0,
		Longitude: 0,
		Latitude:  0,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	rows, err := utils.GetMysqlClient().Cols(cols...).ID(id).And("user_id=?", userID).Update(updateItem)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}

//AddItem ... 后台添加一条数据
func (t IotGateway) AddItem(item *IotGateway) (bool, errs.ErrInfo) {
	item.CreateDate = time.Now()
	item.WriteDate = time.Now()
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}

//GetItemsByUser ... APP获取用户绑定的Gateway列表数据
func (t IotGateway) GetItemsByUser(userID int) ([]*IotGateway, errs.ErrInfo) {
	if userID < 0 {
		return nil, errs.ErrBadRequest
	}
	var items []*IotGateway
	condition := fmt.Sprintf("user_id=%d", userID)
	err := utils.GetMysqlClient().Where(condition).Find(&items)
	if err != nil {
		glog.Errorf("Get the items from %s failed,err:%+v", t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}
