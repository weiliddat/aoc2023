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

var JokerCardLookup = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

func Part01(input string) (string, error) {
	lines := aoc_util.SplitLines(input)

	hands := []Hand{}
	for _, line := range lines {
		hand, err := ParseHand(&line, false)
		if err != nil {
			return "", err
		}
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, SortHands)

	sum := 0
	for index, hand := range hands {
		sum += (index + 1) * hand.Bid
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	lines := aoc_util.SplitLines(input)

	hands := []Hand{}
	for _, line := range lines {
		hand, err := ParseHand(&line, true)
		if err != nil {
			return "", err
		}
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, SortHandsJoker)

	sum := 0
	for index, hand := range hands {
		sum += (index + 1) * hand.Bid
	}

	return strconv.Itoa(sum), nil
}

func ParseHand(input *string, useJoker bool) (Hand, error) {
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

	counts := []struct {
		count int
		card  string
	}{}
	for card, count := range hand.Counts {
		counts = append(counts, struct {
			count int
			card  string
		}{card: card, count: count})
	}

	slices.SortFunc(counts, func(a, b struct {
		count int
		card  string
	}) int {
		return b.count - a.count
	})

	if useJoker {
		jokers, found := hand.Counts["J"]
		if found {
			usedJoker := false
			if counts[0].card != "J" {
				counts[0].count += jokers
				usedJoker = true
			} else if len(counts) > 1 {
				counts[1].count += jokers
				usedJoker = true
			}

			if usedJoker {
				for index, count := range counts {
					if count.card == "J" {
						counts[index].count -= jokers
					}
				}
			}
		}
	}

	slices.SortFunc(counts, func(a, b struct {
		count int
		card  string
	}) int {
		return b.count - a.count
	})

	if counts[0].count == 5 {
		hand.Type = Fives
	} else if counts[0].count == 4 {
		hand.Type = Fours
	} else if counts[0].count == 3 {
		if counts[1].count == 2 {
			hand.Type = FullHouse
		} else {
			hand.Type = Threes
		}
	} else if counts[0].count == 2 {
		if counts[1].count == 2 {
			hand.Type = TwoPair
		} else {
			hand.Type = OnePair
		}
	} else {
		hand.Type = HighCard
	}

	return hand, nil
}

func SortHands(a, b Hand) int {
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
}

func SortHandsJoker(a, b Hand) int {
	if a.Type != b.Type {
		return int(a.Type - b.Type)
	}

	diff := 0
	for i := 0; i < len(a.Cards); i++ {
		aValue := JokerCardLookup[a.Cards[i:i+1]]
		bValue := JokerCardLookup[b.Cards[i:i+1]]
		if aValue != bValue {
			diff = aValue - bValue
			break
		}
	}
	return diff
}
