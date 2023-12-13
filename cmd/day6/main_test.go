package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code/pkg/file"
)

func TestGetRaceResult(t *testing.T) {
	scanner, f := file.NewScanner("tests/test.txt")
	defer f.Close()
	want := 0
	got := GetRaceResults(scanner)
	assert.Equal(t, want, got)
}
