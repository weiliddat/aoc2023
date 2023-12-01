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

	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	replacedInput := ""
	for _, line := range lines {
		replacedInput += replacer.Replace(line)
		replacedInput += "\n"
	}

	return Part01(replacedInput)
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
