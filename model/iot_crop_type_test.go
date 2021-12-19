package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IotCropTypeGetAllItems(t *testing.T) {
	items, err := IotCropEx.GetAllItems()
	assert.Nil(t, err)
	t.Log(items)
}

func Test_IotCropTypeGetItemsByPage(t *testing.T) {
	items, err := IotCropEx.GetItemsByPage(0)
	assert.Nil(t, err)
	t.Log(items)
}
