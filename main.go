package main

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"aoc2023/day03"
	"aoc2023/day04"
	"aoc2023/day05"
	"aoc2023/day06"
	"aoc2023/day07"
	"aoc2023/day08"
	"aoc2023/day09"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"aoc2023/day13"
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
	case 8:
		outputPart01, outputPart02, err = day08.Solve(day08.Input)
	case 9:
		outputPart01, outputPart02, err = day09.Solve(day09.Input)
	case 10:
		outputPart01, outputPart02, err = day10.Solve(day10.Input)
	case 11:
		outputPart01, outputPart02, err = day11.Solve(day11.Input)
	case 12:
		outputPart01, outputPart02, err = day12.Solve(day12.Input)
	case 13:
		outputPart01, outputPart02, err = day13.Solve(day13.Input)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Day %d part01 %s\n", dayToRun, outputPart01)
		fmt.Printf("Day %d part02 %s\n", dayToRun, outputPart02)
	}
}
