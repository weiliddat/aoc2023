package day10_test

import (
	"aoc2023/aoc_util"
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
	actual, err := day10.Part01(testInput01)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	expected = "8"
	actual, err = day10.Part01(testInput02)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

var testInput03 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`

var testInput04 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

var testInput05 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`

func TestPart02(t *testing.T) {
	expected := "4"
	actual, err := day10.Part02(testInput03)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	expected = "8"
	actual, err = day10.Part02(testInput04)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	expected = "10"
	actual, err = day10.Part02(testInput05)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPathArea(t *testing.T) {
	// square of area 1
	path := day10.Path{
		aoc_util.Tile{"", 0, 0},
		aoc_util.Tile{"", 1, 0},
		aoc_util.Tile{"", 1, 1},
		aoc_util.Tile{"", 0, 1},
	}
	expected := 1.0
	actual := path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}

	// triangle of area 0.5
	path = day10.Path{
		aoc_util.Tile{"", 0, 0},
		aoc_util.Tile{"", 1, 0},
		aoc_util.Tile{"", 1, 1},
	}
	expected = 0.5
	actual = path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}

	// large triangle
	path = day10.Path{
		aoc_util.Tile{"", 0, 0},
		aoc_util.Tile{"", 1, 0},
		aoc_util.Tile{"", 2, 0},
		aoc_util.Tile{"", 2, 1},
		aoc_util.Tile{"", 2, 2},
		aoc_util.Tile{"", 1, 1},
	}
	expected = 2.0
	actual = path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}

	// polygon
	path = day10.Path{
		aoc_util.Tile{"", 4, 0},
		aoc_util.Tile{"", 4, 1},
		aoc_util.Tile{"", 4, 2},
		aoc_util.Tile{"", 3, 3},
		aoc_util.Tile{"", 2, 4},
		aoc_util.Tile{"", 1, 5},
		aoc_util.Tile{"", 0, 2},
		aoc_util.Tile{"", 2, 1},
	}
	expected = 10.
	actual = path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}
}

func TestPathInternalPoints(t *testing.T) {
	// square of area 1
	path := day10.Path{
		aoc_util.Tile{"", 0, 0},
		aoc_util.Tile{"", 1, 0},
		aoc_util.Tile{"", 1, 1},
		aoc_util.Tile{"", 0, 1},
	}
	expected := 0
	actual := path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}

	// triangle of area 0.5
	path = day10.Path{
		aoc_util.Tile{"", 0, 0},
		aoc_util.Tile{"", 1, 0},
		aoc_util.Tile{"", 1, 1},
	}
	expected = 0
	actual = path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}

	// large triangle
	path = day10.Path{
		aoc_util.Tile{"", 0, 0},
		aoc_util.Tile{"", 1, 0},
		aoc_util.Tile{"", 2, 0},
		aoc_util.Tile{"", 2, 1},
		aoc_util.Tile{"", 2, 2},
		aoc_util.Tile{"", 1, 1},
	}
	expected = 0
	actual = path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}

	// polygon
	path = day10.Path{
		aoc_util.Tile{"", 4, 0},
		aoc_util.Tile{"", 4, 1},
		aoc_util.Tile{"", 4, 2},
		aoc_util.Tile{"", 3, 3},
		aoc_util.Tile{"", 2, 4},
		aoc_util.Tile{"", 1, 5},
		aoc_util.Tile{"", 0, 2},
		aoc_util.Tile{"", 2, 1},
	}
	expected = 7
	actual = path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}
}
