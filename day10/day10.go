package day10

import (
	_ "embed"
	"errors"
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

type Tile struct {
	S *string
	X int
	Y int
}

type Path []Tile

func Part01(input string) (string, error) {
	tilemap := NewTileMap(&input)

	start := "S"
	startX, startY := tilemap.Find(start)
	if startX == -1 || startY == -1 {
		return "", errors.New("could not find start")
	}

	// paths := []Path{}

	// // check each direction for valid pipes
	// tile, tileExists := tilemap.Get(startX, startY-1)
	// if tileExists {
	// 	if tile == "|" || tile == "7" || tile == "F" {
	// 		paths = append(paths, Path{
	// 			Tile{&start, startX, startY},
	// 			Tile{&tile, startX, startY - 1},
	// 		})
	// 	}
	// }
	// tile, tileExists = tilemap.Get(startX+1, startY)
	// if tileExists {
	// 	if tile == "-" || tile == "7" || tile == "J" {
	// 		paths = append(paths, Path{
	// 			Tile{&start, startX, startY},
	// 			Tile{&tile, startX + 1, startY},
	// 		})
	// 	}
	// }
	// tile, tileExists = tilemap.Get(startX, startY+1)
	// if tileExists {
	// 	if tile == "|" || tile == "L" || tile == "J" {
	// 		paths = append(paths, Path{
	// 			Tile{&start, startX, startY},
	// 			Tile{&tile, startX, startY + 1},
	// 		})
	// 	}
	// }
	// tile, tileExists = tilemap.Get(startX-1, startY)
	// if tileExists {
	// 	if tile == "-" || tile == "L" || tile == "F" {
	// 		paths = append(paths, Path{
	// 			Tile{&start, startX, startY},
	// 			Tile{&tile, startX - 1, startY},
	// 		})
	// 	}
	// }

	// get next tile

	return "", nil
}

func Part02(input string) (string, error) {
	return "", nil
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

func (t *TileMap) Get(x, y int) (Tile, bool) {
	tile := Tile{}

	if y < 0 || y > len(t.tiles)-1 {
		return tile, false
	}

	row := t.tiles[y]

	if x < 0 || x > len(row)-1 {
		return tile, false
	}

	s := row[x : x+1]
	tile.S = &s
	tile.X = x
	tile.Y = y

	return tile, true
}

func (t *TileMap) Find(tile string) (int, int) {
	for y, row := range t.tiles {
		index := strings.Index(row, tile)
		if index > -1 {
			return index, y
		}
	}

	return -1, -1
}

func (p *Path) AppendNext(t *TileMap) {
	// tile := (*p)[len(*p)-1]
	// prev := (*p)[len(*p)-2]

	// if *tile.S == "|" {
	// 	north, exists := t.Get(tile.X, tile.Y-1)
	// 	if exists && prev.X != tile.X && prev.Y != tile.Y-1 {
	// 		if north == "|" || north == "7" || north == "F" {
	// 			*p = append(*p, Tile{&north, tile.X, tile.Y - 1})
	// 		}
	// 	}
	// 	south, exists := t.Get(tile.X, tile.Y+1)
	// 	if exists && prev.X != tile.X && prev.Y != tile.Y+1 {
	// 		if south == "|" || south == "J" || south == "L" {
	// 			*p = append(*p, Tile{&south, tile.X, tile.Y - 1})
	// 		}
	// 	}
	// }
}
