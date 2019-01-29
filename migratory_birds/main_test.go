package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_migratoryBirds(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int32(3), migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}))
	assert.Equal(int32(4), migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
	assert.Equal(int32(1), migratoryBirds([]int32{1, 1, 2, 2, 3}))
}
