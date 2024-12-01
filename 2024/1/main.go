package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var left []int
	var right []int
	sum := 0

	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		leftInt, err := strconv.Atoi(strings.Split(line, "   ")[0])

		if err != nil {
			log.Fatal(err)
		}

		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(strings.Split(line, "   ")[1])

		if err != nil {
			log.Fatal(err)
		}

		right = append(right, rightInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(left)
	sort.Ints(right)

	// Part One
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		sum += Abs(diff)
	}

	fmt.Println("Part One:", sum)

	// Part Two
	sum = 0

	for i := 0; i < len(left); i++ {
		counter := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				counter++
			}
		}

		sum += left[i] * counter
	}

	fmt.Println("Part Two:", sum)
}
