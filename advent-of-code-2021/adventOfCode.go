package main

import (
	"advent-of-code-2021/day1"
	"advent-of-code-2021/day2"
	"advent-of-code-2021/day3"
	"advent-of-code-2021/day4"
	"advent-of-code-2021/day5"
	"advent-of-code-2021/day6"
	"advent-of-code-2021/day7"
	"advent-of-code-2021/day8"
	"advent-of-code-2021/day9"
	"fmt"
)

type fn func(inputFilePath string) (string, error)

func main() {
	execute(1, 1, day1.Part1, "day1/measurements.csv")
	execute(1, 2, day1.Part2, "day1/measurements.csv")
	execute(2, 1, day2.Part1, "day2/commands.csv")
	execute(2, 2, day2.Part2, "day2/commands.csv")
	execute(3, 1, day3.Part1, "day3/diagnostics.csv")
	execute(3, 2, day3.Part2, "day3/diagnostics.csv")
	execute(4, 1, day4.Part1, "day4/game.txt")
	execute(4, 2, day4.Part2, "day4/game.txt")
	execute(5, 1, day5.Part1, "day5/cent_coordinates.txt")
	execute(5, 2, day5.Part2, "day5/cent_coordinates.txt")
	execute(6, 1, day6.Part1, "day6/lanternfish.csv")
	execute(6, 2, day6.Part2, "day6/lanternfish.csv")
	execute(7, 1, day7.Part1, "day7/crab_positions.csv")
	execute(7, 2, day7.Part2, "day7/crab_positions.csv")
	execute(8, 1, day8.Part1, "day8/signal_patterns.txt")
	execute(8, 2, day8.Part2, "day8/signal_patterns.txt")
	execute(9, 1, day9.Part1, "day9/heightmap.txt")
}

func execute(day, part int, task fn, filePath string) {
	fmt.Printf("Day %d, part %d:\n", day, part)
	answer, err := task(filePath)
	if err != nil {
		fmt.Printf("Failed to execute task. Reason: %s", err.Error())
	}
	fmt.Printf("%s\n\n", answer)
}
