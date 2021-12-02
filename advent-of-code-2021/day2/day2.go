package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontal int
	depth      int
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
	case "forward":
		position.horizontal += amount
	case "down":
		position.depth += amount
	case "up":
		position.depth -= amount
	default:
		panic("Unknown direction in command " + command)
	}
}
