package model

import (
	"encoding/json"
	"testing"
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
