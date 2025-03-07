package day09

import (
	"aoc2023/aoc_util"
	_ "embed"
	"math"
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
	histories, err := ParseHistories(&input)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, history := range histories {
		lastValue := Extrapolate(history, false)
		sum += lastValue
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	histories, err := ParseHistories(&input)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, history := range histories {
		lastValue := Extrapolate(history, true)
		sum += lastValue
	}

	return strconv.Itoa(sum), nil
}

func Extrapolate(history []int, backward bool) int {
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

	if backward {
		for i := range rows {
			slices.Reverse(rows[i])
		}
	}

	for depth, row := range rows {
		if depth == 0 {
			rows[depth] = append(row, row[0])
		} else {
			prevRow := rows[depth-1]
			prevRowLast := prevRow[len(prevRow)-1]
			rowLast := row[len(prevRow)-1]
			nextValue := rowLast + prevRowLast
			if backward {
				nextValue = rowLast - prevRowLast
			}
			rows[depth] = append(row, nextValue)
		}
	}

	lastRow := rows[len(rows)-1]
	lastValue := lastRow[len(lastRow)-1]

	return lastValue
}

func ParseHistories(input *string) ([][]int, error) {
	histories := [][]int{}

	lines := aoc_util.SplitLines(*input)

	for _, line := range lines {
		history := []int{}

		for _, numberText := range strings.Split(line, " ") {
			number, err := strconv.Atoi(numberText)
			if err != nil {
				return histories, err
			}
			history = append(history, number)
		}

		histories = append(histories, history)
	}

	return histories, nil
}

func Part01Lagrange(input string) (string, error) {
	histories, err := ParseHistories(&input)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, history := range histories {
		lastValue := Lagrange(history)
		sum += lastValue
	}

	return strconv.Itoa(sum), nil
}

// https://en.wikipedia.org/wiki/Lagrange_polynomial
// https://mathworld.wolfram.com/LagrangeInterpolatingPolynomial.html
func Lagrange(history []int) int {
	x := len(history)

	sum := 0

	for j, y := range history {
		sum += Pj(j, x, y)
	}

	return sum
}

func Pj(j, n, y int) int {
	prod := float64(y)

	for k := 0; k < n; k++ {
		if k != j {
			prod = prod * float64(n-k) / float64(j-k)
		}
	}

	return int(math.Round(prod))
}
