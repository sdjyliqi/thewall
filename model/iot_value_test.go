package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"thewall/errs"
	"time"
)

func Test_IotValueGetLineItems(t *testing.T) {
	items, err := IotValueModel.GetLineItems(23, 0, time.Now().Unix())
	assert.Nil(t, err)
	for _, v := range items {
		t.Log(v)
	}
}

func Test_IotValueGetLinesByCodes(t *testing.T) {
	ids := []string{"000000000022-01"}
	items, err := IotValueModel.GetItemsByCodes(ids, 0, time.Now().Unix())
	assert.Equal(t, err, errs.Succ)
	for _, v := range items {
		t.Log(v)
	}
}
