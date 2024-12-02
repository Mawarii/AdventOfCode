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

//go:embed input
var input string

func main() {
	sum := 0
	numbsInt := make([]int, 8)
	counter := false

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numbs := strings.Split(line, " ")
		numbsInt = numbsInt[:0]

		for j := 0; j < len(numbs); j++ {
			a, _ := strconv.Atoi(numbs[j])
			numbsInt = append(numbsInt, a)
		}

		for i := 0; i < len(numbsInt)-2; i++ {
			if !(numbsInt[i] > numbsInt[i+1] && numbsInt[i+1] > numbsInt[i+2]) && !(numbsInt[i] < numbsInt[i+1] && numbsInt[i+1] < numbsInt[i+2]) {
				break
			} else {
				if (Abs(numbsInt[i]-numbsInt[i+1]) > 3) || (Abs(numbsInt[i+1]-numbsInt[i+2]) > 3) {
					break
				} else {
					counter = true
				}
			}
			if counter && (i == len(numbsInt)-3) {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}
