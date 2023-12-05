package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	sum := Part1("tests/test.txt")
	fmt.Println(sum)
}

func TestPart2(t *testing.T) {
	sum := Part2("tests/test.txt")
	fmt.Println(sum)
}
