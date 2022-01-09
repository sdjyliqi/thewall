package model

import (
	"testing"
)

func Test_IotCropGetItemsByUser(t *testing.T) {
	items, _ := IotFieldEx.GetItemsByUser(1)
	t.Log(items)
	for _, v := range items {
		t.Log("field:", v.IotField, " SoilType:", v.IotSoilType)
	}
}
