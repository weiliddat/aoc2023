package day13

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
	patterns := aoc_util.SplitBlocks(input)

	sum := 0
	for _, p := range patterns {
		// check vertical reflections
		rows := aoc_util.SplitLines(p)

		vIndices := []int{}
		for i := 1; i < len(rows[0]); i++ {
			vIndices = append(vIndices, i)
		}

		for _, row := range rows {
			vIndices = slices.DeleteFunc(vIndices, func(i int) bool {
				a := row[:i]
				b := row[i:]
				return !isReflected(a, b)
			})
		}

		// check horizontal reflections
		cols := aoc_util.IntoColumns(rows)
		hIndices := []int{}
		for i := 1; i < len(cols[0]); i++ {
			hIndices = append(hIndices, i)
		}

		for _, col := range cols {
			hIndices = slices.DeleteFunc(hIndices, func(i int) bool {
				a := col[:i]
				b := col[i:]
				return !isReflected(a, b)
			})
		}

		// add them up
		for _, v := range vIndices {
			sum += v
		}
		for _, h := range hIndices {
			sum += (h * 100)
		}
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	return "", nil
}

func isReflected(a, b string) bool {
	shortest := len(a)
	if len(b) < len(a) {
		shortest = len(b)
	}

	for i := 0; i < shortest; i++ {
		ab := a[len(a)-1-i]
		bb := b[i]
		if ab != bb {
			return false
		}
	}

	return true
}
