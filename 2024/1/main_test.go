package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

var input string = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestSamplePartOne(t *testing.T) {
	left, right := SortLeftRight(bufio.NewScanner(strings.NewReader(input)))

	n := PartOne(left, right)

	assert.Equal(t, 11, n)
}

func TestSamplePartTwo(t *testing.T) {
	left, right := SortLeftRight(bufio.NewScanner(strings.NewReader(input)))

	n := PartTwo(left, right)

	assert.Equal(t, 31, n)
}

func BenchmarkMain(b *testing.B) {
	main()
}

func BenchmarkPartOne(b *testing.B) {
	file, _ := os.Open("input")
	defer file.Close()

	left, right := SortLeftRight(bufio.NewScanner(file))

	_ = PartOne(left, right)
}

func BenchmarkPartTwo(b *testing.B) {
	file, _ := os.Open("input")
	defer file.Close()

	left, right := SortLeftRight(bufio.NewScanner(file))

	_ = PartTwo(left, right)
}
