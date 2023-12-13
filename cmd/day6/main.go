package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"advent-of-code/pkg/file"
)

type Race struct {
	time     int
	distance int
}

func main() {
	scanner, f := file.NewScanner("tests/input.txt")
	defer f.Close()
	result := GetRaceOtherResults(scanner)
	fmt.Println(result)
}

func mustParse(i int, err error) int {
	return i
}

func GetRaceOtherResults(scanner *bufio.Scanner) int {
	result := 1

	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	d := scanner.Text()

	var timeStrs []string
	var distStrs []string
	for _, c := range t {
		if unicode.IsDigit(c) {
			timeStrs = append(timeStrs, string(c))
		}
	}
	for _, c := range d {
		if unicode.IsDigit(c) {
			distStrs = append(distStrs, string(c))
		}
	}

	race := Race{
		time:     mustParse(strconv.Atoi(strings.Join(timeStrs, ""))),
		distance: mustParse(strconv.Atoi(strings.Join(distStrs, ""))),
	}

	minHold, maxHold := 0, 0
	for i := 1; i < race.time; i++ {
		if i*(race.time-i) > race.distance {
			minHold = i
			break
		}
	}
	for i := race.time; i > 1; i-- {
		if i*(race.time-i) > race.distance {
			maxHold = i
			break
		}
	}
	result *= maxHold - minHold + 1

	return result

}

func GetRaceResults(scanner *bufio.Scanner) int {
	result := 1
	var races []Race
	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	d := scanner.Text()

	timeStrs := strings.Split(t, " ")
	distanceStrs := strings.Split(d, " ")

	var times []int
	var distances []int
	for _, timeStr := range timeStrs {
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			continue
		}

		times = append(times, time)
	}
	for _, distSrt := range distanceStrs {
		dist, err := strconv.Atoi(distSrt)
		if err != nil {
			continue
		}

		distances = append(distances, dist)
	}

	if len(times) != len(distances) {
		return -1
	}

	for i, time := range times {
		races = append(races, Race{time: time, distance: distances[i]})
	}

	for _, race := range races {
		minHold, maxHold := 0, 0
		for i := 1; i < race.time; i++ {
			if i*(race.time-i) > race.distance {
				minHold = i
				break
			}
		}
		for i := race.time; i > 1; i-- {
			if i*(race.time-i) > race.distance {
				maxHold = i
				break
			}
		}
		result *= maxHold - minHold + 1
	}

	return result

}
