package handle

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GetTypeName(t *testing.T) {
	go LoadTranslateDic()
	time.Sleep(5 * time.Second)
	v := GetCropTypeByID(1)
	t.Log(v)
	assert.Equal(t, v, "Cotton")

	v = GetSoilTypeByID(1)
	t.Log(v)
	assert.Equal(t, v, "Sand")

	v = GetProbeTypeByID(1)
	assert.Equal(t, v, "Temperature")

	v = GetReferenceNotice(1, 1)
	t.Log(v)
}
