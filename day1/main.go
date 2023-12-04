package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const regx1 = "[0-9]"
const regx2 = "(twone)|(sevenine)|(oneight)|(threeight)|(fiveight)|(nineight)|(eighthree)|(eightwo)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]"

func main() {
	total1 := getTotal("test/input.txt", regx1)
	fmt.Println(total1)
	total2 := getTotal("test/input.txt", regx2)
	fmt.Println(total2)
}

func getTotal(filename string, regx string) int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	//idx := 1

	for scanner.Scan() {
		r, err := regexp.Compile(regx)
		if err != nil {
			panic(err)
		}
		t := scanner.Text()
		findings := r.FindAllString(scanner.Text(), len(t))

		var c []string
		for _, s := range findings {
			c = append(c, ttois(s)...)
		}

		cal := []string{c[0], c[len(c)-1]}
		v, err := strconv.Atoi(strings.Join(cal, ""))
		if err != nil {
			panic(err)
		}

		//fmt.Printf("%d: %d\n", idx, v)
		//idx++

		sum += v
	}
	return sum
}

func ttois(s string) []string {
	switch s {
	case "one":
		return []string{"1"}
	case "two":
		return []string{"2"}
	case "three":
		return []string{"3"}
	case "four":
		return []string{"4"}
	case "five":
		return []string{"5"}
	case "six":
		return []string{"6"}
	case "seven":
		return []string{"7"}
	case "eight":
		return []string{"8"}
	case "nine":
		return []string{"9"}
	case "twone":
		return []string{"2", "1"}
	case "sevenine":
		return []string{"7", "9"}
	case "oneight":
		return []string{"1", "8"}
	case "threeight":
		return []string{"3", "8"}
	case "fiveight":
		return []string{"5", "8"}
	case "nineight":
		return []string{"9", "8"}
	case "eighthree":
		return []string{"8", "3"}
	case "eightwo":
		return []string{"8", "2"}
	}

	return []string{s}
}
