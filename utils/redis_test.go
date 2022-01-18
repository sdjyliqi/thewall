package utils

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Redis(t *testing.T) {
	key, value := "name", "liqi"
	client := GetRedisClient()
	set := client.Set(context.Background(), key, value, 0)
	t.Log(set)
	result := client.Get(context.Background(), key)
	assert.Equal(t, value, result.Val())
}
