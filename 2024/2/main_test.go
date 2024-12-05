package main

import (
	"testing"

	"gotest.tools/assert"
)

var testInput string = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestSamplePartOne(t *testing.T) {
	n := PartOne(testInput)

	assert.Equal(t, 2, n)
}

func TestSamplePartTwo(t *testing.T) {
	n := PartTwo(testInput)

	assert.Equal(t, 4, n)
}

func BenchmarkMain(b *testing.B) {
	for range b.N {
		main()
	}
}

func BenchmarkPartOne(b *testing.B) {
	for range b.N {
		_ = PartOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for range b.N {
		_ = PartTwo(input)
	}
}
