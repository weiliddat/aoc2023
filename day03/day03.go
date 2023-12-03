package day03

import (
	_ "embed"
	"regexp"
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

type PartType int

const (
	Number PartType = iota
	Symbol
)

type Part struct {
	X      []int
	Y      int
	Type   PartType
	Number int
	Symbol string
}

func Part01(input string) (string, error) {
	lines := strings.Split(input, "\n")

	partMatcher := regexp.MustCompile(`[0-9]+|[^0-9\.]+`)

	schematic := map[int][]Part{}

	// parse schematic
	for ln, line := range lines {
		matches := partMatcher.FindAllStringIndex(line, -1)
		for _, match := range matches {
			part := Part{
				X: match,
				Y: ln,
			}

			matched := line[match[0]:match[1]]
			number, err := strconv.Atoi(matched)

			if err != nil {
				part.Type = Symbol
				part.Symbol = matched
			} else {
				part.Type = Number
				part.Number = number
			}

			schematic[ln] = append(schematic[ln], part)
		}
	}

	partNumbers := []int{}

	// check for parts that have adjacent symbols
	for ln, parts := range schematic {

		for _, part := range parts {
			// ignore symbols
			if part.Type == Symbol {
				continue
			}

			// accumulate surrounding parts
			surroundingParts := []Part{}
			previousRow := schematic[ln-1]
			if len(previousRow) > 0 {
				surroundingParts = append(surroundingParts, previousRow...)
			}
			currentRow := schematic[ln]
			if len(currentRow) > 0 {
				surroundingParts = append(surroundingParts, currentRow...)
			}
			nextRow := schematic[ln+1]
			if len(nextRow) > 0 {
				surroundingParts = append(surroundingParts, nextRow...)
			}

			// check for adjacent symbols
			for _, sp := range surroundingParts {
				// ignore same part
				if sp.X[0] == part.X[0] && sp.Y == part.Y {
					continue
				}

				// ignore numbers
				if sp.Type == Number {
					continue
				}

				// check if adjacent or within part Xs
				if sp.X[0]+1 >= part.X[0] &&
					sp.X[0] <= part.X[1] {
					partNumbers = append(partNumbers, part.Number)
					break
				}
			}
		}
	}

	sum := 0
	for _, pn := range partNumbers {
		sum += pn
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	lines := strings.Split(input, "\n")

	partMatcher := regexp.MustCompile(`[0-9]+|[^0-9\.]+`)

	schematic := map[int][]Part{}

	// parse schematic
	for ln, line := range lines {
		matches := partMatcher.FindAllStringIndex(line, -1)
		for _, match := range matches {
			part := Part{
				X: match,
				Y: ln,
			}

			matched := line[match[0]:match[1]]
			number, err := strconv.Atoi(matched)

			if err != nil {
				part.Type = Symbol
				part.Symbol = matched
			} else {
				part.Type = Number
				part.Number = number
			}

			schematic[ln] = append(schematic[ln], part)
		}
	}

	gearRatios := []int{}

	// check for gears (2 number parts next to a * part)
	for ln, parts := range schematic {

		for _, part := range parts {
			// ignore anything other than *
			if part.Symbol != "*" {
				continue
			}
			gearPart := part

			// accumulate surrounding parts
			surroundingParts := []Part{}
			previousRow := schematic[ln-1]
			if len(previousRow) > 0 {
				surroundingParts = append(surroundingParts, previousRow...)
			}
			currentRow := schematic[ln]
			if len(currentRow) > 0 {
				surroundingParts = append(surroundingParts, currentRow...)
			}
			nextRow := schematic[ln+1]
			if len(nextRow) > 0 {
				surroundingParts = append(surroundingParts, nextRow...)
			}

			adjacentNumbers := []int{}
			// check for adjacent numbers
			for _, sp := range surroundingParts {
				// ignore symbols
				if sp.Type == Symbol {
					continue
				}
				numberPart := sp

				// check if adjacent or within part Xs
				if gearPart.X[0]+1 >= numberPart.X[0] &&
					gearPart.X[0] <= numberPart.X[1] {
					adjacentNumbers = append(adjacentNumbers, numberPart.Number)
				}
			}

			if len(adjacentNumbers) == 2 {
				gearRatios = append(gearRatios, adjacentNumbers[0]*adjacentNumbers[1])
			}
		}
	}

	sum := 0
	for _, pn := range gearRatios {
		sum += pn
	}

	return strconv.Itoa(sum), nil
}
