package main

import (
	"bufio"
	"fmt"

	"advent-of-code/pkg/file"
)

func main() {
	scanner, f := file.NewScanner("tests/input.txt")
	defer f.Close()
	result := GetResult(scanner)
	fmt.Println(result)
}


func GetResult(scanner *bufio.Scanner) int {
	result := 0

	return result

}
