package day09

import (
	"aoc2023/aoc_util"
	_ "embed"
	"slices"
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
	histories := [][]int{}

	lines := aoc_util.SplitLines(input)

	for _, line := range lines {
		history := []int{}

		for _, numberText := range strings.Split(line, " ") {
			number, err := strconv.Atoi(numberText)
			if err != nil {
				return "", err
			}
			history = append(history, number)
		}

		histories = append(histories, history)
	}

	sum := 0
	for _, history := range histories {
		lastValue := Extrapolate(history)
		sum += lastValue
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	return "", nil
}

func Extrapolate(history []int) int {
	rows := [][]int{
		history,
	}

	nextRow := rows[0]

	for {
		row := []int{}

		for i := 0; i < len(nextRow)-1; i++ {
			diff := (nextRow)[i+1] - (nextRow)[i]
			row = append(row, diff)
		}

		rows = append(rows, row)

		if slices.Max(row) == slices.Min(row) {
			break
		} else {
			nextRow = row
		}
	}

	slices.Reverse(rows)

	for depth, row := range rows {
		if depth == 0 {
			rows[depth] = append(row, row[0])
		} else {
			prevRow := rows[depth-1]
			prevRowLast := prevRow[len(prevRow)-1]
			rowLast := row[len(prevRow)-1]
			rows[depth] = append(row, rowLast+prevRowLast)
		}
	}

	lastRow := rows[len(rows)-1]
	lastValue := lastRow[len(lastRow)-1]

	return lastValue
}
