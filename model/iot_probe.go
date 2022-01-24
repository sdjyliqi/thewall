package model

import (
	"fmt"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var ProbeModel IotProbe

type IotProbe struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SensorName   string    `json:"sensor_name" xorm:"comment('关联sensor的id') VARCHAR(128)"`
	Code         string    `json:"code" xorm:"comment('2000年以后的16进制数') VARCHAR(16)"`
	ProbeTypeId  int       `json:"probe_type_id" xorm:"comment('Probe Type') INT(11)"`
	Depth        int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"comment('最后上传数据的时间') DATETIME"`
}

type ProbeItem struct {
	Code  string `json:"code"`
	Depth int    `json:"depth"`
}

type ProbeExtend struct {
	IotSensor `xorm:"extends"`
	IotProbe  `xorm:"extends"`
}

func (t IotProbe) TableName() string {
	return "iot_probe"
}

//GetAllItems  ...获取全量数据
func (t IotProbe) GetAllItems() ([]*IotProbe, error) {
	var items []*IotProbe
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetProbesByFieldID  ...根据土地ID获取所有的传感器上探针的列表
func (t IotProbe) GetProbesByFieldID(field int) ([]*ProbeExtend, errs.ErrInfo) {
	var items []*ProbeExtend
	jsonCondition := fmt.Sprintf("%s.%s=%s.%s", t.TableName(), "sensor_name", SensorModel.TableName(), "name")
	whereCondition := fmt.Sprintf("%s.%s=%d", SensorModel.TableName(), "field_id", field)
	err := utils.GetMysqlClient().Table(t.TableName()).
		Join("LEFT", SensorModel.TableName(), jsonCondition).
		Where(whereCondition).
		Find(&items)
	if err != nil {
		glog.Errorf("The the items by field_id %d from %s failed,err:%+v", field, t.TableName(), err)
		return nil, errs.ErrDBGet
	}
	return items, errs.Succ
}

//UpdateItem ... 更新数据Depth
func (t IotProbe) UpdateItem(item *IotProbe) (bool, errs.ErrInfo) {
	var rows int64 = 0
	var err error = nil
	cols := []string{"depth"}
	updateItem := &IotProbe{
		Depth: item.Depth,
	}
	if item.Id > 0 {
		rows, err = utils.GetMysqlClient().Cols(cols...).ID(item.Id).Update(updateItem)
	}
	if item.Code != "" {
		condition := fmt.Sprintf("code='%s'", item.Code)
		rows, err = utils.GetMysqlClient().Cols(cols...).Where(condition).Update(updateItem)
	}
	//if item.SensorName != "" {
	//	rows, err = utils.GetMysqlClient().Cols(cols...).Where("sensor_id=?", item.SensorId).Update(updateItem)
	//}
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", updateItem, t.TableName(), err)
		return false, errs.ErrDBUpdate
	}
	return rows > 0, errs.Succ
}
