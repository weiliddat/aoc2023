package day07

import (
	"aoc2023/aoc_util"
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var Input string

func Solve(input string) (string, string, error) {
	part01, err := Part01(input)

	if err != nil {
		return "", "", err
	}

	part02, err := Part02(input)

	if err != nil {
		return part01, "", err
	}

	return part01, part02, nil
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	Threes
	FullHouse
	Fours
	Fives
)

type Hand struct {
	Cards  string
	Bid    int
	Counts map[string]int
	Type   HandType
}

var CardLookup = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func Part01(input string) (string, error) {
	lines := aoc_util.SplitLines(input)

	hands := []Hand{}
	for _, line := range lines {
		hand, err := ParseHand(&line)
		if err != nil {
			return "", err
		}
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Type != b.Type {
			return int(a.Type - b.Type)
		}

		diff := 0
		for i := 0; i < len(a.Cards); i++ {
			aValue := CardLookup[a.Cards[i:i+1]]
			bValue := CardLookup[b.Cards[i:i+1]]
			if aValue != bValue {
				diff = aValue - bValue
				break
			}
		}
		return diff
	})

	sum := 0
	for index, hand := range hands {
		sum += (index + 1) * hand.Bid
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	return "", nil
}

func ParseHand(input *string) (Hand, error) {
	var hand Hand

	cardText, bidText, _ := strings.Cut(*input, " ")
	hand.Cards = cardText

	bid, err := strconv.Atoi(bidText)
	if err != nil {
		return hand, err
	}
	hand.Bid = bid

	hand.Counts = map[string]int{}
	for _, r := range cardText {
		char := string(r)
		hand.Counts[char]++
	}

	counts := []int{}
	for _, count := range hand.Counts {
		counts = append(counts, count)
	}
	slices.Sort(counts)
	slices.Reverse(counts)

	if counts[0] == 5 {
		hand.Type = Fives
	} else if counts[0] == 4 {
		hand.Type = Fours
	} else if counts[0] == 3 {
		if counts[1] == 2 {
			hand.Type = FullHouse
		} else {
			hand.Type = Threes
		}
	} else if counts[0] == 2 {
		if counts[1] == 2 {
			hand.Type = TwoPair
		} else {
			hand.Type = OnePair
		}
	} else {
		hand.Type = HighCard
	}

	return hand, nil
}
