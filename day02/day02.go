package day02

import (
	_ "embed"
	"slices"
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

type Game struct {
	Round int
	Sets  []CubeSet
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func Part01(input string) (string, error) {
	constraint := CubeSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0
	for _, line := range lines {
		game, err := ParseGameText(line)
		if err != nil {
			return "", err
		}

		isGameImpossible := slices.ContainsFunc(game.Sets, func(cs CubeSet) bool {
			return cs.Red > constraint.Red ||
				cs.Green > constraint.Green ||
				cs.Blue > constraint.Blue
		})

		if !isGameImpossible {
			sum += game.Round
		}
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	return "", nil
}

func ParseGameText(input string) (Game, error) {
	splitResult := strings.Split(input, ":")
	gameText := splitResult[0]
	roundText := splitResult[1]

	game := Game{}
	splitResult = strings.Split(gameText, " ")
	gameRound, err := strconv.Atoi(splitResult[1])
	if err != nil {
		return game, err
	}
	game.Round = gameRound

	setsText := strings.Split(roundText, ";")
	for _, setText := range setsText {
		drawTexts := strings.Split(setText, ",")

		cubeSet := CubeSet{}
		for _, drawText := range drawTexts {
			splitResult := strings.Split(strings.TrimSpace(drawText), " ")
			drawNumber, err := strconv.Atoi(splitResult[0])
			if err != nil {
				return game, err
			}
			drawColor := splitResult[1]
			switch drawColor {
			case "red":
				cubeSet.Red = drawNumber
			case "green":
				cubeSet.Green = drawNumber
			case "blue":
				cubeSet.Blue = drawNumber
			}
		}
		game.Sets = append(game.Sets, cubeSet)
	}

	return game, nil
}
