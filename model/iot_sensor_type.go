package models

import (
	"time"
)

type IotSensorType struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Name       string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code       string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotSensorType) TableName() string {
	return "iot_sensor_type"
}
