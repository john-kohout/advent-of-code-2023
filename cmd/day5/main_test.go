package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code/pkg/file"
)

func TestGetSeeds(t *testing.T) {
	want := []int{79, 14, 55, 13}
	scanner, f := file.NewScanner("tests/seeds.txt")
	defer f.Close()

	scanner.Scan()
	line := scanner.Text()
	got := getSeeds(line)
	assert.Equal(t, want, got)
}

func TestGetMap(t *testing.T) {
	want := make(map[int]int)
	for i := 0; i < 8; i++ {
		want[i+53] = i + 49
	}
	for i := 0; i < 42; i++ {
		want[i+11] = i + 0
	}
	for i := 0; i < 7; i++ {
		want[i+0] = i + 42
	}
	for i := 0; i < 4; i++ {
		want[i+7] = i + 57
	}

	scanner, f := file.NewScanner("tests/map.txt")
	defer f.Close()

	scanner.Scan()
	got := getMaps(scanner)
	assert.Equal(t, want, got)
}

func TestGetLowestLoc(t *testing.T) {
	scanner, f := file.NewScanner("tests/test.txt")
	defer f.Close()

	want1 := 35
	want2 := 46
	got1, got2 := GetLowsetLoc(scanner)
	assert.Equal(t, want1, got1)
	assert.Equal(t, want2, got2)
}
