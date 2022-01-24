package model

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_IotProbeGetProbesByFieldID(t *testing.T) {
	items, err := ProbeModel.GetProbesByFieldID(1)
	t.Log(err)
	for _, v := range items {
		b, _ := json.Marshal(v.IotSensor)
		c, _ := json.Marshal(v.IotProbe)
		t.Log("IOT sensor:", string(b))
		t.Log("IOT probe:", string(c))
	}
}

func Test_IotProbeGetProbesByProbeCode(t *testing.T) {
	item, err := ProbeModel.GetProbesByProbeCode("000000000022-01")
	t.Log(err)
	t.Log(item.IotProbe)
	t.Log(item.IotSensor)
}

func Test_IotProbeUpdateItemByCols(t *testing.T) {
	item := IotProbe{
		Code:         "000000000022-01",
		LastValue:    0,
		LastReceived: time.Now(),
	}
	cols := []string{"last_value", "last_received"}
	err := ProbeModel.UpdateItemByCols(&item, cols)
	t.Log(err)

}
