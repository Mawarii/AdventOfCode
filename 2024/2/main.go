package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isStrictlyDescending(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] >= slice[i-1] {
			return false
		}
	}
	return true
}

func isStrictlyAscending(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] <= slice[i-1] {
			return false
		}
	}
	return true
}

//go:embed input
var input string

func main() {
	fmt.Println(PartOne(input))
}

func PartOne(input string) int {
	sum := 0
	numbers := make([]int, 8)
	counter := false

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numbersString := strings.Split(line, " ")
		numbers = numbers[:0]

		for j := 0; j < len(numbersString); j++ {
			a, _ := strconv.Atoi(numbersString[j])
			numbers = append(numbers, a)
		}

		isDescending := isStrictlyDescending(numbers)

		isAscending := isStrictlyAscending(numbers)

		for i := 0; i < len(numbers)-2; i++ {
			if !isDescending && !isAscending {
				break
			} else {
				if (Abs(numbers[i]-numbers[i+1]) > 3) || (Abs(numbers[i+1]-numbers[i+2]) > 3) {
					break
				} else {
					counter = true
				}
			}
			if counter && (i == len(numbers)-3) {
				sum += 1
			}
		}
	}
	return sum
}
