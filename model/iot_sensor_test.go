package model

import (
	"testing"
)

func Test_IotSensorGetItemsByField(t *testing.T) {
	items, err := SensorModel.GetItemsByField(1)
	t.Log(err)
	t.Log(items)
	for _, v := range items {
		t.Log(v)
	}
}

func Test_IotSensorGetItemsByUser(t *testing.T) {
	items, err := SensorModel.GetItemsByUser(1)
	t.Log(err)
	t.Log(items)
	for _, v := range items {
		t.Log(v)
	}
}

func Test_GetItemByID(t *testing.T) {
	item, err := SensorModel.GetItemsByID(4009, 1)
	t.Log(err)
	t.Log(item)
}
