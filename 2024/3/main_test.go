package main

import (
	"testing"

	"gotest.tools/assert"
)

var testInput string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestSamplePartOne(t *testing.T) {
	n := PartOne(testInput)

	assert.Equal(t, 161, n)
}

func TestSamplePartTwo(t *testing.T) {
	n := PartTwo(testInput)

	assert.Equal(t, 48, n)
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
