package aoc_util

import (
	"slices"
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
