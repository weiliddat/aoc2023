package day05_test

import (
	"aoc2023/day05"
	"reflect"
	"testing"
)

var testInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestPart01(t *testing.T) {
	expected := "35"
	actual, _ := day05.Part01(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "46"
	actual, _ := day05.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4
`
	expected := day05.Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: map[string]day05.AlmanacMap{
			"seed": {
				Src:  "seed",
				Dest: "soil",
				Ranges: []day05.AlmanacMapRange{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			"fertilizer": {
				Src:  "fertilizer",
				Dest: "water",
				Ranges: []day05.AlmanacMapRange{
					{49, 53, 8},
					{0, 11, 42},
					{42, 0, 7},
					{57, 7, 4},
				},
			},
		},
	}

	actual, err := day05.ParseInput(input)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}

func TestAlmanacLookup(t *testing.T) {
	almanac := day05.Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: map[string]day05.AlmanacMap{
			"seed": {
				Src:  "seed",
				Dest: "soil",
				Ranges: []day05.AlmanacMapRange{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			"soil": {
				Src:  "soil",
				Dest: "fertilizer",
				Ranges: []day05.AlmanacMapRange{
					{0, 15, 37},
					{37, 52, 2},
					{39, 0, 15},
				},
			},
			"fertilizer": {
				Src:  "fertilizer",
				Dest: "water",
				Ranges: []day05.AlmanacMapRange{
					{49, 53, 8},
					{0, 11, 42},
					{42, 0, 7},
					{57, 7, 4},
				},
			},
			"water": {
				Src:  "water",
				Dest: "light",
				Ranges: []day05.AlmanacMapRange{
					{88, 18, 7},
					{18, 25, 70},
				},
			},
			"light": {
				Src:  "light",
				Dest: "temperature",
				Ranges: []day05.AlmanacMapRange{
					{45, 77, 23},
					{81, 45, 19},
					{68, 64, 13},
				},
			},
			"temperature": {
				Src:  "temperature",
				Dest: "humidity",
				Ranges: []day05.AlmanacMapRange{
					{0, 69, 1},
					{1, 0, 69},
				},
			},
			"humidity": {
				Src:  "humidity",
				Dest: "location",
				Ranges: []day05.AlmanacMapRange{
					{60, 56, 37},
					{56, 93, 4},
				},
			},
		},
	}

	expectedName, expectedNum := "soil", 81
	actualName, actualNum := almanac.Lookup("seed", 79)

	if expectedName != actualName || expectedNum != actualNum {
		t.Errorf("Expected %s, %d got %s, %d", expectedName, expectedNum, actualName, actualNum)
	}

	expectedName, expectedNum = "fertilizer", 81
	actualName, actualNum = almanac.Lookup("soil", 81)

	if expectedName != actualName || expectedNum != actualNum {
		t.Errorf("Expected %s, %d got %s, %d", expectedName, expectedNum, actualName, actualNum)
	}

	expectedName, expectedNum = "light", 74
	actualName, actualNum = almanac.Lookup("water", 81)

	if expectedName != actualName || expectedNum != actualNum {
		t.Errorf("Expected %s, %d got %s, %d", expectedName, expectedNum, actualName, actualNum)
	}

	expectedName, expectedNum = "temperature", 45
	actualName, actualNum = almanac.Lookup("light", 77)

	if expectedName != actualName || expectedNum != actualNum {
		t.Errorf("Expected %s, %d got %s, %d", expectedName, expectedNum, actualName, actualNum)
	}
}
