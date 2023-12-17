package aoc_util

import "strings"

type TileMap struct {
	Tiles  []string
	length int
}

type Tile struct {
	S string
	X int
	Y int
}

func NewTileMap(input string) TileMap {
	length := strings.Index(input, "\n")
	tilemap := TileMap{
		length: length,
		Tiles:  make([]string, 0, len(input)/(length+1)),
	}

	s := input
	for {
		before, after, _ := strings.Cut(s, "\n")

		tilemap.Tiles = append(tilemap.Tiles, before)

		if len(after) < length {
			break
		} else {
			s = after
		}
	}

	return tilemap
}

func (tm *TileMap) Get(x, y int) (Tile, bool) {
	tile := Tile{}

	if y < 0 || y > len(tm.Tiles)-1 {
		return tile, false
	}

	row := tm.Tiles[y]

	if x < 0 || x > len(row)-1 {
		return tile, false
	}

	s := row[x : x+1]
	tile.S = s
	tile.X = x
	tile.Y = y

	return tile, true
}

func (tm *TileMap) Find(s string) (Tile, bool) {
	tile := Tile{}

	for y, row := range tm.Tiles {
		index := strings.Index(row, s)
		if index > -1 {
			tile.S = s
			tile.X = index
			tile.Y = y
			return tile, true
		}
	}

	return tile, false
}
