package day08

import (
	"aoc2023/aoc_util"
	_ "embed"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
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
	instructionText, nodeText, _ := strings.Cut(input, "\n\n")

	instructions := []int{}
	for i := 0; i < len(instructionText); i++ {
		if instructionText[i:i+1] == "L" {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}

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

	stepChan := make(chan int)
	foundChan := make(chan bool)

	for _, startNode := range startNodes {
		go nodeStepper(startNode, &nodes, stepChan, foundChan)
	}

	steps := 0
	found := atomic.Uint32{}

	for {
		instruction := instructions[steps%len(instructions)]
		steps++

		for range startNodes {
			stepChan <- instruction
		}

		for range startNodes {
			if <-foundChan {
				found.Add(1)
			}
		}

		if int(found.Load()) == len(startNodes) {
			break
		} else {
			found.Store(0)
		}
	}

	return strconv.Itoa(steps), nil
}

func nodeStepper(
	startNode *Node,
	nodes *map[string]Node,
	stepChan <-chan int,
	foundChan chan<- bool,
) {
	nextNode := *startNode

	for step := range stepChan {
		if step == 0 {
			nextNode = (*nodes)[nextNode.Left]
		} else {
			nextNode = (*nodes)[nextNode.Right]
		}

		if nextNode.Name[2:3] == "Z" {
			foundChan <- true
		} else {
			foundChan <- false
		}
	}
}
