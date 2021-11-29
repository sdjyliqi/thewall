package models

import (
	"time"
)

type IotProject struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Name       string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code       string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	UserId     int       `json:"user_id" xorm:"comment('User') index INT(11)"`
	PartnerId  int       `json:"partner_id" xorm:"comment('Partner') index INT(11)"`
	CompanyId  int       `json:"company_id" xorm:"comment('Company') index INT(11)"`
	AddressId  int       `json:"address_id" xorm:"comment('Address') index INT(11)"`
	SoilTypeId int       `json:"soil_type_id" xorm:"comment('soil_type_id') index INT(11)"`
	CropTypeId int       `json:"crop_type_id" xorm:"comment('crop_type_id') index INT(11)"`
	StartDate  time.Time `json:"start_date" xorm:"comment('Start Date') DATE"`
	EndDate    time.Time `json:"end_date" xorm:"comment('End Date') DATE"`
	Amount     string    `json:"amount" xorm:"comment('Amount') DECIMAL(65,30)"`
	State      string    `json:"state" xorm:"not null comment('Status') LONGTEXT"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotProject) TableName() string {
	return "iot_project"
}
