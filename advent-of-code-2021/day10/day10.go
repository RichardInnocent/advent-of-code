package day10

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

type enclosure struct {
	startCharacter, endCharacter int32
	completionScore              int
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

func (enclosures enclosureSet) getEnclosureForStartCharacter(startCharacter int32) (enclosure enclosure, found bool) {
	for _, e := range enclosures {
		if e.startCharacter == startCharacter {
			enclosure = e
			found = true
			return
		}
	}
	return
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

var enclosures = enclosureSet{
	{
		startCharacter:  '(',
		endCharacter:    ')',
		completionScore: 1,
		errorScore:      3,
	},
	{
		startCharacter:  '[',
		endCharacter:    ']',
		completionScore: 2,
		errorScore:      57,
	},
	{
		startCharacter:  '{',
		endCharacter:    '}',
		completionScore: 3,
		errorScore:      1_197,
	},
	{
		startCharacter:  '<',
		endCharacter:    '>',
		completionScore: 4,
		errorScore:      25_137,
	},
}

func Part1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open chunks file at %q. %w", filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

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

func Part2(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open chunks file at %q. %w", filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var completionScores []int

	for scanner.Scan() {
		completionScore, corrupted, err := getCompletionScore(scanner.Text(), enclosures)

		if err != nil {
			return "", err
		}

		if corrupted {
			continue
		}

		completionScores = append(completionScores, completionScore)
	}

	sort.Ints(completionScores)
	middleScore := completionScores[len(completionScores)/2]
	return fmt.Sprintf("Middle completion score: %d", middleScore), nil
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
			charInStack, charFound := charStack.pop()
			if !charFound || charInStack != enclosureType.startCharacter {
				score = enclosureType.errorScore
				return
			}
		}
	}
	return
}

func getCompletionScore(line string, enclosures enclosureSet) (int, bool, error) {
	enclosuresRequired, corrupted, err := getEnclosuresRequiredToComplete(line, enclosures)

	if err != nil {
		return 0, false, err
	}
	if corrupted {
		return 0, corrupted, nil
	}

	completionScore := 0
	for _, enclosureRequired := range enclosuresRequired {
		completionScore = (completionScore * 5) + enclosureRequired.completionScore
	}

	return completionScore, corrupted, err
}

func getEnclosuresRequiredToComplete(line string, enclosures enclosureSet) (enclosuresRequired []enclosure, corrupted bool, err error) {
	var charStack int32Stack
	for _, character := range line {
		if enclosures.isStartCharacter(character) {
			charStack.push(character)
		} else {
			enclosureType, found := enclosures.getEnclosureForEndCharacter(character)
			if !found {
				err = errors.New(fmt.Sprintf("invalid character '%c'", character))
			}
			charInStack, charFound := charStack.pop()
			if !charFound || charInStack != enclosureType.startCharacter {
				corrupted = true
				return
			}
		}
	}

	for {
		unmatchedStartChar, stackIsPopulated := charStack.pop()
		if !stackIsPopulated {
			break
		}
		requiredEnclosure, found := enclosures.getEnclosureForStartCharacter(unmatchedStartChar)
		if !found {
			err = errors.New(fmt.Sprintf("invalid character '%c'", unmatchedStartChar))
			return
		}

		enclosuresRequired = append(enclosuresRequired, requiredEnclosure)
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
