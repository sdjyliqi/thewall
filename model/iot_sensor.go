package model

import (
	"errors"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var SensorModel IotSensor

type IotSensor struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name            string    `json:"name" xorm:"comment('Name') LONGTEXT"`
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
func (t IotSensor) GetItemByID(id int64) (*IotSensor, error) {
	var item *IotSensor
	_, err := utils.GetMysqlClient().ID(id).Get(item)
	if err != nil {
		glog.Errorf("The the item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return nil, err
	}
	return item, nil
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

//UpdateItemByID ... 根据数据ID，更新该条数据数据
func (t IotSensor) UpdateItemByID(item *IotSensor) (int64, error) {
	rows, err := utils.GetMysqlClient().Id(item.Id).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", item, t.TableName(), err)
		return 0, err
	}
	return rows, nil
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
		FieldId:    0,
		CreateDate: time.Now(),
		WriteDate:  time.Now(),
	}
	cols := []string{"field_id", "create_date", "write_date"}
	_, err := utils.GetMysqlClient().Cols(cols...).Where("field_id=?", field).Update(unbindSensor)
	if err != nil {
		glog.Errorf("Update items by field_id %d from %s failed,err:%+v", field, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	//重新针对sensor的id 实施绑定
	bindSensor := &IotSensor{
		FieldId:    field,
		UserId:     userID,
		CreateDate: time.Now(),
		WriteDate:  time.Now(),
	}
	_, err = utils.GetMysqlClient().Cols(cols...).In("id", sensorIDs).Update(bindSensor)
	if err != nil {
		glog.Errorf("Update items by field_ids %+v from %s failed,err:%+v", sensorIDs, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}

//SensorBindGateway ...重新绑定网关
func (t IotSensor) SensorBindGateway(sensorIDs []int, gatewayId, userID int) errs.ErrInfo {
	//先根据sensor绑定的userID和要绑定网关的sensorIDs查出ids
	var ids []int
	err := utils.GetMysqlClient().Select("id").In("id", sensorIDs).And("user_id=?", userID).Find(&ids)
	if err != nil {
		glog.Errorf("Get items by user %d from %s failed,err:%+v", userID, t.TableName(), err)
		return errs.ErrDBGet
	}
	//重新针对查询结果sensor的ids更新绑定网关
	cols := []string{"gateway_id", "create_date", "write_date"}
	bindSensor := &IotSensor{
		FieldId:    gatewayId,
		CreateDate: time.Now(),
		WriteDate:  time.Now(),
	}
	_, err = utils.GetMysqlClient().Cols(cols...).In("id", ids).Update(bindSensor)
	if err != nil {
		glog.Errorf("Update items by field_ids %+v from %s failed,err:%+v", sensorIDs, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}
