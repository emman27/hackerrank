package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMoneySpent(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int32(9), getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
	assert.Equal(int32(-1), getMoneySpent([]int32{5}, []int32{4}, 5))
	assert.Equal(int32(9), getMoneySpent([]int32{5}, []int32{4}, 9))
}
