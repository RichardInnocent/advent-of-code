package day9

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type heightmap struct {
	heights []int
	width   int
}

func (heightmap heightmap) getHeightAt(heightmapIndex int) (int, bool) {
	if heightmapIndex < 0 || heightmapIndex > len(heightmap.heights) {
		return 0, false
	}
	return heightmap.heights[heightmapIndex], true
}

func (heightmap heightmap) getPointAbove(heightmapIndex int) (int, bool) {
	indexAbove := heightmapIndex - heightmap.width
	if indexAbove < 0 {
		return 0, false
	}
	return heightmap.getHeightAt(indexAbove)
}

func (heightmap heightmap) getPointBelow(heightmapIndex int) (int, bool) {
	indexBelow := heightmapIndex + heightmap.width
	if indexBelow > len(heightmap.heights) {
		return 0, false
	}
	return heightmap.getHeightAt(indexBelow)
}

func (heightmap heightmap) getPointToLeft(heightmapIndex int) (int, bool) {
	if heightmapIndex%heightmap.width == 0 {
		return 0, false
	}
	return heightmap.getHeightAt(heightmapIndex - 1)
}

func (heightmap heightmap) getPointToRight(heightmapIndex int) (int, bool) {
	indexToRight := heightmapIndex + 1
	if indexToRight%heightmap.width == 0 {
		return 0, false
	}
	return heightmap.getHeightAt(indexToRight)
}

func (heightmap heightmap) isLowPoint(heightmapIndex int) bool {
	point, pointExists := heightmap.getHeightAt(heightmapIndex)
	if !pointExists {
		return false
	}

	pointAbove, pointAboveExists := heightmap.getPointAbove(heightmapIndex)
	if pointAboveExists && pointAbove <= point {
		return false
	}

	pointBelow, pointBelowExists := heightmap.getPointBelow(heightmapIndex)
	if pointBelowExists && pointBelow <= point {
		return false
	}

	pointToLeft, pointToLeftExists := heightmap.getPointToLeft(heightmapIndex)
	if pointToLeftExists && pointToLeft <= point {
		return false
	}

	pointToRight, pointToRightExists := heightmap.getPointToRight(heightmapIndex)
	if pointToRightExists && pointToRight <= point {
		return false
	}

	return true
}

func (heightmap heightmap) getRiskLevels() (riskLevels []int) {
	for index, height := range heightmap.heights {
		if heightmap.isLowPoint(index) {
			riskLevels = append(riskLevels, 1+height)
		}
	}
	return
}

func (heightmap heightmap) getTotalOfRiskLevels() (total int) {
	riskLevels := heightmap.getRiskLevels()
	for _, riskLevel := range riskLevels {
		total += riskLevel
	}
	return
}

func Part1(filePath string) (string, error) {
	hmap, err := readMap(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Total of risk levels: %d", hmap.getTotalOfRiskLevels()), nil
}

func readMap(filePath string) (hmap heightmap, err error) {
	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		err = fmt.Errorf("could not open file containing hMap at %q. %w", filePath, fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var heightmapWidth *int
	var heightmapValues []int
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		if heightmapWidth == nil {
			heightmapWidth = &lineLength
		} else if *heightmapWidth != lineLength {
			err = errors.New("not all input lines are of the same length")
			return
		}
		lineAsIntArray, convErr := stringToIntArray(line)
		if convErr != nil {
			err = fmt.Errorf("could not convert input line %q to int array. %w", line, convErr)
			return
		}
		heightmapValues = append(heightmapValues, lineAsIntArray...)
	}

	hmap = heightmap{
		heights: heightmapValues,
		width:   *heightmapWidth,
	}

	return
}

func stringToIntArray(value string) ([]int, error) {
	intArray := make([]int, len(value))
	for i := 0; i < len(value); i++ {
		character := value[i : i+1]
		intValue, convErr := strconv.Atoi(character)
		if convErr != nil {
			return []int{}, fmt.Errorf("could not convert character %q to integer. %w", character, convErr)
		}
		intArray[i] = intValue
	}
	return intArray, nil
}
