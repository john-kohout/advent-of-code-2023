package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//  467..114..
//  ...*......
//  ..35..633.
//  ......#...
//  617*......
//  .....+.58.
//  ..592.....
//  ......755.
//  ...$.*....
//  .664.598..

func TestGetNumbers_Small(t *testing.T) {
	got := GetLines("test/smallTest.txt")
	want := []Line{
		{
			ID: 0,
			numbers: []Number{
				{
					value: 467,
					start: 0,
					end:   3,
				},
				{
					value: 114,
					start: 5,
					end:   8,
				},
			},
		},
		{
			ID:           1,
			gearsIndices: []int{3},
		},
	}
	assert.Equal(t, want, got)
}

func TestGetNumbers(t *testing.T) {
	var gears []int
	var numbers []Number

	got := GetLines("test/test.txt")
	want := []Line{
		{
			ID:           0,
			gearsIndices: gears,
			numbers: []Number{
				{
					value: 467,
					start: 0,
					end:   3,
				},
				{
					value: 114,
					start: 5,
					end:   8,
				},
			},
		},
		{
			ID:           1,
			gearsIndices: []int{3},
			numbers:      numbers,
		},
		{
			ID:           2,
			gearsIndices: gears,
			numbers: []Number{
				{
					value: 35,
					start: 2,
					end:   4,
				},
				{
					value: 633,
					start: 6,
					end:   9,
				},
			},
		},
		{
			ID:           3,
			gearsIndices: gears,
			numbers:      numbers,
		},
		{
			ID:           4,
			gearsIndices: []int{3},
			numbers: []Number{
				{
					value: 617,
					start: 0,
					end:   3,
				},
			},
		},
		{
			ID:           5,
			gearsIndices: gears,
			numbers: []Number{
				{
					value: 58,
					start: 7,
					end:   9,
				},
			},
		},
		{
			ID:           6,
			gearsIndices: gears,
			numbers: []Number{
				{
					value: 592,
					start: 2,
					end:   5,
				},
			},
		},
		{
			ID:           7,
			gearsIndices: gears,
			numbers: []Number{
				{
					value: 755,
					start: 6,
					end:   9,
				},
			},
		},
		{
			ID:           8,
			gearsIndices: []int{5},
			numbers:      numbers,
		},
		{
			ID:           9,
			gearsIndices: gears,
			numbers: []Number{
				{
					value: 664,
					start: 1,
					end:   4,
				},
				{
					value: 598,
					start: 5,
					end:   8,
				},
			},
		},
	}
	assert.Equal(t, want, got)
}

func TestGetSumGearRatio(t *testing.T) {
	got := GetSumGearRatio("test/test.txt")
	want := 467835
	assert.Equal(t, want, got)
}

func TestGetSumGearRatio_Big(t *testing.T) {
	got := GetSumGearRatio("test/bigTest.txt")
	want := 1881250
	assert.Equal(t, want, got)
}

func TestPart1_Small(t *testing.T) {
	sum := Part1("test/smallTest.txt")
	assert.Equal(t, 467, sum)
}

func TestPart1(t *testing.T) {
	sum := Part1("test/test.txt")
	assert.Equal(t, 4361, sum)
}

func TestPart1_Big(t *testing.T) {
	sum := Part1("test/bigTest.txt")
	assert.Equal(t, 18816, sum)
}
