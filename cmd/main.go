package main

import (
	"os"

	"advent-of-code/pkg/file"
)

func main() {
	scanner, f := file.NewScanner("cmd/transcript.txt")
	fOut, _ := os.Create("output.txt")

	defer f.Close()
	defer fOut.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 6 {
			continue
		}
		fOut.WriteString(line + "\n")
	}

}
