package day14

import (
	"aoc2023/aoc_util"
	"reflect"
	"slices"
	"strings"
	"testing"
)

var testInput = `O....#....
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
	actual, err := Part01(testInput)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "64"
	actual, err := Part02(testInput)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestTiltLeft(t *testing.T) {
	expected := aoc_util.Transpose(aoc_util.SplitLines(`OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`))
	input := aoc_util.Transpose(aoc_util.SplitLines(testInput))
	actual := slices.Clone(input)
	tilt(&actual, false)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected\n%v\n\ngot\n\n%v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
	}
}

func TestMoveRocksRight(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{".....#....", ".....#...."},
		{".....#..O.", ".....#...O"},
		{"...O.#....", "....O#...."},
		{"...O......", ".........O"},
		{"...O...O..", "........OO"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			actual := moveRocksRight(tC.input)
			if tC.expected != actual {
				t.Errorf("Expected %s got %s", tC.expected, actual)
			}
		})
	}
}

func TestCycle(t *testing.T) {
	expected := aoc_util.SplitLines(`.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`)
	actual := aoc_util.SplitLines(testInput)
	cycle(&actual)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected\n%v\n\ngot\n\n%v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
	}

	expected = aoc_util.SplitLines(`.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`)
	cycle(&actual)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected\n%v\n\ngot\n\n%v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
	}

	expected = aoc_util.SplitLines(`.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`)
	cycle(&actual)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected\n%v\n\ngot\n\n%v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
	}
}

func BenchmarkCycle(b *testing.B) {
	actual := aoc_util.SplitLines(testInput)

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			cycle(&actual)
		}
	}
}
