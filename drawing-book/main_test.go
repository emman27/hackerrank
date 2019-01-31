package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pageCount(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int32(0), pageCount(6, 1))
	assert.Equal(int32(1), pageCount(6, 2))
	assert.Equal(int32(1), pageCount(6, 3))
	assert.Equal(int32(1), pageCount(6, 4))
	assert.Equal(int32(1), pageCount(6, 5))
	assert.Equal(int32(0), pageCount(6, 6))

	assert.Equal(int32(0), pageCount(5, 5))
	assert.Equal(int32(0), pageCount(5, 4))
	assert.Equal(int32(1), pageCount(5, 3))
	assert.Equal(int32(1), pageCount(5, 2))
	assert.Equal(int32(0), pageCount(5, 1))
}
