package day12

import (
	"aoc2023/aoc_util"
	_ "embed"
	"fmt"
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
		possible, err := findPossible(line)
		if err != nil {
			return "", err
		}
		sum += possible
	}

	return "", nil
}

func findPossible(input string) (int, error) {
	conditionText, springsText, _ := strings.Cut(input, " ")

	damagedSprings := []int{}
	for _, springText := range strings.Split(springsText, ",") {
		spring, err := strconv.Atoi(springText)
		if err != nil {
			return 0, err
		}
		damagedSprings = append(damagedSprings, spring)
	}

	fmt.Println(conditionText, damagedSprings)

	return 0, nil
}

func fit(record string, damagedSprings []int, depth string) int {
	total := 0

	// if remaining springs cannot possibly fit in the record by length
	if sum(damagedSprings) > len(record)-strings.Count(record, ".") {
		return total
	}

	for i := 0; i < len(record); i++ {
		possible := 0
		spring := damagedSprings[0]

		// if the current spring exceeds the remaining record
		if spring > len(record[i:]) {
			break
		}

		here := record[i : i+spring]
		if !strings.Contains(here, ".") {
			fmt.Println(depth, "found", record, "at", i, "len", spring)
			possible++

			remainingSprings := damagedSprings[1:]

			// if we have no remaining springs to fit
			if len(remainingSprings) == 0 {
				total += possible
				break
			}

			// need to find next ?/. after current spring fit
			nextSeparator := strings.IndexAny(record[i+spring:], "?.")
			if nextSeparator == -1 {
				// there are remaining springs but we don't have a separator
				// impossible branch
				possible = 0
				break
			}
			fmt.Println(depth, "foundSp", record[i+spring:], "at", nextSeparator)
			nextSeparator = nextSeparator + i + spring + 1

			// offset of 1 needed after fitting spring as separator
			remainingRecordLength := len(record) - nextSeparator

			// if we have remaining springs but they don't fit
			if sum(remainingSprings) > remainingRecordLength {
				// this branch is impossible
				possible = 0
				break
			}

			next := fit(record[nextSeparator:], remainingSprings, depth+"  ")
			possible = possible * next
			total += possible
		}
	}

	fmt.Println(depth, "fit result", record, damagedSprings, total)
	return total
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Part02(input string) (string, error) {
	return "", nil
}
