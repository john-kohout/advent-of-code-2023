package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const regx1 = "[0-9]"

type Card struct {
	ID             int
	WinningNumbers []int
	GameNumbers    []int
	Matches        int
}

func main() {
	sum := Part1("tests/input.txt")
	fmt.Println(sum)
	sum2 := Part2("tests/input.txt")
	fmt.Println(sum2)
}

func GetCard(c string) Card {
	split := strings.Split(c, ":")
	r, _ := regexp.Compile(regx1)
	idSet := r.FindAllString(split[0], -1)
	id, _ := strconv.Atoi(strings.Join(idSet, ""))

	numbers := strings.Split(split[1], "|")
	winning := GetNumbers(numbers[0])
	game := GetNumbers(numbers[1])

	matches := 0
	for _, winner := range winning {
		for _, num := range game {
			if num == winner {
				matches++
			}
		}
	}

	return Card{
		ID:             id,
		WinningNumbers: winning,
		GameNumbers:    game,
		Matches:        matches,
	}
}

func Part2(filename string) int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var cards []Card
	totalWon := 0
	copies := make(map[int]int)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		card := GetCard(scanner.Text())
		cards = append(cards, card)
	}

	for _, card := range cards {
		numCopies := copies[card.ID]
		totalWon++
		for i := 0; i < 1+numCopies; i++ {
			for j := 1; j < card.Matches+1; j++ {
				copies[j+card.ID]++
				totalWon++
			}
		}

	}

	return totalWon
}

func GetNumbers(line string) []int {
	var number []string
	var numbers []int
	for _, char := range line {
		if unicode.IsSpace(char) {
			if len(number) > 0 {
				num, _ := strconv.Atoi(strings.Join(number, ""))
				numbers = append(numbers, num)
				number = make([]string, 0)
			}
			continue
		}
		number = append(number, string(char))
	}
	if len(number) > 0 {
		num, _ := strconv.Atoi(strings.Join(number, ""))
		numbers = append(numbers, num)
	}

	return numbers
}

func GetMatches(line string) int {
	split := strings.Split(line, ":")
	numbers := strings.Split(split[1], "|")
	winning := GetNumbers(numbers[0])
	game := GetNumbers(numbers[1])

	matches := 0
	for _, winner := range winning {
		for _, num := range game {
			if num == winner {
				matches++
			}
		}
	}

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}

func Part1(filename string) int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		sum += GetMatches(scanner.Text())
	}

	return sum
}
