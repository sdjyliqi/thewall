package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"thewall/errs"
	"thewall/utils"
	"time"
)

var UCModel IotUc

type IotUc struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Email     string    `json:"email" xorm:"not null pk default '' unique VARCHAR(64)"`
	Country   string    `json:"country" xorm:"default '' comment('用户所在国家') VARCHAR(32)"`
	Nickname  string    `json:"nickname" xorm:"VARCHAR(32)"`
	Token     string    `json:"token" xorm:"VARCHAR(255)"`
	Password  string    `json:"password" xorm:"VARCHAR(32)"`
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

//ChkUserExisted ...查询某用户是否存在
func (t *IotUc) ChkUserExisted(id int) (bool, errs.ErrInfo) {
	var item IotUc
	ok, err := utils.GetMysqlClient().Where("id=?", id).Get(&item)
	if err != nil {
		glog.Errorf("Get item by id %d from table %s failed,err:%+v", id, t.TableName(), err)
		return false, errs.ErrDBGet
	}
	if !ok {
		return false, errs.Succ
	}
	return true, errs.Succ
}

//UpdateToken  ...更新token
func (t *IotUc) UpdateToken(email, token string) errs.ErrInfo {
	item := IotUc{Email: email, Token: token, LastLogin: time.Now()}
	cols := []string{"token", "last_login"}
	_, err := utils.GetMysqlClient().Cols(cols...).Where(fmt.Sprintf("email='%s'", email)).Update(&item)
	if err != nil {
		glog.Errorf("Update item  %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return errs.ErrDBUpdate
	}
	return errs.Succ
}

//Register  ...用户注册
func (t *IotUc) Register(user IotUc) (bool, errs.ErrInfo) {
	var item IotUc
	ok, err := utils.GetMysqlClient().Where(fmt.Sprintf("email='%s'", user.Email)).Or(fmt.Sprintf("nickname='%s'", user.Nickname)).Get(&item)
	if err != nil {
		glog.Errorf("Get item by email %s or nickname %s from table %s failed,err:%+v", user.Email, user.Nickname, t.TableName(), err)
		return false, errs.ErrDBGet
	}
	if ok {
		return false, errs.ErrUCUserExisted
	}
	rows, insertErr := utils.GetMysqlClient().InsertOne(user)
	if insertErr != nil {
		glog.Errorf("Insert user %+v from table %s failed,err:%+v", user, t.TableName(), insertErr)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}

func (t *IotUc) RegisterSession(sess *xorm.Session, user IotUc) (bool, errs.ErrInfo) {
	var item IotUc
	ok, err := sess.Where(fmt.Sprintf("email='%s'", user.Email)).Or(fmt.Sprintf("nickname='%s'", user.Nickname)).Get(&item)
	if err != nil {
		glog.Errorf("Get item by email %s or nickname %s from table %s failed,err:%+v", user.Email, user.Nickname, t.TableName(), err)
		return false, errs.ErrDBGet
	}
	if ok {
		return false, errs.ErrUCUserExisted
	}
	rows, insertErr := sess.InsertOne(user)
	if insertErr != nil {
		glog.Errorf("Insert user %+v from table %s failed,err:%+v", user, t.TableName(), insertErr)
		return false, errs.ErrDBInsert
	}
	return rows > 0, errs.Succ
}

//ResetPassword  ...重置密码
func (t *IotUc) ResetPassword(user IotUc) (bool, errs.ErrInfo) {
	//设置新密码
	user.Desc = "更新密码"
	//设置一下影响的列
	cols := []string{"password", "desc"}
	count, updateErr := utils.GetMysqlClient().Cols(cols...).Where("email=?", user.Email).Update(&user)
	if updateErr != nil {
		glog.Errorf("Update item  %+v from table %s failed,err:%+v", user, t.TableName(), updateErr)
		return false, errs.ErrDBUpdate
	}
	return count > 0, errs.Succ
}
