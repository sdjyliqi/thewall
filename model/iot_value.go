package model

import (
	"time"
)

type IotValue struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	EtlTimestamp int       `json:"etl_timestamp" xorm:"not null comment('Etl Time') INT(11)"`
	Code         string    `json:"code" xorm:"not null comment('Code') LONGTEXT"`
	Time         time.Time `json:"time" xorm:"comment('Check Time') DATETIME"`
	ProjectId    int       `json:"project_id" xorm:"comment('Project') index INT(11)"`
	GatewayId    int       `json:"gateway_id" xorm:"comment('Gateway') index INT(11)"`
	DeviceId     int       `json:"device_id" xorm:"comment('Device') index INT(11)"`
	SensorId     int       `json:"sensor_id" xorm:"comment('Sensor') index INT(11)"`
	SensorTypeId int       `json:"sensor_type_id" xorm:"comment('Sensor Type') index INT(11)"`
	Latitude     string    `json:"latitude" xorm:"comment('Latitude') DECIMAL(65,30)"`
	Longitude    string    `json:"longitude" xorm:"comment('Longitude') DECIMAL(65,30)"`
	Depth        int       `json:"depth" xorm:"comment('Depth') INT(11)"`
	Value        int       `json:"value" xorm:"comment('Value') INT(11)"`
	CreateUid    int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate   time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid     int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate    time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotValue) TableName() string {
	return "iot_value"
}
