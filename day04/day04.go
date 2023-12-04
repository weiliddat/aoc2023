package day04

import (
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
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		card := ParseScratchCard(line)

		cardMatches := 0
		for _, cardNumber := range card.CardNumbers {
			if card.WinningNumbers[cardNumber] {
				cardMatches += 1
			}
		}

		if cardMatches > 0 {
			sum += int(1) << (uint64(cardMatches) - 1)
		}
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	return "", nil
}

type ScratchCard struct {
	WinningNumbers map[int]bool
	CardNumbers    []int
}

func ParseScratchCard(input string) ScratchCard {
	card := ScratchCard{
		WinningNumbers: map[int]bool{},
		CardNumbers:    []int{},
	}

	_, numberList, _ := strings.Cut(input, ": ")
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
		}
	}

	return card
}
