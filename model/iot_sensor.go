package models

import (
	"time"
)

type IotSensor struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	Name         string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code         string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	DeviceId     int       `json:"device_id" xorm:"comment('Device') index INT(11)"`
	SensorTypeId int       `json:"sensor_type_id" xorm:"comment('Sensor Type') index INT(11)"`
	Depth        int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	CreateUid    int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate   time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid     int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate    time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotSensor) TableName() string {
	return "iot_sensor"
}
