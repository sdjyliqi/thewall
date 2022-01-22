package model

import (
	"time"
)

type IotReference struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SoilTypeId  int       `json:"soil_type_id" xorm:"not null INT(11)"`
	HumidityMin float32   `json:"humidity_min" xorm:"FLOAT"`
	CropTypeId  int       `json:"crop_type_id" xorm:"not null INT(11)"`
	HumidityMax float32   `json:"humidity_max" xorm:"FLOAT"`
	WriteDate   time.Time `json:"write_date" xorm:"comment('修改时间') DATETIME"`
}

func (t IotReference) TableName() string {
	return "iot_reference"
}
