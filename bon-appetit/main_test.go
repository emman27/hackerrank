package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("5", solve([]int32{3, 10, 2, 9}, 1, 12))
	assert.Equal("Bon Appetit", solve([]int32{3, 10, 2, 9}, 1, 7))
}
