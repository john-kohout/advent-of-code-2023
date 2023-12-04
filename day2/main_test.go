package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code/day2/pkg/cubes"
)

func TestPart1(t *testing.T) {
	sum := part1("test.txt", cubes.Pull{
		Red:   12,
		Green: 13,
		Blue:  14,
	})
	fmt.Println(sum)
	assert.Equal(t, 8, sum)
}

func TestPart2(t *testing.T) {
	sum := part2("test.txt")
	fmt.Println(sum)
	assert.Equal(t, 2286, sum)
}
