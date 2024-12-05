package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

func FindMuls(input string) []string {
	expr, _ := regexp.Compile(`mul\(\d+,\d+\)`)

	return expr.FindAllString(input, -1)
}

func Multiply(var1 string, var2 string) int {
	int1, _ := strconv.Atoi(var1)
	int2, _ := strconv.Atoi(var2)

	return int1 * int2
}

//go:embed input
var input string

func main() {
	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input string) int {
	sum := 0

	allFindings := FindMuls(input)

	for _, line := range allFindings {
		exprInt, _ := regexp.Compile(`\d+`)
		foundInts := exprInt.FindAllString(line, -1)
		sum += Multiply(foundInts[0], foundInts[1])
	}

	return sum
}

func PartTwo(input string) int {
	sum := 0

	exprDo, _ := regexp.Compile(`do\(\).+don't\(\)`)

	allDos := exprDo.FindAllString(input, -1)

	for _, dos := range allDos {
		allMuls := FindMuls(dos)

		for _, line := range allMuls {
			exprInt, _ := regexp.Compile(`\d+`)
			foundInts := exprInt.FindAllString(line, -1)
			sum += Multiply(foundInts[0], foundInts[1])
		}
	}

	return sum
}
