package model

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var SensorModel IotSensor

type IotSensor struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name            string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code            string    `json:"code" xorm:"comment('2000年以后的16进制数') VARCHAR(16)"`
	FieldId         int       `json:"field_id" xorm:"comment('农场id') INT(11)"`
	UserId          int       `json:"user_id" xorm:"INT(11)"`
	GatewayId       int       `json:"gateway_id" xorm:"comment('gateway_id') INT(11)"`
	SensorTypeId    int       `json:"sensor_type_id" xorm:"comment('Sensor Type') INT(11)"`
	Depth           int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	LastRecivedTime time.Time `json:"last_recived_time" xorm:"comment('最后上传数据的时间') DATETIME"`
	CreateUid       int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate      time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid        int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate       time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

//SensorExtend ...多表查询
type SensorExtend struct {
	//IotSensor  `xorm:"extends"`
	//IotField   `xorm:"extends"`
	//IotGateway `xorm:"extends"`
	Id          int    `json:"id"`
	Depth       int    `json:"depth"`
	FieldName   string `json:"field_name"`
	GatewayCode string `json:"gateway_code"`
}

func (t IotSensor) TableName() string {
	return "iot_sensor"
}

//GetAllItems  ...获取全量数据
func (t IotSensor) GetAllItems() ([]*IotSensor, error) {
	var items []*IotSensor
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsByPage  ...分页获取全量数据
func (t IotSensor) GetItemsByPage(pageID int) ([]*IotSensor, error) {
	if pageID < 0 {
		return nil, errors.New("invalid-request")
	}
	var items []*IotSensor
	pageCount := 100
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemByID ...根据ID获取对应某条记录
func (t IotSensor) GetItemByID(id int) (*SensorExtend, errs.ErrInfo) {
	if id <= 0 {
		return nil, errs.ErrBadRequest
	}
	var items []*SensorExtend
	joinSelect := fmt.Sprintf("%s.id,%s.depth,%s.name as field_name,%s.code as gateway_code", t.TableName(), t.TableName(), FieldModel.TableName(), GatewayModel.TableName())
	joinField := fmt.Sprintf("%s.field_id=%s.id", t.TableName(), FieldModel.TableName())
	joinGateway := fmt.Sprintf("%s.gateway_id=%s.id", t.TableName(), GatewayModel.TableName())
	condition := fmt.Sprintf("%s.id=%d", t.TableName(), id)
	err := utils.GetMysqlClient().Table(t.TableName()).Select(joinSelect).
		Join("LEFT", FieldModel.TableName(), joinField).
		Join("LEFT", GatewayModel.TableName(), joinGateway).
		Where(condition).Find(&items)
	if err != nil {
		glog.Errorf("Get the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	var item *SensorExtend = nil
	if len(items) > 0 {
		item = items[0]
	}
	return item, errs.Succ
}

//GetItemsByField ...获取当前某农场绑定的传感器列表
func (t IotSensor) GetItemsByField(fieldID int) ([]*IotSensor, errs.ErrInfo) {
	var items []*IotSensor
	err := utils.GetMysqlClient().Where("field_id=?", fieldID).Find(&items)
	if err != nil {
		glog.Errorf("Get items by field %d from %s failed,err:%+v", fieldID, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

//GetItemsByGateway ...获取绑定Gateway的传感器列表
func (t IotSensor) GetItemsByGateway(gatewayId int) ([]*IotSensor, errs.ErrInfo) {
	var items []*IotSensor
	err := utils.GetMysqlClient().Where("gateway_id=?", gatewayId).Find(&items)
	if err != nil {
		glog.Errorf("Get items by field %d from %s failed,err:%+v", gatewayId, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

//GetItemsByUser ...获取当前某农场绑定的传感器列表
func (t IotSensor) GetItemsByUser(userID int) ([]*IotSensor, errs.ErrInfo) {
	var items []*IotSensor
	err := utils.GetMysqlClient().Where("user_id=?", userID).Find(&items)
	if err != nil {
		glog.Errorf("Get items by user %d from %s failed,err:%+v", userID, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

func (t IotSensor) UpdateItemByUser(item *IotSensor) (bool, errs.ErrInfo) {
	if item.Id <= 0 || item.UserId <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"depth", "write_uid", "write_date"}
	updateItem := &IotSensor{
		Depth:     item.Depth,
		WriteUid:  item.UserId,
		WriteDate: time.Now(),
	}
	condition := fmt.Sprintf("user_id=%d", item.UserId)
	rows, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).And(condition).Update(updateItem)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}

//AddItem ... 添加一条数据
func (t IotSensor) AddItem(item *IotSensor) (bool, errs.ErrInfo) {
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}

//SensorBindFiled ...重新绑定某地的传感器
func (t IotSensor) SensorBindFiled(sensorIDs []int, field, userID int) errs.ErrInfo {
	//先根据field的sensor实施重置，接触绑定
	unbindSensor := &IotSensor{
		FieldId:   0,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	cols := []string{"field_id", "write_uid", "write_date"}
	_, err := utils.GetMysqlClient().Cols(cols...).Where("field_id=?", field).And("user_id=?", userID).Update(unbindSensor)
	if err != nil {
		glog.Errorf("Update items by field_id %d from %s failed,err:%+v", field, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	//重新针对sensor的id 实施绑定
	bindSensor := &IotSensor{
		FieldId:   field,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	_, err = utils.GetMysqlClient().Cols(cols...).In("id", sensorIDs).And("user_id=?", userID).Update(bindSensor)
	if err != nil {
		glog.Errorf("Update items by ids %+v from %s failed,err:%+v", sensorIDs, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}

//SensorBindGateway ...重新绑定网关
func (t IotSensor) SensorBindGateway(sensorIDs []int, gatewayId, userID int) errs.ErrInfo {
	//先根据GatewayId解除绑定
	unbindSensor := &IotSensor{
		GatewayId: 0,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	cols := []string{"gateway_id", "write_uid", "write_date"}
	_, err := utils.GetMysqlClient().Cols(cols...).Where("gateway_id=?", gatewayId).And("user_id=?", userID).Update(unbindSensor)
	if err != nil {
		glog.Errorf("Update items by gateway_id %d from %s failed,err:%+v", gatewayId, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	//重新绑定sensorIDs的GatewayId
	bindSensor := &IotSensor{
		GatewayId: gatewayId,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	_, err = utils.GetMysqlClient().Cols(cols...).In("id", sensorIDs).And("user_id=?", userID).Update(bindSensor)
	if err != nil {
		glog.Errorf("Update items by ids %+v from %s failed,err:%+v", sensorIDs, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}

//BindItemByUser ... APP绑定用户ID
func (t IotSensor) BindItemByUser(code string, userID int) (bool, errs.ErrInfo) {
	if code == "" || userID <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"user_id", "write_uid", "write_date"}
	updateItem := &IotSensor{
		UserId:    userID,
		WriteUid:  userID,
		WriteDate: time.Now(),
	}
	condition := fmt.Sprintf("code=%s", code)
	rows, err := utils.GetMysqlClient().Cols(cols...).Where(condition).Update(updateItem)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}

//UnbindItemByUser ... APP用户解绑
func (t IotSensor) UnbindItemByUser(id, userID int) (bool, errs.ErrInfo) {
	if id <= 0 || userID <= 0 {
		return false, errs.ErrBadRequest
	}
	cols := []string{"field_id", "gateway_id", "depth", "user_id", "write_uid", "write_date"}
	updateItem := &IotSensor{
		FieldId:   0,
		GatewayId: 0,
		UserId:    0,
		Depth:     0,
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
