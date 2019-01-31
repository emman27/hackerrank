package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_catAndMouse(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Cat B", catAndMouse(1, 2, 3))
	assert.Equal("Mouse C", catAndMouse(1, 3, 2))
}
