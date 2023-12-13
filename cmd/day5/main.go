package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"advent-of-code/pkg/file"
)

func main() {
	scanner, f := file.NewScanner("tests/input.txt")
	defer f.Close()

	loc, rangedLoc := GetLowsetLoc(scanner)
	fmt.Println(loc)
	fmt.Println(rangedLoc)
}

func GetLowsetLoc(scanner *bufio.Scanner) (int, int) {
	var seeds []int
	var seedsToSoil [][]int
	var soilToFert [][]int
	var fertToWater [][]int
	var waterToLight [][]int
	var lightToTemp [][]int
	var tempToHum [][]int
	var humToLoc [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds") {
			seeds = getSeeds(line)
		}
		if strings.Contains(line, "seed-to-soil") {
			seedsToSoil = getMaps(scanner)
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			soilToFert = getMaps(scanner)
		}
		if strings.Contains(line, "fertilizer-to-water") {
			fertToWater = getMaps(scanner)
		}
		if strings.Contains(line, "water-to-light") {
			waterToLight = getMaps(scanner)
		}
		if strings.Contains(line, "light-to-temperature") {
			lightToTemp = getMaps(scanner)
		}
		if strings.Contains(line, "temperature-to-humidity") {
			tempToHum = getMaps(scanner)
		}
		if strings.Contains(line, "humidity-to-location") {
			humToLoc = getMaps(scanner)
		}
	}

	lowestLoc := math.MaxInt
	for _, seed := range seeds {
		soil := getMappedValue(seed, seedsToSoil)
		fert := getMappedValue(soil, soilToFert)
		water := getMappedValue(fert, fertToWater)
		light := getMappedValue(water, waterToLight)
		temp := getMappedValue(light, lightToTemp)
		hum := getMappedValue(temp, tempToHum)
		loc := getMappedValue(hum, humToLoc)

		if loc < lowestLoc {
			lowestLoc = loc
		}
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < (len(seeds)/2)+1; i += 2 {
		wg.Add(1)
		start := seeds[i]
		length := seeds[i+1]
		go func() {
			l := math.MaxInt
			for j := start; j < start+length; j++ {
				soil := getMappedValue(j, seedsToSoil)
				fert := getMappedValue(soil, soilToFert)
				water := getMappedValue(fert, fertToWater)
				light := getMappedValue(water, waterToLight)
				temp := getMappedValue(light, lightToTemp)
				hum := getMappedValue(temp, tempToHum)
				loc := getMappedValue(hum, humToLoc)

				if loc < l {
					l = loc
				}
			}

			ch <- l
			// Signal that the goroutine is done
			wg.Done()
		}()
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	lowestLocRange := math.MaxInt
	for val := range ch {
		if val < lowestLocRange {
			lowestLocRange = val
		}
	}

	return lowestLoc, lowestLocRange
}

func getSeeds(line string) []int {
	var seeds []int
	seedStrs := strings.Split(line, " ")
	for _, s := range seedStrs {
		seed, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func getMaps(scanner *bufio.Scanner) [][]int {
	var maps [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		paramStrs := strings.Split(line, " ")
		var params []int
		for _, paramStr := range paramStrs {
			param, err := strconv.Atoi(paramStr)
			if err != nil {
				continue
			}
			params = append(params, param)
		}

		maps = append(maps, params)
	}

	return maps
}

func getMappedValue(val int, maps [][]int) int {
	for _, m := range maps {
		if val < m[1] || val >= m[1]+m[2] {
			continue
		}
		return val - m[1] + m[0]
	}

	return val
}
