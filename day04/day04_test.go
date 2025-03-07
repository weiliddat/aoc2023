package day04_test

import (
	"aoc2023/day04"
	"reflect"
	"testing"
)

var testInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func TestPart01(t *testing.T) {
	expected := "13"
	actual, err := day04.Part01(testInput)

	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "30"
	actual, err := day04.Part02(testInput)

	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestParseScratchCards(t *testing.T) {
	testInput := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := day04.ScratchCard{
		ID: 1,
		WinningNumbers: map[int]bool{
			41: true,
			48: true,
			83: true,
			86: true,
			17: true,
		},
		CardNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
		Matches:     []int{83, 86, 17, 48},
	}
	actual, err := day04.ParseScratchCard(testInput)

	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}

func BenchmarkParseScratchCard(b *testing.B) {
	testInput := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

	for i := 0; i < b.N; i++ {
		day04.ParseScratchCard(testInput)
	}
}
