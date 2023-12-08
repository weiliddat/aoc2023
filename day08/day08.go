package day08

import (
	"aoc2023/aoc_util"
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

type Node struct {
	Name  string
	Left  string
	Right string
}

func Part01(input string) (string, error) {
	startNodeName := "AAA"
	endNodeName := "ZZZ"

	instructionText, nodeText, _ := strings.Cut(input, "\n\n")

	nodeMatcher := regexp.MustCompile(`(\w{3})`)
	nodes := map[string]Node{}
	lines := aoc_util.SplitLines(nodeText)
	for _, line := range lines {
		matches := nodeMatcher.FindAllString(line, -1)
		node := Node{matches[0], matches[1], matches[2]}
		nodes[node.Name] = node
	}

	nextNode := nodes[startNodeName]
	instructionIndex := 0
	counter := 0
	for {
		instruction := instructionText[instructionIndex : instructionIndex+1]

		instructionIndex++
		counter++

		if instructionIndex >= len(instructionText) {
			instructionIndex = 0
		}

		if instruction == "L" {
			nextNode = nodes[nextNode.Left]
		} else {
			nextNode = nodes[nextNode.Right]
		}

		if nextNode.Name == endNodeName {
			break
		}
	}

	return strconv.Itoa(counter), nil
}

func Part02(input string) (string, error) {
	return "", nil
}
