package model

import (
	"time"
)

type IotSoilType struct {
	Id                  int       `json:"id" xorm:"not null pk INT(11)"`
	Name                string    `json:"name" xorm:"comment('Name') LONGTEXT"`
	Code                string    `json:"code" xorm:"comment('Code') LONGTEXT"`
	FieldCapacity       string    `json:"field_capacity" xorm:"comment('Field Capacity') DECIMAL(65,30)"`
	TotalAvailableWater string    `json:"total_available_water" xorm:"comment('Total Available Water') DECIMAL(65,30)"`
	Pwp                 string    `json:"pwp" xorm:"comment('PWP') DECIMAL(65,30)"`
	CreateUid           int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate          time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid            int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate           time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotSoilType) TableName() string {
	return "iot_soil_type"
}
