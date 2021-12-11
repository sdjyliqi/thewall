package models

import (
	"email-center/errs"
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var UCModel IotUc

type IotUc struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Email     string    `json:"email" xorm:"not null default '按照邮箱登录' VARCHAR(128)"`
	Nickname  string    `json:"nickname" xorm:"VARCHAR(128)"`
	Passport  string    `json:"passport" xorm:"not null default '密码' VARCHAR(128)"`
	Token     string    `json:"token" xorm:"VARCHAR(128)"`
	LastLogin time.Time `json:"last_login" xorm:"DATETIME"`
	Desc      string    `json:"desc" xorm:"default '' comment('描述信息') VARCHAR(1024)"`
}

func (t IotUc) TableName() string {
	return "iot_uc"
}

//Login ...用户登录
func (t IotUc) Login(email, password string) (bool, errs.ErrInfo) {
	var item IotUc
	ok, err := utils.GetMysqlClient().Where("email='%s'", email).Get(&item)
	if err != nil {
		glog.Errorf("Get item by email %s from table %s failed,err:%+v", email, t.TableName(), err)
		return false, errs.ErrDBGet
	}
	if !ok {
		return false, errs.ErrUCNoUser
	}
	return utils.EncodingPassword(password) == item.Passport, errs.Succ
}

//UpdateToken  ...更新token
func (t IotUc) UpdateToken(email, token string) errs.ErrInfo {
	item := IotUc{Email: email, Token: token, LastLogin: time.Now()}
	_, err := utils.GetMysqlClient().Where("email='%s'", email).Update(&item)
	if err != nil {
		glog.Errorf("Update item  %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}

//Register  ...用户注册
func (t IotUc) Register() error {
	return nil
}
