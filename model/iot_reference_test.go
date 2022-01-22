package model

import (
	"testing"
)

func Test_IotReferenceGetAllItems(t *testing.T) {
	items, err := ReferenceModel.GetAllItems()
	t.Log(err)
	t.Log(len(items))
}
