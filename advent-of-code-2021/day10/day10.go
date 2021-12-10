package day10

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type enclosure struct {
	startCharacter, endCharacter int32
	errorScore                   int
}

type enclosureSet []enclosure

func (enclosures enclosureSet) isStartCharacter(startCharacter int32) bool {
	for _, e := range enclosures {
		if e.startCharacter == startCharacter {
			return true
		}
	}
	return false
}

func (enclosures enclosureSet) getEnclosureForEndCharacter(endCharacter int32) (enclosure enclosure, found bool) {
	for _, e := range enclosures {
		if e.endCharacter == endCharacter {
			enclosure = e
			found = true
			return
		}
	}
	return
}

func Part1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open chunks file at %q. %w", filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	enclosures := enclosureSet{
		{
			startCharacter: '(',
			endCharacter:   ')',
			errorScore:     3,
		},
		{
			startCharacter: '[',
			endCharacter:   ']',
			errorScore:     57,
		},
		{
			startCharacter: '{',
			endCharacter:   '}',
			errorScore:     1_197,
		},
		{
			startCharacter: '<',
			endCharacter:   '>',
			errorScore:     25_137,
		},
	}

	totalErrorScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		errorScore, err := getSyntaxErrorScore(line, enclosures)

		if err != nil {
			return "", err
		}

		totalErrorScore += errorScore
	}
	return fmt.Sprintf("Total error score: %d", totalErrorScore), nil
}

func getSyntaxErrorScore(line string, enclosures enclosureSet) (score int, err error) {
	var charStack int32Stack
	for _, character := range line {
		if enclosures.isStartCharacter(character) {
			charStack.push(character)
		} else {
			enclosureType, found := enclosures.getEnclosureForEndCharacter(character)
			if !found {
				err = errors.New(fmt.Sprintf("invalid character '%c'", character))
			}
			charInStack, charFound := charStack.peek()
			if !charFound || charInStack != enclosureType.startCharacter {
				score = enclosureType.errorScore
				return
			}
			charStack.pop()
		}
	}
	return
}

type int32Stack []int32

func (s *int32Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *int32Stack) push(value int32) {
	*s = append(*s, value)
}

func (s *int32Stack) peek() (value int32, found bool) {
	if s.isEmpty() {
		return
	}
	found = true
	value = (*s)[len(*s)-1]
	return
}

func (s *int32Stack) pop() (value int32, found bool) {
	if s.isEmpty() {
		return
	}
	found = true
	index := len(*s) - 1
	value = (*s)[index]
	*s = (*s)[:index]
	return
}
