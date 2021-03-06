package model

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var IotValueModel IotValue

type IotValue struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Code         string    `json:"code" xorm:"VARCHAR(32)"`
	EtlTimestamp int       `json:"etl_timestamp" xorm:"not null comment('Etl Time') INT(11)"`
	FieldId      int       `json:"field_id" xorm:"comment('field_id') INT(11)"`
	Depth        int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	Value        int       `json:"value" xorm:"comment('Value') INT(11)"`
	CreateUid    int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate   time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid     int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate    time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotValue) TableName() string {
	return "iot_value"
}

//探针和农田的关联查询
type ValueEx struct {
	IotValue `xorm:"extends"`
	IotField `xorm:"extends"`
	IotProbe `xorm:"extends"`
}

//GetItemsByPage  ...分页获取全量数据
func (t IotValue) GetItemsByPage(pageID int) ([]*IotValue, error) {
	if pageID < 0 {
		return nil, errors.New("invalid-request")
	}
	var items []*IotValue
	pageCount := 100
	err := utils.GetMysqlClient().Limit(pageCount, pageID*pageCount).Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//AddItem ... 添加一条数据
func (t IotValue) AddItem(item *IotValue) (bool, errs.ErrInfo) {
	item.WriteDate = time.Now()
	rows, err := utils.GetMysqlClient().InsertOne(item)
	if err != nil {
		glog.Errorf("Insert item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}

//GetLineItems  ...获取某个传感器的时间区间内的所有数据
func (t IotValue) GetLineItems(sensorID int, startTS, stopTS int64) ([]*IotValue, errs.ErrInfo) {
	var items []*IotValue
	cols := []string{"id", "etl_timestamp", "value"}
	err := utils.GetMysqlClient().Cols(cols...).Where("sensor_id=?", sensorID).
		And("etl_timestamp >=?", startTS).
		And("etl_timestamp <=?", stopTS).
		OrderBy("etl_timestamp").
		Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

//GetItemsByCodes  ...获取某个些探针的在某个时间段内的数据
func (t IotValue) GetItemsByCodes(probeCodes []string, startTS, stopTS int64) ([]*IotValue, errs.ErrInfo) {
	var items []*IotValue
	cols := []string{"id", "etl_timestamp", "value", "code", "depth"}
	err := utils.GetMysqlClient().
		Cols(cols...).
		In("code", probeCodes).
		And("etl_timestamp >=?", startTS).
		And("etl_timestamp <=?", stopTS).
		OrderBy("etl_timestamp").
		Find(&items)
	if err != nil {
		glog.Errorf("Find the items from %s failed,err:%+v", t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

//UpdateItemByCols ... 按照probe code更新数据，只修改固定列
func (t IotValue) GetProbeWithField(probeID string, startTS, stopTS int64) ([]*ValueEx, errs.ErrInfo) {
	var items []*ValueEx
	joinCon := fmt.Sprintf("%s.field_id=%s.id", t.TableName(), FieldModel.TableName())
	whereCon := fmt.Sprintf("%s.code='%s'", t.TableName(), probeID)
	err := utils.GetMysqlClient().Table(t.TableName()).
		Join("LEFT", FieldModel.TableName(), joinCon).
		Join("LEFT", ProbeModel.TableName(), "iot_value.code=iot_probe.code").
		Where(whereCon).
		And("etl_timestamp >=?", startTS).
		And("etl_timestamp <=?", stopTS).
		Find(&items)
	if err != nil {
		glog.Errorf("Get the item by probe code %s from %s failed,err:%+v", probeID, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}
