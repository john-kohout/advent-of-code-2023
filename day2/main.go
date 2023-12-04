package main

import (
	"bufio"
	"fmt"
	"os"

	"advent-of-code/day2/pkg/cubes"
)

func main() {
	sum := part1("input.txt", cubes.Pull{
		Red:   12,
		Green: 13,
		Blue:  14,
	})
	fmt.Printf("Part 1: %d", sum)

	sum = part2("input.txt")
	fmt.Printf("Part 2: %d", sum)
}

func part1(filename string, p cubes.Pull) int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		g, err := cubes.NewGame(scanner.Text())
		if err != nil {
			panic(err)
		}

		countGame := true
		for _, pull := range g.Pulls {
			if pull.Blue > p.Blue || pull.Red > p.Red || pull.Green > p.Green {
				countGame = false
				break
			}
		}

		if !countGame {
			continue
		}
		sum += g.ID
	}

	return sum
}

func part2(filename string) int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		g, err := cubes.NewGame(scanner.Text())
		if err != nil {
			panic(err)
		}

		sum += g.Power()
	}

	return sum
}
