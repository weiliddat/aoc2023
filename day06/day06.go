package day06

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

func Part01(input string) (string, error) {
	races, err := ParseRaces01(&input)
	if err != nil {
		return "", nil
	}

	sum := 1

	for totalTime, distToBeat := range races {
		min, max := FindHoldRange(totalTime, distToBeat)
		sum = sum * (max - min + 1)
	}

	return strconv.Itoa(sum), nil
}

func Part02(input string) (string, error) {
	totalTime, distToBeat, err := ParseRaces02(&input)
	if err != nil {
		return "", nil
	}

	min, max := FindHoldRange(totalTime, distToBeat)
	sum := (max - min + 1)

	return strconv.Itoa(sum), nil
}

func ParseRaces01(input *string) (map[int]int, error) {
	races := map[int]int{}

	numberMatcher := regexp.MustCompile(`([\d]+)`)
	timeText, distanceText, _ := strings.Cut(*input, "\n")

	timeMatches := numberMatcher.FindAllStringSubmatch(timeText, -1)
	distanceMatches := numberMatcher.FindAllStringSubmatch(distanceText, -1)

	for index, timeMatch := range timeMatches {
		time, err := strconv.Atoi(timeMatch[0])
		if err != nil {
			return nil, err
		}
		distance, err := strconv.Atoi(distanceMatches[index][0])
		if err != nil {
			return nil, err
		}
		races[time] = distance
	}

	return races, nil
}

func ParseRaces02(input *string) (int, int, error) {
	numberMatcher := regexp.MustCompile(`([\d]+)`)
	timeText, distanceText, _ := strings.Cut(*input, "\n")

	timeMatches := numberMatcher.FindAllString(timeText, -1)
	distMatches := numberMatcher.FindAllString(distanceText, -1)

	time, err := strconv.Atoi(strings.Join(timeMatches, ""))
	if err != nil {
		return 0, 0, err
	}

	dist, err := strconv.Atoi(strings.Join(distMatches, ""))
	if err != nil {
		return 0, 0, err
	}

	return time, dist, nil
}

func FindHoldRange(totalTime int, distToBeat int) (int, int) {
	offset := 0

	for timeHeld := 0; timeHeld <= totalTime; timeHeld++ {
		dist := FindDistance(timeHeld, totalTime)

		if dist > distToBeat {
			offset = timeHeld
			break
		}
	}

	return offset, totalTime - offset
}

func FindDistance(timeHeld int, totalTime int) int {
	speed := timeHeld
	distance := (totalTime - timeHeld) * speed
	return distance
}
