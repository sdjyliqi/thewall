package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IotCropGetItemsByUser(t *testing.T) {
	items, err := IotFieldEx.GetItemsByUser(1)
	assert.Nil(t, err)
	t.Log(items)
	for _, v := range items {
		t.Log("field:", v.IotField, " CropType:", v.IotCropType)
	}
}
