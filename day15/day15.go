package day15

import (
	_ "embed"
	"errors"
	"regexp"
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

func Part01(input string) (string, error) {
	instructions := strings.Split(strings.TrimSpace(input), ",")
	sum := 0
	for _, ins := range instructions {
		sum += hash(ins)
	}
	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	type lens struct {
		label       string
		labelHash   int
		focalLength int
	}

	instructionTexts := strings.Split(strings.TrimSpace(input), ",")
	boxes := map[int][]lens{}

	matcher := regexp.MustCompile(`^(\w+)([\-=])(\d+)*`)
	for _, text := range instructionTexts {
		matched := matcher.FindStringSubmatch(text)
		if matched == nil || len(matched) != 4 {
			return "", errors.New("could not parse instruction: " + text)
		}
		label := matched[1]
		op := matched[2]
		hash := hash(label)
		if op == "=" {
			_, exists := boxes[hash]
			if !exists {
				boxes[hash] = []lens{}
			}

			focalLength, err := strconv.Atoi(matched[3])
			if err != nil {
				return "", err
			}
			l := lens{
				label:       label,
				labelHash:   hash,
				focalLength: focalLength,
			}

			boxIndex := slices.IndexFunc(boxes[hash], func(l lens) bool {
				return l.label == label
			})
			if boxIndex > -1 {
				boxes[hash][boxIndex] = l
			} else {
				boxes[hash] = append(boxes[hash], l)
			}
		} else {
			box, exists := boxes[hash]
			if exists {
				boxIndex := slices.IndexFunc(box, func(l lens) bool {
					return l.label == label
				})
				if boxIndex > -1 {
					boxes[hash] = slices.Delete(box, boxIndex, boxIndex+1)
				}
			}
		}
	}

	sum := 0
	for box, lenses := range boxes {
		for slot, lens := range lenses {
			sum = sum + (1+box)*(1+slot)*lens.focalLength
		}
	}

	return strconv.Itoa(sum), nil
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
