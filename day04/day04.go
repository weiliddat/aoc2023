package day04

import (
	"aoc2023/aoc_util"
	_ "embed"
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

func Part01(input string) (string, error) {
	lines := aoc_util.SplitLines(input)

	sum := 0
	for _, line := range lines {
		card, err := ParseScratchCard(line)
		if err != nil {
			return "", err
		}

		cardMatches := len(card.Matches)

		if cardMatches > 0 {
			sum += int(1) << (uint64(cardMatches) - 1)
		}
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	lines := aoc_util.SplitLines(input)

	sum := 0
	cardCounts := map[int]*struct {
		count   int
		matches int
	}{}

	// initial scratchcards
	for _, line := range lines {
		card, err := ParseScratchCard(line)
		if err != nil {
			return "", err
		}
		cardCounts[card.ID] = &struct {
			count   int
			matches int
		}{
			count:   1,
			matches: len(card.Matches),
		}
	}

	// loop in order
	for id := 1; id <= len(cardCounts); id++ {
		count := cardCounts[id]
		for m := 1; m <= count.matches; m++ {
			card := cardCounts[id+m]
			card.count += count.count
		}
	}

	for _, count := range cardCounts {
		sum += count.count
	}

	return strconv.Itoa(sum), nil
}

type ScratchCard struct {
	ID             int
	WinningNumbers map[int]bool
	CardNumbers    []int
	Matches        []int
}

func ParseScratchCard(input string) (ScratchCard, error) {
	card := ScratchCard{
		WinningNumbers: map[int]bool{},
		CardNumbers:    []int{},
		Matches:        []int{},
	}

	cardAndNumberText, numberList, _ := strings.Cut(input, ": ")
	_, cardNumberText, _ := strings.Cut(cardAndNumberText, " ")
	cardNumber, err := strconv.Atoi(strings.TrimSpace(cardNumberText))
	if err != nil {
		return card, err
	}
	card.ID = cardNumber

	winningNumbersText, cardNumbersText, _ := strings.Cut(numberList, " | ")
	for _, winningNumberText := range strings.Split(winningNumbersText, " ") {
		if len(winningNumberText) > 0 {
			winningNumber, _ := strconv.Atoi(winningNumberText)
			card.WinningNumbers[winningNumber] = true
		}
	}

	for _, cardNumberText := range strings.Split(cardNumbersText, " ") {
		if len(cardNumberText) > 0 {
			cardNumber, _ := strconv.Atoi(cardNumberText)
			card.CardNumbers = append(card.CardNumbers, cardNumber)
			if card.WinningNumbers[cardNumber] {
				card.Matches = append(card.Matches, cardNumber)
			}
		}
	}

	return card, nil
}
