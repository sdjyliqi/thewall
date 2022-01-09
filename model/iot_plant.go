package model

import (
	"time"
)

type IotFieldPlant struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	FieldId    int       `json:"field_id" xorm:"comment('field_id') INT(11)"`
	CropTypeId int       `json:"crop_type_id" xorm:"comment('crop_type_id') INT(11)"`
	StartDate  time.Time `json:"start_date" xorm:"comment('Start Date') DATE"`
	EndDate    time.Time `json:"end_date" xorm:"comment('End Date') DATE"`
	Amount     float32   `json:"amount" xorm:"comment('Amount') FLOAT(11,2)"`
	StateId    int       `json:"state_id" xorm:"comment('种植周期阶段') INT(11)"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotFieldPlant) TableName() string {
	return "iot_field_plant"
}
