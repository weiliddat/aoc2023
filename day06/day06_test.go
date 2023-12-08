package day06_test

import (
	"aoc2023/day06"
	"fmt"
	"reflect"
	"testing"
)

var testInput = `Time:      7  15   30
Distance:  9  40  200
`

func TestPart01(t *testing.T) {
	expected := "288"
	actual, _ := day06.Part01(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "71503"
	actual, _ := day06.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestParseRaces01(t *testing.T) {
	expected := map[int]int{
		7:  9,
		15: 40,
		30: 200,
	}

	actual, err := day06.ParseRaces01(&testInput)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}

func TestParseRaces02(t *testing.T) {
	expectedTime, expectedDist := 71530, 940200

	actualTime, actualDist, err := day06.ParseRaces02(&testInput)

	if err != nil {
		t.Error(err)
	}

	if expectedTime != actualTime || expectedDist != actualDist {
		t.Errorf(
			"Expected %d, %d got %d, %d",
			expectedTime,
			expectedDist,
			actualTime,
			actualDist,
		)
	}
}

func TestFindDistance(t *testing.T) {
	testCases := []struct {
		timeHeld  int
		totalTime int
		expected  int
	}{
		{
			timeHeld:  0,
			totalTime: 7,
			expected:  0,
		},
		{
			timeHeld:  1,
			totalTime: 7,
			expected:  6,
		},
		{
			timeHeld:  2,
			totalTime: 7,
			expected:  10,
		},
		{
			timeHeld:  3,
			totalTime: 7,
			expected:  12,
		},
		{
			timeHeld:  4,
			totalTime: 7,
			expected:  12,
		},
		{
			timeHeld:  5,
			totalTime: 7,
			expected:  10,
		},
		{
			timeHeld:  6,
			totalTime: 7,
			expected:  6,
		},
		{
			timeHeld:  7,
			totalTime: 7,
			expected:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(
			fmt.Sprintf("held %d total %d", tC.timeHeld, tC.totalTime),
			func(t *testing.T) {
				actual := day06.FindDistance(tC.timeHeld, tC.totalTime)

				if tC.expected != actual {
					t.Errorf("Expected %d got %d", tC.expected, actual)
				}
			},
		)
	}
}
