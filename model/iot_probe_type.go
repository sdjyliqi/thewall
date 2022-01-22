package model

import (
	"time"
)

type IotProbeType struct {
	Id        int       `json:"id" xorm:"not null pk INT(11)"`
	Name      string    `json:"name" xorm:"comment('Name') VARCHAR(32)"`
	Code      string    `json:"code" xorm:"comment('Code') VARCHAR(32)"`
	NameCn    string    `json:"name_cn" xorm:"comment('中文名') VARCHAR(32)"`
	WriteDate time.Time `json:"write_date" xorm:"comment('Last Updated on') DATETIME"`
}

func (t IotProbeType) TableName() string {
	return "iot_probe_type"
}
