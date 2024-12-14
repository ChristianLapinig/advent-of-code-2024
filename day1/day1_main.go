package main

import (
	"fmt"
	"os"

	day1 "github.com/ChristianLapinig/advent-of-code-2024/day1/src"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no input passed. exiting...")
		os.Exit(1)
	}

	solution := day1.DayOne{FilePath: args[0]}

	fmt.Println("part 1:", solution.CalculateTotalDistance())
	fmt.Println("part 2:", solution.CalculateTotalSimilarityScore())
}
