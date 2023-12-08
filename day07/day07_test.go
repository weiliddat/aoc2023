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
	expected := "5905"
	actual, _ := day07.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	testInput := `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41
`

	expected = "6839"
	actual, _ = day07.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestParseHand(t *testing.T) {
	testCases := []struct {
		input    string
		useJoker bool
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

		// J as Jokers
		{
			input:    "32T3K 765",
			useJoker: true,
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
			input:    "T55J5 684",
			useJoker: true,
			expected: day07.Hand{
				Cards: "T55J5",
				Bid:   684,
				Counts: map[string]int{
					"5": 3,
					"T": 1,
					"J": 1,
				},
				Type: day07.Fours,
			},
		},
		{
			input:    "KK677 28",
			useJoker: true,
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
			input:    "KTJJT 220",
			useJoker: true,
			expected: day07.Hand{
				Cards: "KTJJT",
				Bid:   220,
				Counts: map[string]int{
					"T": 2,
					"J": 2,
					"K": 1,
				},
				Type: day07.Fours,
			},
		},
		{
			input:    "QQQJA 483",
			useJoker: true,
			expected: day07.Hand{
				Cards: "QQQJA",
				Bid:   483,
				Counts: map[string]int{
					"J": 1,
					"Q": 3,
					"A": 1,
				},
				Type: day07.Fours,
			},
		},
		{
			input: "8J29Q 523",
			expected: day07.Hand{
				Cards: "8J29Q",
				Bid:   523,
				Counts: map[string]int{
					"2": 1,
					"8": 1,
					"9": 1,
					"J": 1,
					"Q": 1,
				},
				Type: day07.HighCard,
			},
		},
		{
			input:    "8J29Q 523",
			useJoker: true,
			expected: day07.Hand{
				Cards: "8J29Q",
				Bid:   523,
				Counts: map[string]int{
					"2": 1,
					"8": 1,
					"9": 1,
					"J": 1,
					"Q": 1,
				},
				Type: day07.OnePair,
			},
		},
		{
			input:    "TAJA2 998",
			useJoker: true,
			expected: day07.Hand{
				Cards: "TAJA2",
				Bid:   998,
				Counts: map[string]int{
					"2": 1,
					"T": 1,
					"J": 1,
					"A": 2,
				},
				Type: day07.Threes,
			},
		},
		{
			input: "JJJJJ 145",
			expected: day07.Hand{
				Cards: "JJJJJ",
				Bid:   145,
				Counts: map[string]int{
					"J": 5,
				},
				Type: day07.Fives,
			},
		},
		{
			input:    "JJJJJ 145",
			useJoker: true,
			expected: day07.Hand{
				Cards: "JJJJJ",
				Bid:   145,
				Counts: map[string]int{
					"J": 5,
				},
				Type: day07.Fives,
			},
		},
		{
			input:    "JQJ55 793",
			useJoker: true,
			expected: day07.Hand{
				Cards: "JQJ55",
				Bid:   793,
				Counts: map[string]int{
					"J": 2,
					"Q": 1,
					"5": 2,
				},
				Type: day07.Fours,
			},
		},
		{
			input:    "9JJJ9 500",
			useJoker: true,
			expected: day07.Hand{
				Cards: "9JJJ9",
				Bid:   500,
				Counts: map[string]int{
					"J": 3,
					"9": 2,
				},
				Type: day07.Fives,
			},
		},
		{
			input:    "AA55J 500",
			useJoker: true,
			expected: day07.Hand{
				Cards: "AA55J",
				Bid:   500,
				Counts: map[string]int{
					"J": 1,
					"A": 2,
					"5": 2,
				},
				Type: day07.FullHouse,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {

			actual, err := day07.ParseHand(&tC.input, tC.useJoker)

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
