package day3

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func Part1(filePath string) (string, error) {
	rawDiagnosticEntries, fileError := getDiagnosticOutput(filePath)
	if fileError != nil {
		return "", fmt.Errorf("Failed to read diagnostic output. %w", fileError)
	}
	diagnosticEntries, err := toBoolArrays(rawDiagnosticEntries)
	if err != nil {
		return "", fmt.Errorf("failed to read diagnostic output. %w", err)
	}

	rawGammaRate := getRawGammaRate(diagnosticEntries)
	rawEpsilonRate := getRawEpsilonRate(rawGammaRate)

	gammaRate := boolsToInt(rawGammaRate)
	epsilonRate := boolsToInt(rawEpsilonRate)

	return fmt.Sprintf("Power consumption: %d", gammaRate*epsilonRate), nil
}

func Part2(filePath string) (string, error) {
	rawDiagnosticEntries, fileError := getDiagnosticOutput(filePath)
	if fileError != nil {
		return "", fmt.Errorf("Failed to read diagnostic output. %w", fileError)
	}
	diagnosticEntries, err := toBoolArrays(rawDiagnosticEntries)
	if err != nil {
		return "", fmt.Errorf("failed to read diagnostic output. %w", err)
	}

	rawOxygenGeneratorRating := getRawOxygenGeneratorRatings(diagnosticEntries)
	rawCO2ScrubberRating := getRawCO2ScrubberRating(diagnosticEntries)

	oxygenGeneratorRating := boolsToInt(rawOxygenGeneratorRating)
	co2ScrubberRating := boolsToInt(rawCO2ScrubberRating)

	return fmt.Sprintf("Life support rating: %d", oxygenGeneratorRating*co2ScrubberRating), nil
}

func getDiagnosticOutput(filePath string) (*[]string, error) {
	file, readError := os.Open(filePath)
	if readError != nil {
		return nil, fmt.Errorf("failed to read diagnostic file. %w", readError)
	}
	defer file.Close()

	var commands []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	return &commands, nil
}

func toBoolArrays(diagnosticEntries *[]string) (*[][]bool, error) {
	var boolArrays [][]bool

	for _, diagnosticEntry := range *diagnosticEntries {
		boolArray, err := toBoolArray(diagnosticEntry)
		if err != nil {
			return nil, err
		}
		boolArrays = append(boolArrays, *boolArray)
	}

	return &boolArrays, nil
}

func toBoolArray(diagnosticEntry string) (*[]bool, error) {
	var bools []bool
	for _, char := range diagnosticEntry {
		bit, err := toBool(char)
		if err != nil {
			return nil, err
		}

		bools = append(bools, bit)
	}
	return &bools, nil
}

func toBool(char int32) (result bool, err error) {
	if char == '0' {
		result = false
	} else if char == '1' {
		result = true
	} else {
		err = errors.New(fmt.Sprintf("Value %c is not a binary digit", char))
	}
	return
}

func getRawGammaRate(diagnosticEntries *[][]bool) *[]bool {
	var bits []bool
	for i := 0; i < len((*diagnosticEntries)[0]); i++ {
		mostCommonBit, oneBitIsMoreCommon := getMostCommonBit(diagnosticEntries, i)
		if !oneBitIsMoreCommon {
			panic("One bit was not more common")
		}
		bits = append(bits, mostCommonBit)
	}
	return &bits
}

func getMostCommonBit(diagnosticEntries *[][]bool, position int) (mostCommonBit, oneBitIsMoreCommon bool) {
	bitCountMap := make(map[bool]int)

	for _, value := range *diagnosticEntries {
		bit := value[position]
		count, found := bitCountMap[bit]
		if !found {
			bitCountMap[bit] = 1
		} else {
			bitCountMap[bit] = count + 1
		}
	}

	numberOfTimesCountReached := 0
	highestCount := 0
	for b, count := range bitCountMap {
		if count > highestCount {
			highestCount = count
			mostCommonBit = b
			numberOfTimesCountReached = 1
		} else if count == highestCount {
			numberOfTimesCountReached++
		}
	}

	oneBitIsMoreCommon = numberOfTimesCountReached < 2

	return
}

func getRawEpsilonRate(bits *[]bool) *[]bool {
	result := make([]bool, len(*bits))
	for i := 0; i < len(*bits); i++ {
		result[i] = !(*bits)[i]
	}
	return &result
}

func boolsToInt(bools *[]bool) int {
	result := 0
	for i := 0; i < len(*bools); i++ {
		if (*bools)[len(*bools)-i-1] {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}

func getRawOxygenGeneratorRatings(bools *[][]bool) *[]bool {
	result := *bools
	numberOfBits := len((*bools)[0])
	for i := 0; i < numberOfBits && len(result) > 1; i++ {
		mostCommonBit, oneBitIsMoreCommon := getMostCommonBit(&result, i)
		desiredBit := true
		if oneBitIsMoreCommon {
			desiredBit = mostCommonBit
		}
		result = filterBoolArrayWithBoolAtPosition(&result, desiredBit, i)
	}
	return &result[0]
}

func filterBoolArrayWithBoolAtPosition(bools *[][]bool, desiredValue bool, position int) (result [][]bool) {
	for _, boolArray := range *bools {
		if boolArray[position] == desiredValue {
			result = append(result, boolArray)
		}
	}
	return
}

func getRawCO2ScrubberRating(bools *[][]bool) *[]bool {
	result := *bools
	numberOfBits := len((*bools)[0])
	for i := 0; i < numberOfBits && len(result) > 1; i++ {
		mostCommonBit, oneBitIsMoreCommon := getMostCommonBit(&result, i)
		desiredBit := false
		if oneBitIsMoreCommon {
			desiredBit = !mostCommonBit
		}
		result = filterBoolArrayWithBoolAtPosition(&result, desiredBit, i)
	}
	return &result[0]
}
