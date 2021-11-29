package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IotCropTypeGetAllItems(t *testing.T) {
	items, err := IotCropEx.GetAllItems()
	assert.Nil(t, err)
	t.Log(items)
}
