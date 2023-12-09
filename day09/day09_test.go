package day09_test

import (
	"aoc2023/day09"
	"testing"
)

var testInput01 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestPart01(t *testing.T) {
	expected := "114"
	actual, _ := day09.Part01(testInput01)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

var testInput02 = ``

func TestPart02(t *testing.T) {
	expected := ""
	actual, _ := day09.Part02(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestExtrapolate(t *testing.T) {
	input := []int{
		-3, -6, -10, -3, 44, 181, 485, 1069, 2097, 3823, 6706, 11715, 21054, 39774, 79253, 164654, 350909, 756900, 1637990, 3538025, 7603072,
	}
	expected := 0
	actual := day09.Extrapolate(input)

	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}
}
