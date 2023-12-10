package day10

import (
	_ "embed"
	"errors"
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

type TileMap struct {
	Tiles  []string
	length int
}

type Tile struct {
	S string
	X int
	Y int
}

type Path []Tile

func Part01(input string) (string, error) {
	tilemap := NewTileMap(input)

	start, found := tilemap.Find("S")
	if !found {
		return "", errors.New("could not find start")
	}

	paths := []*Path{}

	// check each direction for valid pipes
	tile, tileExists := tilemap.Get(start.X, start.Y-1)
	if tileExists {
		if tile.S == "|" || tile.S == "7" || tile.S == "F" {
			paths = append(paths, &Path{start, tile})
		}
	}
	tile, tileExists = tilemap.Get(start.X+1, start.Y)
	if tileExists {
		if tile.S == "-" || tile.S == "7" || tile.S == "J" {
			paths = append(paths, &Path{start, tile})
		}
	}
	tile, tileExists = tilemap.Get(start.X, start.Y+1)
	if tileExists {
		if tile.S == "|" || tile.S == "L" || tile.S == "J" {
			paths = append(paths, &Path{start, tile})
		}
	}
	tile, tileExists = tilemap.Get(start.X-1, start.Y)
	if tileExists {
		if tile.S == "-" || tile.S == "L" || tile.S == "F" {
			paths = append(paths, &Path{start, tile})
		}
	}

	// get next tile
	for {
		for _, path := range paths {
			next := path.FindNext(&tilemap)
			if next != nil {
				*path = append(*path, *next)
			} else {
				return "", errors.New("could not find next value")
			}
		}

		firstPath := *paths[0]
		secondPath := *paths[1]

		if firstPath[len(firstPath)-1] == secondPath[len(secondPath)-1] {
			break
		}
	}

	length := len(*paths[0]) - 1 // exclude start

	return strconv.Itoa(length), nil
}

func Part02(input string) (string, error) {
	return "", nil
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

func (p *Path) FindNext(t *TileMap) *Tile {
	curr := (*p)[len(*p)-1]
	prev := (*p)[len(*p)-2]

	// check north
	if curr.S == "|" || curr.S == "J" || curr.S == "L" {
		next, found := t.Get(curr.X, curr.Y-1)
		if found && !(prev.X == next.X && prev.Y == next.Y) {
			if next.S == "|" || next.S == "7" || next.S == "F" {
				return &next
			}
		}
	}

	// check south
	if curr.S == "|" || curr.S == "7" || curr.S == "F" {
		next, found := t.Get(curr.X, curr.Y+1)
		if found && !(prev.X == next.X && prev.Y == next.Y) {
			if next.S == "|" || next.S == "J" || next.S == "L" {
				return &next
			}
		}
	}

	// check east
	if curr.S == "-" || curr.S == "L" || curr.S == "F" {
		next, found := t.Get(curr.X+1, curr.Y)
		if found && !(prev.X == next.X && prev.Y == next.Y) {
			if next.S == "-" || next.S == "7" || next.S == "J" {
				return &next
			}
		}
	}

	// check west
	if curr.S == "-" || curr.S == "7" || curr.S == "J" {
		next, found := t.Get(curr.X-1, curr.Y)
		if found && !(prev.X == next.X && prev.Y == next.Y) {
			if next.S == "-" || next.S == "L" || next.S == "F" {
				return &next
			}
		}
	}

	return nil
}
