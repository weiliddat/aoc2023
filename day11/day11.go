package day11

import (
	"aoc2023/aoc_util"
	_ "embed"
	"slices"
	"strconv"
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
	rows := ParseAndExpand(input)

	type point struct {
		x, y int
	}

	galaxies := []point{}

	for y := range rows {
		for x := range rows[y] {
			here := rows[y][x : x+1]
			if here != "." {
				galaxies = append(galaxies, point{x, y})
			}
		}
	}

	sum := 0

	for i := range galaxies {
		curr := galaxies[i]
		for _, other := range galaxies[i+1:] {
			sum = sum + abs(curr.x-other.x) + abs(curr.y-other.y)
		}
	}

	return strconv.Itoa(sum), nil
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func ParseAndExpand(input string) []string {
	rows := aoc_util.SplitLines(input)

	columns := make([]string, len(rows[0]))

	for _, row := range rows {
		for x := 0; x < len(row); x++ {
			columns[x] = columns[x] + row[x:x+1]
		}
	}

	colsToInsert := []int{}
	for x, col := range columns {
		if IsAll(col, ".") {
			colsToInsert = append(colsToInsert, x)
		}
	}

	maxY := len(rows)
	for i := 0; i < maxY; i++ {
		for j := 0; j < len(colsToInsert); j++ {
			x := colsToInsert[j] + len(colsToInsert[:j])
			rows[i] = rows[i][0:x] + "." + rows[i][x:]
		}
		if IsAll(rows[i], ".") {
			rows = slices.Insert(rows, i, rows[i])
			i++
			maxY = len(rows)
		}
	}

	return rows
}

func IsAll(s, c string) bool {
	for i := 0; i < len(s); i++ {
		if s[i:i+1] != c {
			return false
		}
	}

	return true
}

func Part02(input string) (string, error) {
	return "", nil
}
