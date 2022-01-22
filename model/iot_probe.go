package model

import (
	"github.com/golang/glog"
	"thewall/utils"
	"time"
)

var ProbeModel IotProbe

type IotProbe struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SensorId     string    `json:"sensor_id" xorm:"comment('Name') LONGTEXT"`
	Code         string    `json:"code" xorm:"comment('2000年以后的16进制数') VARCHAR(16)"`
	ProbeTypeId  int       `json:"probe_type_id" xorm:"comment('Probe Type') INT(11)"`
	Depth        int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"comment('最后上传数据的时间') DATETIME"`
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
