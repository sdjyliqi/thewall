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

func Test_IotSensorGetItemByID(t *testing.T) {
	item, err := SensorModel.GetItemsByID(4009, 1)
	t.Log(err)
	t.Log(item)
}

func Test_IotSensorGetItemByName(t *testing.T) {
	item, err := SensorModel.GetItemByName("000000000022")
	t.Log(err)
	t.Log(item)
}
