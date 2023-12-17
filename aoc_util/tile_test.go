package aoc_util_test

import (
	"aoc2023/aoc_util"
	"aoc2023/day10"
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkNewTileMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := aoc_util.NewTileMap(day10.Input)
		t.Get(1, 1)
		t.Get(2, 2)
	}
}

func TestString(t *testing.T) {
	tilemap := aoc_util.NewTileMap(day10.Input)
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

var testInput01 = `.....
.S-7.
.|.|.
.L-J.
.....
`

func TestTileMapGet(t *testing.T) {
	tilemap := aoc_util.NewTileMap(testInput01)

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
	tilemap := aoc_util.NewTileMap(testInput01)

	expected := aoc_util.Tile{"S", 1, 1}
	actual, found := tilemap.Find("S")

	if !found {
		t.Errorf("Expected %#v got %#v", true, found)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = aoc_util.Tile{".", 0, 0}
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
