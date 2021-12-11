package models

import (
	"email-center/errs"
	"email-center/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"time"
)

var UCModel IotUc

type IotUc struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Email     string    `json:"email" xorm:"default '' unique VARCHAR(64)"`
	Token     string    `json:"token" xorm:"VARCHAR(255)"`
	Password  string    `json:"password" xorm:"VARCHAR(128)"`
	LastLogin time.Time `json:"last_login" xorm:"DATETIME"`
	Desc      string    `json:"desc" xorm:"default '' comment('描述信息') VARCHAR(1024)"`
}

func (t *IotUc) TableName() string {
	return "iot_uc"
}

//Login ...用户登录
func (t *IotUc) Login(email, password string) (bool, errs.ErrInfo) {
	var item IotUc
	ok, err := utils.GetMysqlClient().Where(fmt.Sprintf("email='%s'", email)).Get(&item)
	if err != nil {
		glog.Errorf("Get item by email %s from table %s failed,err:%+v", email, t.TableName(), err)
		return false, errs.ErrDBGet
	}
	if !ok {
		return false, errs.ErrUCNoUser
	}
	return utils.EncodingPassword(password) == item.Password, errs.Succ
}

//UpdateToken  ...更新token
func (t *IotUc) UpdateToken(email, token string) errs.ErrInfo {
	item := IotUc{Email: email, Token: token, LastLogin: time.Now()}
	_, err := utils.GetMysqlClient().Where("email='%s'", email).Update(&item)
	if err != nil {
		glog.Errorf("Update item  %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}

//Register  ...用户注册
func (t *IotUc) Register() error {
	return nil
}
