package model

type IotProjectResUsersRel struct {
	IotProjectId int `json:"iot_project_id" xorm:"not null index INT(11)"`
	ResUsersId   int `json:"res_users_id" xorm:"not null index INT(11)"`
}

func (t IotProjectResUsersRel) TableName() string {
	return "iot_project_res_users_rel"
}
