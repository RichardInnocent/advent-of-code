package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type boardElement struct {
	value  int
	marked bool
}

type board struct {
	size     int
	elements *[]*boardElement
}

func Part1(filePath string) (string, error) {
	inputs, boards, err := readInputFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read input file. %w", err)
	}

	for _, input := range *inputs {
		for _, board := range *boards {
			hasWon, winningScore := board.mark(input)
			if hasWon {
				return fmt.Sprintf("Winning score: %d", winningScore), nil
			}
		}
	}

	return "No board will win", nil
}

func Part2(filePath string) (string, error) {
	inputs, boards, err := readInputFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read input file. %w", err)
	}

	remainingBoards := *boards
	boardsNotWon := len(remainingBoards)

	for _, input := range *inputs {
		for index, board := range remainingBoards {
			if board == nil {
				continue
			}
			hasWon, winningScore := board.mark(input)
			if hasWon {
				if boardsNotWon == 1 {
					return fmt.Sprintf("Winning score: %d", winningScore), nil
				}
				remainingBoards[index] = nil
				boardsNotWon--
			}
		}
	}

	return "No board will win", nil
}

func readInputFile(filePath string) (*[]int, *[]*board, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open input file %q. %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	drawnNumbers, err := getDrawnNumbers(scanner)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read drawn numbers. %w", err)
	}

	boards, err := createBoards(scanner)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create game boards. %w", err)
	}

	return drawnNumbers, boards, nil
}

func getDrawnNumbers(scanner *bufio.Scanner) (*[]int, error) {
	scanner.Scan()
	input := scanner.Text()
	rawDrawnNumbers := strings.Split(input, ",")
	return stringArrayToIntArray(&rawDrawnNumbers)
}

func stringArrayToIntArray(stringArray *[]string) (*[]int, error) {
	result := make([]int, len(*stringArray))
	for index, value := range *stringArray {
		number, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value %q to integer. %w", value, err)
		}
		result[index] = number
	}
	return &result, nil
}

func createBoards(scanner *bufio.Scanner) (*[]*board, error) {
	var boards []*board
	for {
		board, err := createNextBoard(scanner)
		if err != nil {
			return nil, fmt.Errorf("failed to create board. %w", err)
		}
		if board == nil {
			return &boards, nil
		}
		boards = append(boards, board)
	}
}

func createNextBoard(scanner *bufio.Scanner) (*board, error) {
	boardRead := false
	var rawBoardValues []string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if boardRead {
				break
			}
		} else {
			boardRead = true
			rawRowValues := strings.Fields(text)
			rawBoardValues = append(rawBoardValues, rawRowValues...)
		}
	}

	if boardRead == false {
		return nil, nil
	}

	boardValues, err := stringArrayToIntArray(&rawBoardValues)
	if err != nil {
		return nil, fmt.Errorf("failed to convert values to board. %w", err)
	}
	board := newBoard(boardValues)
	return board, nil
}

func newBoard(values *[]int) *board {
	boardElements := toBoardElements(values)
	board := board{
		size:     int(math.Sqrt(float64(len(*values)))),
		elements: boardElements,
	}
	return &board
}

func toBoardElements(values *[]int) *[]*boardElement {
	boardElements := make([]*boardElement, len(*values))
	for index, value := range *values {
		boardElement := toBoardElement(value)
		boardElements[index] = &boardElement
	}
	return &boardElements
}

func toBoardElement(value int) boardElement {
	return boardElement{
		value:  value,
		marked: false,
	}
}

func (board *board) mark(value int) (won bool, winningScore int) {
	for index, boardValue := range *board.elements {
		if boardValue.value == value {
			(*(*board).elements)[index].marked = true
		}
	}
	won = board.hasWon()
	if won {
		winningScore = value * board.sumUnmarkedElements()
	}
	return
}

func (board *board) hasWon() bool {
	return board.hasRowFullyMarked() || board.hasColumnFullyMarked()
}

func (board *board) hasRowFullyMarked() bool {
rowLoop:
	for row := 0; row < board.size; row++ {
		for column := 0; column < board.size; column++ {
			element := board.getElementAt(row, column)
			if !element.marked {
				continue rowLoop
			}
		}
		return true
	}
	return false
}

func (board *board) hasColumnFullyMarked() bool {
columnLoop:
	for column := 0; column < board.size; column++ {
		for row := 0; row < board.size; row++ {
			element := board.getElementAt(row, column)
			if !element.marked {
				continue columnLoop
			}
		}
		return true
	}
	return false
}

func (board *board) sumUnmarkedElements() (sum int) {
	for _, element := range *(*board).elements {
		if !element.marked {
			sum += element.value
		}
	}
	return
}

func (board *board) getElementAt(row, column int) *boardElement {
	index := column + row*board.size
	return (*(*board).elements)[index]
}
