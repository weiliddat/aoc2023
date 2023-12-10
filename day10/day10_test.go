package day10_test

import (
	"aoc2023/day10"
	"testing"
)

var testInput01 = `.....
.S-7.
.|.|.
.L-J.
.....
`

var testInput02 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

func TestPart01(t *testing.T) {
	expected := "4"
	actual, _ := day10.Part01(testInput01)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	expected = "8"
	actual, _ = day10.Part01(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := ""
	actual, _ := day10.Part02(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func BenchmarkParseToSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Part01(day10.Input)
	}
}

func BenchmarkNewTileMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.NewTileMap(&day10.Input)
	}
}
