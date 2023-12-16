package day13

import (
	"fmt"
	"testing"
)

var testInput01 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

func TestPart01(t *testing.T) {
	expected := "405"
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

func TestIsReflected(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected bool
	}{
		{"1", "123", true},
		{"21", "123", true},
		{"321", "123", true},
		{"331", "123", false},
		{"321", "133", false},
		{"31", "134", true},
		{"5431", "134", true},
		{"", "134", true},
		{"", "", true},
		{"a", "", true},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%s <> %s", tC.a, tC.b), func(t *testing.T) {
			actual := isReflected(tC.a, tC.b)

			if tC.expected != actual {
				t.Errorf("Expected %v got %v", tC.expected, actual)
			}
		})
	}
}
