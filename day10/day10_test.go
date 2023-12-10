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

func BenchmarkNewTileMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := day10.NewTileMap(&day10.Input)
		t.Get(1, 1)
		t.Get(2, 2)
	}
}

func TestTileMap(t *testing.T) {
	tilemap := day10.NewTileMap(&testInput01)

	expected := "."
	actual, found := tilemap.Get(0, 0)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != *actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = "."
	actual, found = tilemap.Get(4, 4)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != *actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = "S"
	actual, found = tilemap.Get(1, 1)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != *actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = "J"
	actual, found = tilemap.Get(3, 3)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != *actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = ""
	actual, found = tilemap.Get(5, 1)
	if found {
		t.Errorf("Expected %#v got %#v", false, found)
	}

	expected = ""
	actual, found = tilemap.Get(1, 5)
	if found {
		t.Errorf("Expected %#v got %#v", false, found)
	}

	expected = ""
	actual, found = tilemap.Get(-1, 5)
	if found {
		t.Errorf("Expected %#v got %#v", false, found)
	}

	expected = ""
	actual, found = tilemap.Get(3, -2)
	if found {
		t.Errorf("Expected %#v got %#v", false, found)
	}

	expectedX := 1
	expectedY := 1
	actualX, actualY := tilemap.Find("S")
	if expectedX != actualX || expectedY != actualY {
		t.Errorf(
			"Expected %#v, %#v got %#v, %#v",
			expectedX,
			expectedY,
			actualX,
			actualY,
		)
	}

	expectedX = 0
	expectedY = 0
	actualX, actualY = tilemap.Find(".")
	if expectedX != actualX || expectedY != actualY {
		t.Errorf(
			"Expected %#v, %#v got %#v, %#v",
			expectedX,
			expectedY,
			actualX,
			actualY,
		)
	}

	expectedX = -1
	expectedY = -1
	actualX, actualY = tilemap.Find("F")
	if expectedX != actualX || expectedY != actualY {
		t.Errorf(
			"Expected %#v, %#v got %#v, %#v",
			expectedX,
			expectedY,
			actualX,
			actualY,
		)
	}
}
