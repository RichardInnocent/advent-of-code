package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	forward = "forward"
	up      = "up"
	down    = "down"
)

type Position struct {
	horizontal int
	depth      int
}

type ReUnderstoodPosition struct {
	horizontal int
	depth      int
	aim        int
}

func Part1() string {
	commands := getCommands()
	position := executeCommandsAndDetermineFinalPosition(commands)
	return "depth x horizontal = " + strconv.Itoa(position.depth*position.horizontal)
}

func getCommands() *[]string {
	file, readError := os.Open("day2/commands.csv")
	if readError != nil {
		panic(readError)
	}
	defer file.Close()

	var commands []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	return &commands
}

func executeCommandsAndDetermineFinalPosition(commands *[]string) *Position {
	position := Position{horizontal: 0, depth: 0}

	for _, command := range *commands {
		updatePositionFromCommand(&position, command)
	}

	return &position
}

func updatePositionFromCommand(position *Position, command string) {
	components := strings.Split(command, " ")
	if len(components) != 2 {
		panic("Unexpected input: " + command)
	}
	direction := components[0]
	amount, e := strconv.Atoi(components[1])
	if e != nil {
		panic("Amount not numeric in command: " + command)
	}

	switch direction {
	case forward:
		position.horizontal += amount
	case down:
		position.depth += amount
	case up:
		position.depth -= amount
	default:
		panic("Unknown direction in command " + command)
	}
}

func Part2() string {
	commands := getCommands()
	position := executeReUnderstoodCommandsAndDetermineFinalPosition(commands)
	return "depth x horizontal = " + strconv.Itoa(position.depth*position.horizontal)
}

func executeReUnderstoodCommandsAndDetermineFinalPosition(commands *[]string) *ReUnderstoodPosition {
	position := newReUnderstoodPosition()

	for _, command := range *commands {
		updatePositionFromReUnderstoodCommand(position, command)
	}

	return position
}

func updatePositionFromReUnderstoodCommand(position *ReUnderstoodPosition, command string) {
	components := strings.Split(command, " ")
	if len(components) != 2 {
		panic("Unexpected input: " + command)
	}
	direction := components[0]
	amount, e := strconv.Atoi(components[1])
	if e != nil {
		panic("Amount not numeric in command: " + command)
	}

	switch direction {
	case forward:
		position.horizontal += amount
		position.depth += position.aim * amount
	case down:
		position.aim += amount
	case up:
		position.aim -= amount
	default:
		panic("Unknown direction in command " + command)
	}
}

func newReUnderstoodPosition() *ReUnderstoodPosition {
	position := ReUnderstoodPosition{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}
	return &position
}
