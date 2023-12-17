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

// transpose rows into columns
// assumes all rows have the same length
func Transpose(rows []string) []string {
	rowLength := len(rows[0])
	cols := make([]string, 0, rowLength)

	for ci := 0; ci < rowLength; ci++ {
		cs := strings.Builder{}
		for _, row := range rows {
			cs.WriteByte(row[ci])
		}
		cols = append(cols, cs.String())
	}

	return cols
}

// split blocks
func SplitBlocks(input string) []string {
	blocks := strings.Split(input, "\n\n")

	for bi := range blocks {
		blocks[bi] = strings.TrimSpace(blocks[bi])
	}

	return blocks
}

func StringToNums(input string, sep string) ([]int, error) {
	numbers := []int{}

	for _, numStr := range strings.Split(input, sep) {
		spring, err := strconv.Atoi(numStr)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, spring)
	}

	return numbers, nil
}
