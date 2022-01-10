package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IotGatGetItemsByUser(t *testing.T) {
	items, err := GatewayModel.GetItemsByUser(1)
	assert.Nil(t, err)
	t.Log(items)
	for _, v := range items {
		t.Log(v)
	}
}
