package day09_test

import (
	"aoc2023/day09"
	"testing"
)

var testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestPart01(t *testing.T) {
	expected := "114"
	actual, _ := day09.Part01(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "2"
	actual, _ := day09.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
