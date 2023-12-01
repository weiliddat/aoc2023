package main

import (
	"aoc2023/day01"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dayToRun, err := strconv.ParseInt(os.Args[1], 10, 8)

	if err != nil {
		fmt.Printf("Error converting %s to integer", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("Running day %d\n", dayToRun)

	outputPart01 := "Not implemented yet!"
	outputPart02 := "Not implemented yet!"

	switch dayToRun {
	case 1:
		outputPart01, outputPart02, err = day01.Solve(day01.Input)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Day %d part01 %s\n", dayToRun, outputPart01)
		fmt.Printf("Day %d part02 %s\n", dayToRun, outputPart02)
	}
}
