package day10

import (
	"aoc2023/aoc_util"
	_ "embed"
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

type TileMap struct {
	tiles  []string
	length int
}

func Part01(input string) (string, error) {
	return "", nil
}

func Part02(input string) (string, error) {
	return "", nil
}

func ParseToSlices(input *string) [][]rune {
	lines := aoc_util.SplitLines(*input)

	tilemap := [][]rune{}

	for y, line := range lines {
		if len(tilemap)-1 < y {
			row := []rune{}
			tilemap = append(tilemap, row)
		}

		for _, tile := range line {
			tilemap[y] = append(tilemap[y], tile)
		}
	}

	return tilemap
}

func NewTileMap(input *string) TileMap {
	length := strings.Index(*input, "\n")
	tilemap := TileMap{
		length: length,
		tiles:  make([]string, 0, len(*input)/(length+1)),
	}

	s := *input
	for {
		before, after, _ := strings.Cut(s, "\n")

		tilemap.tiles = append(tilemap.tiles, before)

		if len(after) < length {
			break
		} else {
			s = after
		}
	}

	return tilemap
}

func (t *TileMap) Get(x, y int) (string, bool) {

	return "", true
}
