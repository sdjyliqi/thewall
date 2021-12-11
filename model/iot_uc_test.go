package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IotUcLogin(t *testing.T) {
	items, err := UCModel.Login("sdjyliqi@163.com", "abcd1234")
	assert.Nil(t, err)
	t.Log(items)
}
