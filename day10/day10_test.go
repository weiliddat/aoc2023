package day10_test

import (
	"aoc2023/day10"
	"reflect"
	"testing"
	"unsafe"
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

func BenchmarkNewTileMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := day10.NewTileMap(day10.Input)
		t.Get(1, 1)
		t.Get(2, 2)
	}
}

func TestString(t *testing.T) {
	tilemap := day10.NewTileMap(day10.Input)
	tile, _ := tilemap.Get(0, 0)

	ipd := unsafe.StringData(day10.Input)
	tmpd := unsafe.StringData(tilemap.Tiles[0])
	tpd := unsafe.StringData(tile.S)

	if ipd != tmpd || tmpd != tpd {
		t.Errorf(
			"Underlying string arrays are different\n%v %v %v\n",
			ipd, tmpd, tpd,
		)
	}
}

func TestTileMapGet(t *testing.T) {
	tilemap := day10.NewTileMap(testInput01)

	expected := "."
	actual, found := tilemap.Get(0, 0)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = "."
	actual, found = tilemap.Get(4, 4)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = "S"
	actual, found = tilemap.Get(1, 1)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != actual.S {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = "J"
	actual, found = tilemap.Get(3, 3)
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if expected != actual.S {
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
}

func TestTileMapFind(t *testing.T) {
	tilemap := day10.NewTileMap(testInput01)

	expected := day10.Tile{"S", 1, 1}
	actual, found := tilemap.Find("S")

	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = day10.Tile{".", 0, 0}
	actual, found = tilemap.Find(".")
	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	_, found = tilemap.Find("F")
	if found {
		t.Errorf("Expected %#v got %#v", false, found)
	}
}

func TestPathArea(t *testing.T) {
	// square of area 1
	path := day10.Path{
		day10.Tile{"", 0, 0},
		day10.Tile{"", 1, 0},
		day10.Tile{"", 1, 1},
		day10.Tile{"", 0, 1},
	}
	expected := 1.0
	actual := path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}

	// triangle of area 0.5
	path = day10.Path{
		day10.Tile{"", 0, 0},
		day10.Tile{"", 1, 0},
		day10.Tile{"", 1, 1},
	}
	expected = 0.5
	actual = path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}

	// large triangle
	path = day10.Path{
		day10.Tile{"", 0, 0},
		day10.Tile{"", 1, 0},
		day10.Tile{"", 2, 0},
		day10.Tile{"", 2, 1},
		day10.Tile{"", 2, 2},
		day10.Tile{"", 1, 1},
	}
	expected = 2.0
	actual = path.Area()
	if expected != actual {
		t.Errorf("Expected %f got %f", expected, actual)
	}

	// polygon
	path = day10.Path{
		day10.Tile{"", 4, 0},
		day10.Tile{"", 4, 1},
		day10.Tile{"", 4, 2},
		day10.Tile{"", 3, 3},
		day10.Tile{"", 2, 4},
		day10.Tile{"", 1, 5},
		day10.Tile{"", 0, 2},
		day10.Tile{"", 2, 1},
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
		day10.Tile{"", 0, 0},
		day10.Tile{"", 1, 0},
		day10.Tile{"", 1, 1},
		day10.Tile{"", 0, 1},
	}
	expected := 0
	actual := path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}

	// triangle of area 0.5
	path = day10.Path{
		day10.Tile{"", 0, 0},
		day10.Tile{"", 1, 0},
		day10.Tile{"", 1, 1},
	}
	expected = 0
	actual = path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}

	// large triangle
	path = day10.Path{
		day10.Tile{"", 0, 0},
		day10.Tile{"", 1, 0},
		day10.Tile{"", 2, 0},
		day10.Tile{"", 2, 1},
		day10.Tile{"", 2, 2},
		day10.Tile{"", 1, 1},
	}
	expected = 0
	actual = path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}

	// polygon
	path = day10.Path{
		day10.Tile{"", 4, 0},
		day10.Tile{"", 4, 1},
		day10.Tile{"", 4, 2},
		day10.Tile{"", 3, 3},
		day10.Tile{"", 2, 4},
		day10.Tile{"", 1, 5},
		day10.Tile{"", 0, 2},
		day10.Tile{"", 2, 1},
	}
	expected = 7
	actual = path.InternalPoints()
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}
}
