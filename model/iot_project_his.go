package model

import (
	"time"
)

type IotProjectHis struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	ProjectId  int       `json:"project_id" xorm:"comment('project_id') index INT(11)"`
	CropTypeId int       `json:"crop_type_id" xorm:"comment('crop_type_id') index INT(11)"`
	StartDate  time.Time `json:"start_date" xorm:"comment('Start Date') DATE"`
	EndDate    time.Time `json:"end_date" xorm:"comment('End Date') DATE"`
	Amount     string    `json:"amount" xorm:"comment('Amount') DECIMAL(65,30)"`
	CreateUid  int       `json:"create_uid" xorm:"comment('Created by') index INT(11)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('Created on') DATETIME"`
	WriteUid   int       `json:"write_uid" xorm:"comment('Last Updated by') index INT(11)"`
	WriteDate  time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotProjectHis) TableName() string {
	return "iot_project_his"
}
