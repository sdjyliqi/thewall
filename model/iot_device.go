package models

import (
	"time"
)

type IotDevice struct {
	Id          int       `json:"id" xorm:"not null pk INT(11)"`
	UserId      int       `json:"user_id" xorm:"comment('User') index INT(11)"`
	PartnerId   int       `json:"partner_id" xorm:"comment('Partner') index INT(11)"`
	CompanyId   int       `json:"company_id" xorm:"comment('Company') index INT(11)"`
	GatewayId   int       `json:"gateway_id" xorm:"comment('Gateway') index INT(11)"`
	ProjectId   int       `json:"project_id" xorm:"comment('Project') index INT(11)"`
	Name        string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code        string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	Latitude    string    `json:"latitude" xorm:"comment('Latitude') DECIMAL(65,30)"`
	Longitude   string    `json:"longitude" xorm:"comment('Longitude') DECIMAL(65,30)"`
	Interval    int       `json:"interval" xorm:"comment('Interval') INT(11)"`
	LastEtlTime time.Time `json:"last_etl_time" xorm:"comment('Last Check Time') DATETIME"`
	CreateUid   int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate  time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid    int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate   time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotDevice) TableName() string {
	return "iot_device"
}
