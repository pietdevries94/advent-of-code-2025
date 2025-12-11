package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineToSystem(t *testing.T) {
	assert := assert.New(t)
	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	expected := &System{
		indicatorTarget: []bool{false, true, true, false},
		buttons:         [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
		joltage:         []int{3, 5, 4, 7},
	}
	actual := LineToSystem(input)
	assert.Equal(expected, actual)
}

func TestParseLine(t *testing.T) {
	assert := assert.New(t)

	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	expected := 10
	actual := ParseLine(input)
	assert.Equal(expected, actual)

	input = "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}"
	expected = 12
	actual = ParseLine(input)
	assert.Equal(expected, actual)

	input = "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}"
	expected = 11
	actual = ParseLine(input)
	assert.Equal(expected, actual)
}
