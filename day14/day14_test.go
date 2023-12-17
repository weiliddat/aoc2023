package day14

import (
	"aoc2023/aoc_util"
	"reflect"
	"slices"
	"strings"
	"testing"
)

var testInput01 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

func TestPart01(t *testing.T) {
	expected := "136"
	actual, err := Part01(testInput01)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

var testInput02 = ``

func TestPart02(t *testing.T) {
	expected := ""
	actual, err := Part02(testInput02)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestTiltLeft(t *testing.T) {
	expected := aoc_util.IntoColumns(aoc_util.SplitLines(`OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`))
	input := aoc_util.IntoColumns(aoc_util.SplitLines(testInput01))
	actual := slices.Clone(input)
	tiltLeft(&actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected\n%v\n\ngot\n\n%v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
	}
}
