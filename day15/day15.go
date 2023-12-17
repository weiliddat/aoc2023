package day15

import (
	_ "embed"
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

func Part01(input string) (string, error) {
	instructions := strings.Split(strings.TrimSpace(input), ",")
	sum := 0
	for _, ins := range instructions {
		sum += hash(ins)
	}
	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	return "", nil
}

func hash(s string) int {
	c := 0
	for _, r := range s {
		c += int(r)
		c *= 17
		c %= 256
	}
	return c
}
