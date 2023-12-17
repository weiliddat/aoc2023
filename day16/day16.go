package day16

import (
	"aoc2023/aoc_util"
	_ "embed"
	"fmt"
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
	lightmap := aoc_util.NewTileMap(input)

	fmt.Println(lightmap.Get(0, 1))

	return "", nil
}

func Part02(input string) (string, error) {
	return "", nil
}
