package day02_test

import (
	"aoc2023/day02"
	"reflect"
	"testing"
)

var testInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func TestPart01(t *testing.T) {
	expected := "8"
	actual, _ := day02.Part01(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestParseGameText(t *testing.T) {
	text := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`
	actual, err := day02.ParseGameText(text)

	if err != nil {
		t.Fatal(err)
	}

	sets := []day02.CubeSet{
		{
			Red:  4,
			Blue: 3,
		},
		{
			Red:   1,
			Green: 2,
			Blue:  6,
		},
		{
			Green: 2,
		},
	}

	expected := day02.Game{
		Round: 1,
		Sets:  sets,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}
