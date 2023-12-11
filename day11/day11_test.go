package day11_test

import (
	"aoc2023/aoc_util"
	"aoc2023/day11"
	"reflect"
	"testing"
)

var testInput01 = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func TestPart01(t *testing.T) {
	expected := "374"
	actual, err := day11.Part01(testInput01)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "82000210"
	actual, err := day11.Part02(testInput01)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestIsAll(t *testing.T) {
	expected := true
	actual := day11.IsAll(".....", ".")
	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = false
	actual = day11.IsAll("..#..", ".")
	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}
}

func TestParseAndExpand(t *testing.T) {
	expected := aoc_util.SplitLines(`....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
`)

	actual := day11.ParseAndExpand(testInput01)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

}

func TestFindDistanceBetweenGalaxies(t *testing.T) {
	expected := 374
	actual := day11.FindDistBetweenGalaxies(testInput01, 2)

	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = 1030
	actual = day11.FindDistBetweenGalaxies(testInput01, 10)

	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	expected = 8410
	actual = day11.FindDistBetweenGalaxies(testInput01, 100)

	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day11.Part01(day11.Input)
	}
}

func BenchmarkPart02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day11.Part02(day11.Input)
	}
}
