package aoc_util

import (
	"slices"
	"strconv"
	"strings"
)

// split and removes empty lines
func SplitLines(input string) []string {
	lines := strings.Split(input, "\n")
	lines = slices.DeleteFunc(lines, func(line string) bool {
		return len(line) == 0
	})
	return lines
}

func StringToNums(input string, sep string) ([]int, error) {
	numbers := []int{}

	for _, numStr := range strings.Split(input, ",") {
		spring, err := strconv.Atoi(numStr)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, spring)
	}

	return numbers, nil
}
