package day07_test

import (
	"aoc2023/day07"
	"reflect"
	"testing"
)

var testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestPart01(t *testing.T) {
	expected := "6440"
	actual, _ := day07.Part01(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := ""
	actual, _ := day07.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestParseHand(t *testing.T) {
	testCases := []struct {
		input    string
		expected day07.Hand
	}{
		{
			input: "32T3K 765",
			expected: day07.Hand{
				Cards: "32T3K",
				Bid:   765,
				Counts: map[string]int{
					"2": 1,
					"3": 2,
					"T": 1,
					"K": 1,
				},
				Type: day07.OnePair,
			},
		},
		{
			input: "T55J5 684",
			expected: day07.Hand{
				Cards: "T55J5",
				Bid:   684,
				Counts: map[string]int{
					"5": 3,
					"T": 1,
					"J": 1,
				},
				Type: day07.Threes,
			},
		},
		{
			input: "KK677 28",
			expected: day07.Hand{
				Cards: "KK677",
				Bid:   28,
				Counts: map[string]int{
					"6": 1,
					"7": 2,
					"K": 2,
				},
				Type: day07.TwoPair,
			},
		},
		{
			input: "KTJJT 220",
			expected: day07.Hand{
				Cards: "KTJJT",
				Bid:   220,
				Counts: map[string]int{
					"T": 2,
					"J": 2,
					"K": 1,
				},
				Type: day07.TwoPair,
			},
		},
		{
			input: "QQQJA 483",
			expected: day07.Hand{
				Cards: "QQQJA",
				Bid:   483,
				Counts: map[string]int{
					"J": 1,
					"Q": 3,
					"A": 1,
				},
				Type: day07.Threes,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {

			actual, err := day07.ParseHand(&tC.input)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(tC.expected, actual) {
				t.Errorf(
					"Expected %#v got %#v",
					tC.expected,
					actual,
				)
			}
		})
	}
}
