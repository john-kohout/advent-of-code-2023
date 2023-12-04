package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type Line struct {
	ID           int
	numbers      []Number
	gearsIndices []int
}

type Number struct {
	value int
	start int
	end   int
}

const findDigit = "[0-9]"
const findCharacter = "[^0-9.]"

func main() {
	sum := Part1("test/input.txt")
	fmt.Println(sum)
	sum2 := GetSumGearRatio("test/input.txt")
	fmt.Println(sum2)
}

func consume(number []rune) (int, []int) {
	var incl []int
	if len(number) == 0 {
		return 0, incl
	}

	num, err := strconv.Atoi(string(number))
	if err != nil {
		panic(err)
	}
	return num, append(incl, num)
}

func getLine(id int, line string) Line {
	var numbers []Number
	var gearsIndices []int
	var runes []rune
	startIdx := -1

	for index, char := range line {
		if unicode.IsDigit(char) {
			runes = append(runes, char)
			if startIdx == -1 {
				startIdx = index
			}
			continue
		}
		if char == '*' {
			gearsIndices = append(gearsIndices, index)
		}

		if len(runes) != 0 {
			num, err := strconv.Atoi(string(runes))
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, Number{
				value: num,
				start: startIdx,
				end:   index,
			})

			runes = make([]rune, 0)
			startIdx = -1
		}
	}

	if len(runes) != 0 {
		num, err := strconv.Atoi(string(runes))
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, Number{
			value: num,
			start: startIdx,
			end:   len(line),
		})
	}

	return Line{
		ID:           id,
		numbers:      numbers,
		gearsIndices: gearsIndices,
	}
}

func getCurrentLine(prevLine, currentLine, nextLine string) (int, []int) {
	sum := 0
	var include []int

	chars := regexp.MustCompile(findCharacter)

	prevLineChars := chars.FindAllStringIndex(prevLine, -1)
	currentLineChars := chars.FindAllStringIndex(currentLine, -1)
	nextLineChars := chars.FindAllStringIndex(nextLine, -1)

	var number []rune
	startIdx := -1

	for index, char := range currentLine {
		if unicode.IsDigit(char) {
			number = append(number, char)
			if startIdx == -1 {
				startIdx = index
			}
			continue
		} else if char != '.' {
			if len(number) != 0 {
				num, incl := consume(number)
				include = append(include, incl...)
				sum += num
				number = make([]rune, 0)
				startIdx = -1
			}
			continue
		}
		if len(number) != 0 {
			num, incl := validateAndConsume(number, startIdx, index, prevLineChars, currentLineChars, nextLineChars)
			include = append(include, incl...)
			sum += num
			number = make([]rune, 0)
			startIdx = -1
		}
	}

	if len(number) != 0 {
		num, incl := validateAndConsume(number, startIdx, len(currentLine)-1, prevLineChars, currentLineChars, nextLineChars)
		include = append(include, incl...)
		sum += num
	}

	return sum, include
}

func validateAndConsume(number []rune, startIdx int, index int, prevLineChars [][]int, currentLineChars [][]int, nextLineChars [][]int) (int, []int) {
	for _, c := range currentLineChars {
		if c[0] == startIdx-1 {
			return consume(number)
		}
	}

	for _, c := range prevLineChars {
		if c[0] >= startIdx-1 && c[0] <= index {
			return consume(number)
		}
	}

	for _, c := range nextLineChars {
		if c[0] >= startIdx-1 && c[0] <= index {
			return consume(number)
		}
	}

	return 0, []int{}
}

func Part1(filename string) int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)

	var included [][]int

	prevLine, currentLine, nextLine := "", "", ""

	for scanner.Scan() {
		if prevLine == "" {
			prevLine = scanner.Text()
			continue
		}
		if currentLine == "" {
			currentLine = scanner.Text()
			num, incl := getCurrentLine("", prevLine, currentLine)
			sum += num
			included = append(included, incl)
			continue
		}
		nextLine = scanner.Text()
		num, incl := getCurrentLine(prevLine, currentLine, nextLine)
		sum += num
		included = append(included, incl)
		prevLine = currentLine
		currentLine = nextLine
	}

	num, incl := getCurrentLine(prevLine, currentLine, "")
	sum += num
	included = append(included, incl)

	return sum
}

func GetLines(filename string) []Line {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var lines []Line

	scanner := bufio.NewScanner(f)
	idx := 0

	for scanner.Scan() {
		lines = append(lines, getLine(idx, scanner.Text()))
		idx++
	}

	return lines
}

func GetGearRatios(prevLine, currentLine, nextLine Line) int {
	gearRatios := 0
	for _, gear := range currentLine.gearsIndices {
		var values []int
		for _, number := range prevLine.numbers {
			if number.start-1 <= gear && number.end >= gear {
				values = append(values, number.value)
			}
		}

		for _, number := range currentLine.numbers {
			//if number.start-1 == gear && number.end == gear {
			if number.start-1 == gear || number.end == gear {
				values = append(values, number.value)
			}
		}

		for _, number := range nextLine.numbers {
			if number.start-1 <= gear && number.end >= gear {
				values = append(values, number.value)
			}
		}
		if len(values) == 2 {
			gearRatios += values[0] * values[1]
		} else {
			fmt.Printf("%d:%d\n", currentLine.ID+1, gear+1)
		}
	}

	return gearRatios
}

func GetSumGearRatio(filename string) int {
	lines := GetLines(filename)

	sum := 0

	for i := 0; i < len(lines); i++ {
		if i == 0 {
			sum += GetGearRatios(Line{}, lines[0], lines[1])
			continue
		}
		if i == len(lines)-1 {
			sum += GetGearRatios(lines[i-1], lines[i], Line{})
			continue
		}
		sum += GetGearRatios(lines[i-1], lines[i], lines[i+1])
	}

	return sum
}
