package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minBribes_chaotic(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Too chaotic", minBribes([]int32{2, 5, 1, 3, 4}))
}

func Test_minBribes_singleSwap(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1", minBribes([]int32{1, 2, 3, 5, 4}))
}

func Test_minBribes_multiSwap(t *testing.T) {
	assert.Equal(t, "3", minBribes([]int32{2, 1, 5, 3, 4}))
}
