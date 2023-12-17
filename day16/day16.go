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
)

type Step struct {
	aoc_util.Tile
	direction Direction
}

func (d Direction) String() string {
	switch d {
	case north:
		return "^"
	case south:
		return "v"
	case west:
		return "<"
	case east:
		return ">"
	default:
		return "?"
	}
}

func (s Step) String() string {
	return fmt.Sprintf("%d %d %s", s.X, s.Y, s.direction)
}

func Part01(input string) (string, error) {
	lightmap := aoc_util.NewTileMap(input)

	entryTile, ok := lightmap.Get(0, 0)
	if !ok {
		return "", errors.New("cannot find entry tile")
	}
	entryStep := Step{entryTile, east}
	energized := findEnergized(lightmap, entryStep)

	return strconv.Itoa(energized), nil
}

func findEnergized(lightmap aoc_util.TileMap, entryStep Step) int {
	stepsTaken := []Step{}
	stepsToCheck := []Step{entryStep}
	nextSteps := []Step{}
	for {
		for _, step := range stepsToCheck {
			stepsTaken = append(stepsTaken, step)

			directions := []Direction{}

			if step.S == "." {
				directions = append(directions, step.direction)
			} else if step.S == "/" {
				switch step.direction {
				case north:
					directions = append(directions, east)
				case east:
					directions = append(directions, north)
				case south:
					directions = append(directions, west)
				case west:
					directions = append(directions, south)
				}
			} else if step.S == "\\" {
				switch step.direction {
				case north:
					directions = append(directions, west)
				case west:
					directions = append(directions, north)
				case south:
					directions = append(directions, east)
				case east:
					directions = append(directions, south)
				}
			} else if step.S == "-" {
				switch step.direction {
				case west, east:
					directions = append(directions, step.direction)
				case north, south:
					directions = append(directions, west)
					directions = append(directions, east)
				}
			} else if step.S == "|" {
				switch step.direction {
				case north, south:
					directions = append(directions, step.direction)
				case west, east:
					directions = append(directions, north)
					directions = append(directions, south)
				}
			}

			for _, d := range directions {
				nextX := (d % 2) * (2 - d)
				nextY := ((d + 1) % 2) * (d - 1)
				next, ok := lightmap.Get(step.X+int(nextX), step.Y+int(nextY))
				if !ok {
					continue
				}
				nextStep := Step{next, d}
				if slices.Contains(stepsTaken, nextStep) {
					continue
				}
				nextSteps = append(nextSteps, nextStep)
			}
		}

		if len(nextSteps) == 0 {
			break
		}

		stepsToCheck = nextSteps
		nextSteps = []Step{}
	}

	energized := []aoc_util.Tile{}

	for _, step := range stepsTaken {
		if !slices.Contains(energized, step.Tile) {
			energized = append(energized, step.Tile)
		}
	}

	return len(energized)
}

func Part02(input string) (string, error) {
	lightmap := aoc_util.NewTileMap(input)
	width := len(lightmap.Tiles[0])
	height := len(lightmap.Tiles)

	highestEnergized := 0

	// north
	for x := 0; x < width; x++ {
		entryTile, ok := lightmap.Get(x, 0)
		if !ok {
			return "", errors.New("cannot find entry tile")
		}
		entryStep := Step{entryTile, south}
		energized := findEnergized(lightmap, entryStep)
		if energized > highestEnergized {
			highestEnergized = energized
		}
	}

	// east
	for y := 0; y < height; y++ {
		entryTile, ok := lightmap.Get(width-1, y)
		if !ok {
			return "", errors.New("cannot find entry tile")
		}
		entryStep := Step{entryTile, west}
		energized := findEnergized(lightmap, entryStep)
		if energized > highestEnergized {
			highestEnergized = energized
		}
	}

	// south
	for x := 0; x < width; x++ {
		entryTile, ok := lightmap.Get(x, height-1)
		if !ok {
			return "", errors.New("cannot find entry tile")
		}
		entryStep := Step{entryTile, north}
		energized := findEnergized(lightmap, entryStep)
		if energized > highestEnergized {
			highestEnergized = energized
		}
	}

	// west
	for y := 0; y < height; y++ {
		entryTile, ok := lightmap.Get(0, y)
		if !ok {
			return "", errors.New("cannot find entry tile")
		}
		entryStep := Step{entryTile, east}
		energized := findEnergized(lightmap, entryStep)
		if energized > highestEnergized {
			highestEnergized = energized
		}
	}

	return strconv.Itoa(highestEnergized), nil
}
