package model

type IotStateType struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name   string `json:"name" xorm:"VARCHAR(16)"`
	NameCn string `json:"name_cn" xorm:"comment('中文名') VARCHAR(16)"`
}

func (t IotStateType) TableName() string {
	return "iot_state_type"
}
