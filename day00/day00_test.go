package day00_test

import (
	"aoc2023/day00"
	"testing"
)

var testInput01 = ``

func TestPart01(t *testing.T) {
	expected := ""
	actual, _ := day00.Part01(testInput01)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

var testInput02 = ``

func TestPart02(t *testing.T) {
	expected := ""
	actual, _ := day00.Part02(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
