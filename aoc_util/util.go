package aoc_util

import (
	"slices"
	"strings"
)

func SplitLines(input string) []string {
	lines := strings.Split(input, "\n")
	lines = slices.DeleteFunc(lines, func(line string) bool {
		return len(line) == 0
	})
	return lines
}
