package main

import (
	"bufio"
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
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	PartOne(bufio.NewScanner(file))

	file, err = os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	PartTwo(bufio.NewScanner(file))
}

func SortLeftRight(scanner *bufio.Scanner) ([]int, []int, error) {
	left := make([]int, 1000)
	right := make([]int, 1000)

	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, "   ")

		leftInt, err := strconv.Atoi(splitted[0])

		if err != nil {
			log.Print(err)
			continue
		}

		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(splitted[1])

		if err != nil {
			log.Print(err)
			continue
		}

		right = append(right, rightInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right, nil
}

func PartOne(scanner *bufio.Scanner) int {
	left, right, err := SortLeftRight(scanner)

	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		sum += Abs(diff)
	}

	// fmt.Println("Solution Part One:", sum)

	return sum
}

func PartTwo(scanner *bufio.Scanner) int {
	left, right, err := SortLeftRight(scanner)

	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for i := 0; i < len(left); i++ {
		counter := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				counter++
			}
		}

		sum += left[i] * counter
	}

	// fmt.Println("Solution Part Two:", sum)

	return sum
}
