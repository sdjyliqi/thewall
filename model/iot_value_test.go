package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_IotValueGetLineItems(t *testing.T) {
	items, err := IotValueModel.GetLineItems(23, 0, time.Now().Unix())
	assert.Nil(t, err)
	for _, v := range items {
		t.Log(v)
	}
}
