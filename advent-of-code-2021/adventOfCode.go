package main

import (
	"advent-of-code-2021/day1"
	"advent-of-code-2021/day2"
	"advent-of-code-2021/day3"
	"advent-of-code-2021/day4"
	"advent-of-code-2021/day5"
	"fmt"
)

type fn func() string

type Task struct {
	day  int
	part int
}

type TaskResult struct {
	task   *Task
	result string
}

func main() {
	execute(1, 1, day1.Part1)
	execute(1, 2, day1.Part2)
	execute(2, 1, day2.Part1)
	execute(2, 2, day2.Part2)
	execute(3, 1, day3.Part1)
	execute(3, 2, day3.Part2)
	execute(4, 1, day4.Part1)
	execute(4, 2, day4.Part2)
	execute(5, 1, day5.Part1)
}

func execute(day, part int, runnable fn) {
	fmt.Printf("Day %d, part %d:\n", day, part)
	answer := runnable()
	result := toResult(day, part, answer)
	fmt.Printf("%s\n\n", result.result)
}

func toResult(day, part int, result string) *TaskResult {
	task := Task{
		day:  day,
		part: part,
	}
	return &TaskResult{
		task:   &task,
		result: result,
	}
}
