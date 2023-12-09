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
	counter := 0

	for {
		instruction := instructionText[counter%len(instructionText) : counter%len(instructionText)+1]
		counter++

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
	instructionText, nodeText, _ := strings.Cut(input, "\n\n")

	nodeMatcher := regexp.MustCompile(`(\w{3})`)
	nodes := map[string]Node{}
	lines := aoc_util.SplitLines(nodeText)
	for _, line := range lines {
		matches := nodeMatcher.FindAllString(line, -1)
		node := Node{matches[0], matches[1], matches[2]}
		nodes[node.Name] = node
	}

	startNodes := []*Node{}
	for _, node := range nodes {
		if node.Name[2:3] == "A" {
			startNodes = append(startNodes, &node)
		}
	}

	counters := []int{}
	for _, startNode := range startNodes {
		nextNode := nodes[startNode.Name]
		counter := 0

		for {
			instruction := instructionText[counter%len(instructionText) : counter%len(instructionText)+1]
			counter++

			if instruction == "L" {
				nextNode = nodes[nextNode.Left]
			} else {
				nextNode = nodes[nextNode.Right]
			}

			if nextNode.Name[2:3] == "Z" {
				counters = append(counters, counter)
				break
			}
		}
	}

	lcm := LCM(counters[0], counters[1], counters[2:]...)

	return strconv.Itoa(lcm), nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
