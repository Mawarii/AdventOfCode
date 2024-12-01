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
	n := PartOne(bufio.NewScanner(strings.NewReader(input)))

	assert.Equal(t, 11, n)
}

func TestSamplePartTwo(t *testing.T) {
	n := PartTwo(bufio.NewScanner(strings.NewReader(input)))

	assert.Equal(t, 31, n)
}

func BenchmarkMain(b *testing.B) {
	for range b.N {
		main()
	}
}

func BenchmarkPartOne(b *testing.B) {
	for range b.N {
		file, _ := os.Open("input")
		defer file.Close()
		_ = PartOne(bufio.NewScanner(file))
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for range b.N {
		file, _ := os.Open("input")
		defer file.Close()

		_ = PartTwo(bufio.NewScanner(file))
	}
}
