package main

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"aoc2023/day03"
	"aoc2023/day04"
	"aoc2023/day05"
	"aoc2023/day06"
	"aoc2023/day07"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dayToRun, err := strconv.ParseInt(os.Args[1], 10, 8)

	if err != nil {
		fmt.Printf("Error converting %s to integer\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("Running day %d\n", dayToRun)

	outputPart01 := "Not implemented yet!"
	outputPart02 := "Not implemented yet!"

	switch dayToRun {
	case 1:
		outputPart01, outputPart02, err = day01.Solve(day01.Input)
	case 2:
		outputPart01, outputPart02, err = day02.Solve(day02.Input)
	case 3:
		outputPart01, outputPart02, err = day03.Solve(day03.Input)
	case 4:
		outputPart01, outputPart02, err = day04.Solve(day04.Input)
	case 5:
		outputPart01, outputPart02, err = day05.Solve(day05.Input)
	case 6:
		outputPart01, outputPart02, err = day06.Solve(day06.Input)
	case 7:
		outputPart01, outputPart02, err = day07.Solve(day07.Input)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Day %d part01 %s\n", dayToRun, outputPart01)
		fmt.Printf("Day %d part02 %s\n", dayToRun, outputPart02)
	}
}
