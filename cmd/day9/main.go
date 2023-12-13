package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/pkg/file"
)

func main() {
	scanner, f := file.NewScanner("tests/test.txt")
	defer f.Close()
	result := GetResult(scanner, true)
	fmt.Println(result)
}

type Sequence []int

func IsSequenceZero(s Sequence) bool {
	for _, num := range s {
		if num != 0 {
			return false
		}
	}
	return true
}

func GetResult(scanner *bufio.Scanner, part2 bool) int {
	result := 0

	for scanner.Scan() {
		result += getNextDigit(scanner.Text(), part2)
	}

	return result
}

func getNextDigit(line string, part2 bool) int {
	numbers := strings.Split(line, " ")
	var nums []Sequence
	var num Sequence
	for _, s := range numbers {
		n, _ := strconv.Atoi(s)
		num = append(num, n)
	}

	nums = append(nums, num)
	for !IsSequenceZero(num) {
		var nextNum Sequence
		for i := 0; i < len(num)-1; i++ {
			nextNum = append(nextNum, num[i+1]-num[i])
		}
		num = nextNum
		nums = append(nums, num)
	}

	nextVal := nums[0][0]

	if part2 {
		for i := len(nums) - 1; i > 0; i-- {
			nextVal -= nums[i][0]
		}
		return nextVal
	}

	for _, n := range nums {
		nextVal += n[len(n)-1]
	}

	//if len(nums) == 3 {
	//	return nums[1][len(nums[1])-1] + nums[0][len(nums[0])-1]
	//}
	//
	//nextVal := nums[len(nums)-2][len(nums[len(nums)-2])-1]
	//for i := len(nums) - 3; i >= 0; i-- {
	//	nextVal += nums[i][len(nums[i])-1]
	//}

	return nextVal
}
