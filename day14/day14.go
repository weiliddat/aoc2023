package day14

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
	rows := aoc_util.SplitLines(input)
	cols := aoc_util.Transpose(rows)

	tiltLeft(&cols)

	sum := 0
	for _, col := range cols {
		for i, r := range col {
			if r == 'O' {
				distFromSouth := len(cols) - i
				sum += distFromSouth
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	rows := aoc_util.SplitLines(input)

	cycle(&rows)

	sum := 0

	return strconv.Itoa(sum), nil
}

func tiltLeft(rows *[]string) {
	for i := range *rows {
		(*rows)[i] = moveRocksLeft((*rows)[i])
	}
}

func tiltRight(rows *[]string) {
	for i := range *rows {
		(*rows)[i] = moveRocksRight((*rows)[i])
	}
}

func moveRocksLeft(s string) string {
	rounds := []int{}
	borders := []int{-1} // dish edge is also a border
	numRoundsFromPrevEdge := 0

	for i := range s {
		switch s[i : i+1] {
		case "O":
			prevEdge := borders[len(borders)-1]
			numRoundsFromPrevEdge++
			rounds = append(rounds, prevEdge+numRoundsFromPrevEdge)
		case "#":
			borders = append(borders, i)
			numRoundsFromPrevEdge = 0
		}
	}

	newS := strings.Builder{}
	for i := range s {
		if slices.Contains(rounds, i) {
			newS.WriteRune('O')
		} else if slices.Contains(borders, i) {
			newS.WriteRune('#')
		} else {
			newS.WriteRune('.')
		}
	}

	return newS.String()
}

func moveRocksRight(s string) string {
	rounds := []int{}
	borders := []int{len(s)} // dish edge is also a border
	numRoundsFromPrevEdge := 0

	for i := range s {
		i = len(s) - 1 - i
		switch s[i : i+1] {
		case "O":
			prevEdge := borders[len(borders)-1]
			numRoundsFromPrevEdge--
			rounds = append(rounds, prevEdge+numRoundsFromPrevEdge)
		case "#":
			borders = append(borders, i)
			numRoundsFromPrevEdge = 0
		}
	}

	newS := strings.Builder{}
	for i := range s {
		if slices.Contains(rounds, i) {
			newS.WriteRune('O')
		} else if slices.Contains(borders, i) {
			newS.WriteRune('#')
		} else {
			newS.WriteRune('.')
		}
	}

	return newS.String()
}

func cycle(rows *[]string) {
	// north
	cols := aoc_util.Transpose(*rows)
	tiltLeft(&cols)

	// west
	*rows = aoc_util.Transpose(cols)
	tiltLeft(rows)

	// south
	cols = aoc_util.Transpose(*rows)
	tiltRight(&cols)

	// east
	*rows = aoc_util.Transpose(cols)
	tiltRight(rows)
}
