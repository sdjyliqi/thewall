package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Convert2Int(t *testing.T) {
	result := Convert2Int("100")
	assert.Equal(t, result, 100)

	result64 := Convert2Int64("1642427706")
	t.Log(result64)

}
