package model

import (
	"testing"
)

func Test_IotCropGetItemsByUser(t *testing.T) {
	items, _ := FieldModel.GetItemsByUser(1)
	t.Log(items)
	for _, v := range items {
		t.Log("field:", v.IotField, " SoilType:", v.IotSoilType)
	}
}

func Test_IotFieldGetItemByID(t *testing.T) {
	item, _ := FieldModel.GetItemByID(1)
	t.Log(item)
}

func Test_IotCropAddFieldByUser(t *testing.T) {
	node := &IotField{
		Name:          "test",
		NameCn:        "test",
		UserId:        1,
		Longitude:     10,
		Latitude:      10,
		Area:          10,
		SoilTypeId:    1,
		CropTypeNowId: 1,
		StateNowId:    0,
	}
	dbNode, _ := FieldModel.AddFieldByUser(node)
	t.Log(dbNode)
}
