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
		record, springsText, _ := strings.Cut(line, " ")

		damagedSprings, err := aoc_util.StringToNums(springsText, ",")
		if err != nil {
			return "", err
		}

		possible := fit(record, damagedSprings, "")
		sum += possible
	}

	return strconv.Itoa(sum), nil
}

func fit(record string, damagedSprings []int, depth string) int {
	total := 0

	// if there are no remaining springs and no remaining recorded springs
	if len(damagedSprings) == 0 && !strings.Contains(record, "#") {
		return 1
	}

	// if remaining springs cannot possibly fit in the record by length
	if sum(damagedSprings) > len(record)-strings.Count(record, ".") {
		return 0
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
			// cannot fit if next char is also a spring
			if len(record) > i+spring+1 {
				if record[i+spring:i+spring+1] == "#" {
					continue
				}
			}

			// cannot fit if prev char is a spring
			// usually only for first match since subsequent matches
			// we check for a separator
			if i > 0 {
				if record[i-1:i] == "#" {
					continue
				}
			}

			fmt.Println(depth, "found", record, "at", i, "len", spring)
			possible++

			remainingSprings := damagedSprings[1:]

			// if we have no remaining springs to fit
			if len(remainingSprings) == 0 {
				// if there were remaining springs in the record, it doesn't fit
				if strings.Contains(record[i+spring:], "#") {
					return 0
				}

				total += possible
				continue
			}

			// need to find next ?/. after current spring fit
			nextSeparator := strings.IndexAny(record[i+spring:], "?.")
			if nextSeparator == -1 {
				// there are remaining springs but we don't have a separator
				// impossible branch
				possible = 0
				break
			}
			// fmt.Println(depth, "foundSp", record[i+spring:], "at", nextSeparator)
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
