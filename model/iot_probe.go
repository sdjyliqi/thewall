package model

import (
	"time"
)

type IotProbe struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SensorId     string    `json:"sensor_id" xorm:"comment('Name') LONGTEXT"`
	Code         string    `json:"code" xorm:"comment('2000年以后的16进制数') VARCHAR(16)"`
	SensorTypeId int       `json:"sensor_type_id" xorm:"comment('Sensor Type') INT(11)"`
	Depth        int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"comment('最后上传数据的时间') DATETIME"`
}

func (t IotProbe) TableName() string {
	return "iot_probe"
}
