package model

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"thewall/utils"
)

func Test_IotUcLogin(t *testing.T) {
	items, err := UCModel.Login("sdjyliqi@163.com", "abcd1234")
	assert.Nil(t, err)
	t.Log(items)
}

func Test_IotUcRegister(t *testing.T) {
	email := "yanghao@163.com"
	nickname := "测试用户"
	pwd := utils.EncodingPassword("123456")
	user := IotUc{Email: email, Password: pwd, Nickname: nickname}
	ok, err := UCModel.Register(user)
	assert.Nil(t, err)
	t.Log(ok)
}

//测试
func Test_Transaction(t *testing.T) {
	err := utils.Transaction(utils.GetMysqlClient().NewSession(), func(session *xorm.Session) error {
		email := "yanghao@163.com"
		nickname := "测试用户"
		pwd := utils.EncodingPassword("123456")
		user := IotUc{Email: email, Password: pwd, Nickname: nickname}
		ok, err := UCModel.RegisterSession(session, user)
		fmt.Println(ok, err)
		return errors.New("hello")
	})
	fmt.Println(err)
}
