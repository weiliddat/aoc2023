package day16

import (
	"aoc2023/aoc_util"
	_ "embed"
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

	entryStep := Step{
		aoc_util.Tile{
			S: ".",
			X: -1,
			Y: 0,
		},
		east,
	}

	stepsTaken := []Step{}
	stepsToCheck := []Step{entryStep}
	nextSteps := []Step{}
	for {
		for _, step := range stepsToCheck {
			if slices.Contains(stepsTaken, step) {
				continue
			} else {
				stepsTaken = append(stepsTaken, step)
			}

			nextX := (step.direction % 2) * (2 - step.direction)
			nextY := ((step.direction + 1) % 2) * (step.direction - 1)
			next, ok := lightmap.Get(step.X+int(nextX), step.Y+int(nextY))
			if !ok { // out of bounds
				continue
			}

			nextStep := Step{next, step.direction}

			if next.S == "." {
				nextSteps = append(nextSteps, nextStep)
			} else if next.S == "/" {
				switch step.direction {
				case north:
					nextStep.direction = east
				case east:
					nextStep.direction = north
				case south:
					nextStep.direction = west
				case west:
					nextStep.direction = south
				}
				nextSteps = append(nextSteps, nextStep)
			} else if next.S == "\\" {
				switch step.direction {
				case north:
					nextStep.direction = west
				case west:
					nextStep.direction = north
				case south:
					nextStep.direction = east
				case east:
					nextStep.direction = south
				}
				nextSteps = append(nextSteps, nextStep)
			} else if next.S == "-" {
				switch step.direction {
				case west, east:
					nextSteps = append(nextSteps, nextStep)
				case north, south:
					nextStep.direction = west
					nextSteps = append(nextSteps, nextStep)
					nextSteps = append(nextSteps, Step{
						next,
						east,
					})
				}
			} else if next.S == "|" {
				switch step.direction {
				case north, south:
					nextSteps = append(nextSteps, nextStep)
				case west, east:
					nextStep.direction = north
					nextSteps = append(nextSteps, nextStep)
					nextSteps = append(nextSteps, Step{
						next,
						south,
					})
				}
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
		if step != entryStep && !slices.Contains(energized, step.Tile) {
			energized = append(energized, step.Tile)
		}
	}

	return strconv.Itoa(len(energized)), nil
}

func Part02(input string) (string, error) {
	return "", nil
}
