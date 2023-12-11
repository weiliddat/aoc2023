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

type point struct {
	x, y int
}

func Part01(input string) (string, error) {
	rows := ParseAndExpand(input)

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

func FindDistBetweenGalaxies(input string, emptySpaceMult int) int {
	rows := aoc_util.SplitLines(input)

	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}
	galaxies := []point{}

	for y, row := range rows {
		_, ok := emptyRows[y]
		if !ok {
			emptyRows[y] = true
		}

		for x := 0; x < len(row); x++ {
			_, ok := emptyCols[x]
			if !ok {
				emptyCols[x] = true
			}

			if row[x:x+1] == "#" {
				galaxies = append(galaxies, point{x, y})
				emptyRows[y] = false
				emptyCols[x] = false
			}
		}
	}

	sum := 0
	for i, galaxy := range galaxies {
		for _, other := range galaxies[i+1:] {
			// check cols
			a := galaxy.x
			b := other.x
			if a > b {
				a = other.x
				b = galaxy.x
			}
			for i := a; i < b; i++ {
				if emptyCols[i] {
					sum += emptySpaceMult
				} else {
					sum += 1
				}
			}

			// check rows
			a = galaxy.y
			b = other.y
			if a > b {
				a = other.y
				b = galaxy.y
			}
			for i := a; i < b; i++ {
				if emptyRows[i] {
					sum += emptySpaceMult
				} else {
					sum += 1
				}
			}
		}
	}

	return sum
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
	totalDistance := FindDistBetweenGalaxies(input, 1000000)
	return strconv.Itoa(totalDistance), nil
}
