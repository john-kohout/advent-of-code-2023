package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	sum := getTotal("test/input.txt", regx1)
	assert.Equal(t, 55834, sum)
}

func Test_partTwo(t *testing.T) {
	sum := getTotal("test/input.txt", regx2)
	assert.Equal(t, 53221, sum)
}

func Test_input(t *testing.T) {
	sum := getTotal("test/text.txt", regx2)
	assert.Equal(t, 1098, sum)
}
