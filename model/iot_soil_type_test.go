package model

import "testing"

func Test_SoilTypeGetAllItems(t *testing.T) {
	items, err := IotSoilTypeModel.GetAllItems()
	t.Log(err)
	for _, v := range items {
		t.Log(v)
	}
}
