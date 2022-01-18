package model

import "testing"

func Test_SensorTypeGetAllItems(t *testing.T) {
	items, err := SensorTypeModel.GetAllItems()
	t.Log(err)
	for _, v := range items {
		t.Log(v)
	}
}
