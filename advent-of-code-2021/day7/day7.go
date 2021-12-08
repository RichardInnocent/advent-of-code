package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(filePath string) (string, error) {
	positions, err := getCrabPositions(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read crab positions. %w", err)
	}

	min, _ := min(positions)
	max, _ := max(positions)
	bestPosition := 0
	var fuelRequired *int

	for i := min; i <= max; i++ {
		absoluteDifferences := sumAbsoluteDifferences(positions, i)
		if fuelRequired == nil || absoluteDifferences < *fuelRequired {
			fuelRequired = &absoluteDifferences
			bestPosition = i
		}
	}
	return fmt.Sprintf("Best position: %d. Fuel required: %d", bestPosition, *fuelRequired), nil
}

func Part2(filePath string) (string, error) {
	positions, err := getCrabPositions(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read crab positions. %w", err)
	}

	min, _ := min(positions)
	max, _ := max(positions)
	bestPosition := 0
	var fuelRequired *int

	for i := min; i <= max; i++ {
		absoluteDifferences := sumAbsoluteDifferencesWithReunderstoodFuelCosts(positions, i)
		if fuelRequired == nil || absoluteDifferences < *fuelRequired {
			fuelRequired = &absoluteDifferences
			bestPosition = i
		}
	}
	return fmt.Sprintf("Best position: %d. Fuel required: %d", bestPosition, *fuelRequired), nil
}

func min(values *[]int) (int, bool) {
	var min *int
	for _, value := range *values {
		if min == nil || value < *min {
			newMin := value
			min = &newMin
		}
	}
	if min == nil {
		return 0, false
	}
	return *min, true
}

func max(values *[]int) (int, bool) {
	var max *int
	for _, value := range *values {
		if max == nil || value > *max {
			newMin := value
			max = &newMin
		}
	}
	if max == nil {
		return 0, false
	}
	return *max, true
}

func sumAbsoluteDifferences(array *[]int, point int) int {
	difference := 0
	for _, value := range *array {
		difference += abs(value - point)
	}
	return difference
}

func sumAbsoluteDifferencesWithReunderstoodFuelCosts(array *[]int, point int) int {
	difference := 0
	for _, value := range *array {
		difference += pyramid(abs(value - point))
	}
	return difference
}

func pyramid(value int) int {
	result := 0
	for i := 1; i <= value; i++ {
		result += i
	}
	return result
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func getCrabPositions(filePath string) (*[]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open crab file %q. %w", filePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	return parsePositions(text)
}

func parsePositions(text string) (*[]int, error) {
	rawPositions := strings.Split(text, ",")
	return stringSliceToIntSlice(&rawPositions)
}

func stringSliceToIntSlice(stringSlice *[]string) (*[]int, error) {
	intSlice := make([]int, len(*stringSlice))
	for index, rawValue := range *stringSlice {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, err
		}
		intSlice[index] = value
	}
	return &intSlice, nil
}
