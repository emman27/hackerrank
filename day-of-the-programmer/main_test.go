package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dayOfProgrammer(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("12.09.1984", dayOfProgrammer(1984))
	assert.Equal("13.09.2017", dayOfProgrammer(2017))
	assert.Equal("12.09.2016", dayOfProgrammer(2016))
	assert.Equal("12.09.1800", dayOfProgrammer(1800))
	assert.Equal("26.09.1918", dayOfProgrammer(1918))
}
