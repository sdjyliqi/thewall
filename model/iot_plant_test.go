package model

import (
	"testing"
)

func Test_IotPlantGetHistoryPlant(t *testing.T) {
	items, errEx := PlantModel.GetHistoryPlant(1)
	t.Log(errEx)
	t.Log(items)
	for _, v := range items {
		t.Log(v)
	}
}
