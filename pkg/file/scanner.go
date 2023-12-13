package file

import (
	"bufio"
	"log"
	"os"
)

func NewScanner(filename string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner, file
}
