package day01

import (
	_ "embed"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var Input string

func Part01(input string) (string, error) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		intChars := []rune{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				intChars = append(intChars, char)
			}
		}
		firstAndLast := string(intChars[0]) + string(intChars[len(intChars)-1])

		numberInLine, err := strconv.Atoi(firstAndLast)
		if err != nil {
			return "", err
		}

		sum += numberInLine
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	lines := strings.Split(input, "\n")

	digitLookup := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0

	for _, line := range lines {
		firstValue, firstIndex, lastValue, lastIndex := 0, -1, 0, -1

		for dK, dV := range digitLookup {
			index := strings.Index(line, dK)

			if index == -1 {
				continue
			}

			if firstIndex == -1 {
				firstIndex = index
				firstValue = dV * 10
			} else if index <= firstIndex {
				firstIndex = index
				firstValue = dV * 10
			}

			index = strings.LastIndex(line, dK)

			if index == -1 {
				continue
			}

			if lastIndex == -1 {
				lastIndex = index
				lastValue = dV
			} else if index >= lastIndex {
				lastIndex = index
				lastValue = dV
			}
		}

		sum = sum + firstValue + lastValue
	}

	result := strconv.Itoa(sum)

	return result, nil
}

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
