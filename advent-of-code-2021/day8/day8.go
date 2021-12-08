package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const delimiter = "|"
const (
	top         = 1
	topLeft     = 2
	topRight    = 3
	middle      = 4
	bottomLeft  = 5
	bottomRight = 6
	bottom      = 7
)

type position int

type displayNumber struct {
	signalPattern string
	decodedValue  *int
}

func newDisplayNumber(signalPattern string) displayNumber {
	return displayNumber{
		signalPattern: signalPattern,
	}
}

type display struct {
	inputs  []displayNumber
	outputs []displayNumber
}

func (display display) getSignalPatternsForDecodedValue(number int) (signalPattern string) {
	for _, value := range display.inputs {
		if value.decodedValue != nil && *value.decodedValue == number {
			return value.signalPattern
		}
	}
	return ""
}

func Part1() string {
	file, err := os.Open("day8/signal_patterns.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numberOfKnownValuesInOutput := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := parseEntry(scanner.Text())
		for _, output := range entry.outputs {
			if isKnownNumber(output.signalPattern) {
				numberOfKnownValuesInOutput++
			}
		}
	}
	return fmt.Sprintf("Number of 1, 4, 7 or 8s in the output: %d", numberOfKnownValuesInOutput)
}

func Part2() string {
	file, err := os.Open("day8/signal_patterns.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	sumOutputs := 0
	for scanner.Scan() {
		entry := parseEntry(scanner.Text())
		sumOutputs += entry.decipher()
	}
	return fmt.Sprintf("Sum of outputs: %d", sumOutputs)
}

func parseEntry(rawEntry string) (d display) {
	components := strings.Fields(rawEntry)

	delimiterMet := false

	for _, component := range components {
		if !delimiterMet {
			if component == delimiter {
				delimiterMet = true
			} else {
				d.inputs = append(d.inputs, newDisplayNumber(component))
			}
		} else {
			d.outputs = append(d.outputs, newDisplayNumber(component))
		}
	}

	return
}

func isKnownNumber(signalPattern string) bool {
	length := len(signalPattern)
	return length == 2 || length == 3 || length == 4 || length == 7
}

func (display display) decipher() (sumOutputs int) {

	display.getSignalPatternsForDecodedValue(7)

	for index, input := range display.inputs {
		number, numberDeduced := getNumber(input.signalPattern)
		if numberDeduced {
			input.decodedValue = &number
		}
		display.inputs[index] = input
	}

	positionMap := make(map[position]string)
	positionMap[top] = singleCharDifference(display.getSignalPatternsForDecodedValue(7), display.getSignalPatternsForDecodedValue(1))

	var allSignalPatterns []string
	for _, input := range display.inputs {
		allSignalPatterns = append(allSignalPatterns, input.signalPattern)
	}
	totalCharacterCounts := countCharactersInAll(allSignalPatterns)

	positionMap[bottomLeft], _ = getKeyFromMapWhereValueIs(totalCharacterCounts, 4)
	positionMap[bottomRight], _ = getKeyFromMapWhereValueIs(totalCharacterCounts, 9)
	positionMap[topLeft], _ = getKeyFromMapWhereValueIs(totalCharacterCounts, 6)
	positionMap[topRight] = singleCharDifference(display.getSignalPatternsForDecodedValue(1), positionMap[bottomRight])

	display.setInputNumberWhereMatches(2, func(n displayNumber) bool {
		return !strings.Contains(n.signalPattern, positionMap[bottomRight])
	})

	positionMap[middle] = singleCharDifference(
		display.getSignalPatternsForDecodedValue(4),
		positionMap[topLeft]+positionMap[topRight]+positionMap[bottomRight],
	)
	positionMap[bottom] = singleCharDifference(
		display.getSignalPatternsForDecodedValue(8),
		positionMap[top]+positionMap[topLeft]+positionMap[topRight]+positionMap[middle]+positionMap[bottomLeft]+positionMap[bottomRight],
	)

	// We have now deciphered all signals, so we can work out what the remaining numbers are

	display.setInputNumberWhereInputMatchesAll(
		0, positionMap[top]+positionMap[topLeft]+positionMap[topRight]+positionMap[bottomLeft]+positionMap[bottomRight]+positionMap[bottom],
	)
	display.setInputNumberWhereInputMatchesAll(
		3, positionMap[top]+positionMap[topRight]+positionMap[middle]+positionMap[bottomRight]+positionMap[bottom],
	)
	display.setInputNumberWhereInputMatchesAll(
		5, positionMap[top]+positionMap[topLeft]+positionMap[middle]+positionMap[bottomRight]+positionMap[bottom],
	)
	display.setInputNumberWhereInputMatchesAll(
		6, positionMap[top]+positionMap[topLeft]+positionMap[middle]+positionMap[bottomLeft]+positionMap[bottomRight]+positionMap[bottom],
	)
	display.setInputNumberWhereInputMatchesAll(
		9, positionMap[top]+positionMap[topLeft]+positionMap[topRight]+positionMap[middle]+positionMap[bottomRight]+positionMap[bottom],
	)

	// Don't need to do this, but why not decipher all of the inputs too?
	display.setOutputNumberWhereOutputMatchesAll(0, display.getSignalPatternsForDecodedValue(0))
	display.setOutputNumberWhereOutputMatchesAll(1, display.getSignalPatternsForDecodedValue(1))
	display.setOutputNumberWhereOutputMatchesAll(2, display.getSignalPatternsForDecodedValue(2))
	display.setOutputNumberWhereOutputMatchesAll(3, display.getSignalPatternsForDecodedValue(3))
	display.setOutputNumberWhereOutputMatchesAll(4, display.getSignalPatternsForDecodedValue(4))
	display.setOutputNumberWhereOutputMatchesAll(5, display.getSignalPatternsForDecodedValue(5))
	display.setOutputNumberWhereOutputMatchesAll(6, display.getSignalPatternsForDecodedValue(6))
	display.setOutputNumberWhereOutputMatchesAll(7, display.getSignalPatternsForDecodedValue(7))
	display.setOutputNumberWhereOutputMatchesAll(8, display.getSignalPatternsForDecodedValue(8))
	display.setOutputNumberWhereOutputMatchesAll(9, display.getSignalPatternsForDecodedValue(9))

	for i := 0; i < len(display.outputs); i++ {
		sumOutputs += *display.outputs[len(display.outputs)-1-i].decodedValue * int(math.Pow10(i))
	}
	return
}

func (display display) setInputNumberWhereMatches(value int, matcher func(displayNumber) bool) {
	for index, input := range display.inputs {
		if matcher(input) {
			input.decodedValue = &value
			display.inputs[index] = input
		}
	}
}

func (display display) setInputNumberWhereInputMatchesAll(value int, match string) {
	display.setInputNumberWhereMatches(value, func(n displayNumber) bool {
		return containsAllAndOnlyAll(n.signalPattern, match)
	})
}

func (display display) setOutputNumberWhereMatches(value int, matcher func(number displayNumber) bool) {
	for index, input := range display.outputs {
		if matcher(input) {
			input.decodedValue = &value
			display.outputs[index] = input
		}
	}
}

func (display display) setOutputNumberWhereOutputMatchesAll(value int, match string) {
	display.setOutputNumberWhereMatches(value, func(n displayNumber) bool {
		return containsAllAndOnlyAll(n.signalPattern, match)
	})
}

func (display display) getUndecipheredInputs() (results []string) {
	for _, value := range display.inputs {
		if value.decodedValue == nil {
			results = append(results, value.signalPattern)
		}
	}
	return
}

func singleCharDifference(str1, str2 string) string {
	charDifferences := charDifference(str1, str2)
	if len(charDifferences) > 0 {
		return charDifferences[0]
	}
	return ""
}

func charDifference(str1, str2 string) (charDifference []string) {
	for i := 0; i < len(str1); i++ {
		testCharacter := str1[i : i+1]
		if !strings.Contains(str2, testCharacter) {
			charDifference = append(charDifference, testCharacter)
		}
	}

	for i := 0; i < len(str2); i++ {
		testCharacter := str2[i : i+1]
		if !strings.Contains(str1, testCharacter) {
			charDifference = append(charDifference, testCharacter)
		}
	}

	return
}

func containsAllAndOnlyAll(test, characters string) bool {
	if len(test) != len(characters) {
		return false
	}
	for i := 0; i < len(characters); i++ {
		if !strings.Contains(test, characters[i:i+1]) {
			return false
		}
	}
	return true
}

func countCharactersInAll(strs []string) map[string]int {
	characterCount := make(map[string]int)
	for _, str := range strs {
		localCharacterCount := countCharacters(str)
		for character, count := range localCharacterCount {
			existingCount, found := characterCount[character]
			if !found {
				characterCount[character] = count
			} else {
				characterCount[character] = existingCount + count
			}
		}
	}
	return characterCount
}

func countCharacters(str string) map[string]int {
	characterCount := make(map[string]int)
	for i := 0; i < len(str); i++ {
		value := str[i : i+1]
		existingValue, found := characterCount[value]
		if !found {
			characterCount[value] = 1
		} else {
			characterCount[value] = existingValue + 1
		}
	}
	return characterCount
}

func getNumber(signalPattern string) (number int, numberDeduced bool) {
	numberDeduced = true
	switch len(signalPattern) {
	case 2:
		number = 1
	case 3:
		number = 7
	case 4:
		number = 4
	case 7:
		number = 8
	default:
		numberDeduced = false
	}
	return number, numberDeduced
}

func getKeyFromMapWhereValueIs(m map[string]int, value int) (string, bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	return "", false
}
