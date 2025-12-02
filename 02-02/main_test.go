package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInvalidIdsInRange(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{11, 22}, GetInvalidIdsInRange(11, 22))
	assert.Equal([]int{1188511885}, GetInvalidIdsInRange(1188511880, 1188511890))
}

func TestIsIdInvalid(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsIdInvalid(11))
	assert.True(IsIdInvalid(1212))
	assert.True(IsIdInvalid(121212))
	assert.True(IsIdInvalid(112112112))
	assert.False(IsIdInvalid(1212121))
}
