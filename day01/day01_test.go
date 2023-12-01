package day01_test

import (
	"aoc2023/day01"
	"testing"
)

var testInput01 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func TestPart01(t *testing.T) {
	expected := "142"
	actual, _ := day01.Part01(testInput01)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

var testInput02 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestPart02(t *testing.T) {
	expected := "281"
	actual, _ := day01.Part02(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	testInput02 += "twotwozvvkrml3nine4fouroneightxg\n"
	expected = "309"
	actual, _ = day01.Part02(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
