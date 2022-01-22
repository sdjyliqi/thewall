package model

import (
	"github.com/golang/glog"
	"thewall/utils"
	"time"
)

var ReferenceModel IotReference

type IotReference struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SoilTypeId  int       `json:"soil_type_id" xorm:"not null INT(11)"`
	CropTypeId  int       `json:"crop_type_id" xorm:"not null INT(11)"`
	HumidityMin float32   `json:"humidity_min" xorm:"FLOAT"`
	HumidityMax float32   `json:"humidity_max" xorm:"FLOAT"`
	WriteDate   time.Time `json:"write_date" xorm:"comment('修改时间') DATETIME"`
}

func (t IotReference) TableName() string {
	return "iot_reference"
}

//GetAllItems  ...获取全量数据
func (t IotReference) GetAllItems() ([]*IotReference, error) {
	var items []*IotReference
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("The the items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
