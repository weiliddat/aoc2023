package day16

import (
	"aoc2023/aoc_util"
	_ "embed"
	"errors"
	"fmt"
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

type Direction int

const (
	north Direction = iota
	east
	south
	west
	none
)

type Path struct {
	tiles     []aoc_util.Tile
	direction Direction
}

func Part01(input string) (string, error) {
	lightmap := aoc_util.NewTileMap(input)

	entryTile, ok := lightmap.Get(0, 0)
	if !ok {
		return "", errors.New("cannot find starting tile")
	}
	entryPath := Path{
		tiles:     []aoc_util.Tile{entryTile},
		direction: east,
	}
	paths := []Path{entryPath}

	for {
		remaining := false

		// time.Sleep(time.Second * 2)
		fmt.Println()
		for i := range paths {
			fmt.Println("path", i, paths[i])
			if paths[i].direction == none {
				continue
			}

			last := paths[i].tiles[len(paths[i].tiles)-1]
			nextX := (paths[i].direction % 2) * (2 - paths[i].direction)
			nextY := ((paths[i].direction + 1) % 2) * (paths[i].direction - 1)
			next, ok := lightmap.Get(last.X+int(nextX), last.Y+int(nextY))
			// fmt.Println("path", i, "next", next)

			// next tile is out of map
			if !ok {
				paths[i].direction = none
				continue
			}

			paths[i].tiles = append(paths[i].tiles, next)
			remaining = true

			if next.S == "." {
				// continue
			} else if next.S == "/" {
				switch paths[i].direction {
				case north:
					paths[i].direction = east
				case east:
					paths[i].direction = north
				case south:
					paths[i].direction = west
				case west:
					paths[i].direction = south
				}
			} else if next.S == "\\" {
				switch paths[i].direction {
				case north:
					paths[i].direction = west
				case west:
					paths[i].direction = north
				case south:
					paths[i].direction = east
				case east:
					paths[i].direction = south
				}
			} else if next.S == "-" {
				switch paths[i].direction {
				case north, south:
					paths[i].direction = west
					paths = append(paths, Path{
						tiles:     slices.Clone(paths[i].tiles),
						direction: east,
					})
				}
			} else if next.S == "|" {
				switch paths[i].direction {
				case west, east:
					paths[i].direction = north
					paths = append(paths, Path{
						tiles:     slices.Clone(paths[i].tiles),
						direction: south,
					})
				}
			}
		}

		if !remaining {
			break
		}
	}

	energizedTiles := []aoc_util.Tile{}
	for i, path := range paths {
		fmt.Println("path", i, path.tiles)
		energizedTiles = append(energizedTiles, path.tiles...)
	}
	fmt.Println((energizedTiles))
	compacted := slices.CompactFunc(energizedTiles, func(t1, t2 aoc_util.Tile) bool {
		return t1.X == t2.X && t1.Y == t2.Y
	})

	return strconv.Itoa(len(compacted)), nil
}

func Part02(input string) (string, error) {
	return "", nil
}
