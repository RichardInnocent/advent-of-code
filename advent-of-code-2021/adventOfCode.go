package main

import (
	"advent-of-code-2021/day1"
	"advent-of-code-2021/day2"
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
	execute(2, 1, day2.Part2)
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
