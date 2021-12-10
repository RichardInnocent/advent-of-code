package day9

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type heightmap struct {
	heights []int
	width   int
}

func Part1(filePath string) (string, error) {
	hmap, err := readMap(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Total of risk levels: %d", hmap.getTotalOfRiskLevels()), nil
}

func Part2(filePath string) (string, error) {
	hmap, err := readMap(filePath)
	if err != nil {
		return "", err
	}

	basins := hmap.getBasins()
	basinLengths := make([]int, len(basins))
	for index, basin := range basins {
		basinLengths[index] = len(basin)
	}

	sort.Ints(basinLengths)
	largestBasinSize := basinLengths[len(basinLengths)-1]
	secondLargestBasinSize := basinLengths[len(basinLengths)-2]
	thirdLargestBasinSize := basinLengths[len(basinLengths)-3]

	return fmt.Sprintf("Total size of largest three basins: %d", largestBasinSize*secondLargestBasinSize*thirdLargestBasinSize), nil
}

func (heightmap heightmap) getHeightAt(heightmapIndex int) (int, bool) {
	if heightmapIndex < 0 || heightmapIndex > len(heightmap.heights) {
		return 0, false
	}
	return heightmap.heights[heightmapIndex], true
}

func (heightmap heightmap) getHeightAbove(heightmapIndex int) (int, bool) {
	indexAbove, exists := heightmap.getHeightmapIndexAbove(heightmapIndex)
	if !exists {
		return 0, exists
	}
	return heightmap.getHeightAt(indexAbove)
}

func (heightmap heightmap) getHeightmapIndexAbove(heightmapIndex int) (indexAbove int, exists bool) {
	indexAbove = heightmapIndex - heightmap.width
	if indexAbove >= 0 && indexAbove < len(heightmap.heights) {
		exists = true
	}
	return
}

func (heightmap heightmap) getHeightBelow(heightmapIndex int) (int, bool) {
	indexBelow, exists := heightmap.getHeightmapIndexBelow(heightmapIndex)
	if !exists {
		return 0, exists
	}
	return heightmap.getHeightAt(indexBelow)
}

func (heightmap heightmap) getHeightmapIndexBelow(heightmapIndex int) (indexBelow int, exists bool) {
	indexBelow = heightmapIndex + heightmap.width
	if indexBelow >= 0 && indexBelow < len(heightmap.heights) {
		exists = true
	}
	return
}

func (heightmap heightmap) getHeightToLeft(heightmapIndex int) (int, bool) {
	indexToLeft, exists := heightmap.getHeightmapIndexToLeftOf(heightmapIndex)
	if !exists {
		return 0, exists
	}
	return heightmap.getHeightAt(indexToLeft)
}

func (heightmap heightmap) getHeightmapIndexToLeftOf(heightmapIndex int) (indexToLeft int, exists bool) {
	indexToLeft = heightmapIndex - 1
	if heightmapIndex%heightmap.width != 0 && indexToLeft >= 0 && indexToLeft < len(heightmap.heights) {
		exists = true
	}
	return
}

func (heightmap heightmap) getHeightToRight(heightmapIndex int) (int, bool) {
	indexToRight, exists := heightmap.getHeightmapIndexToRightOf(heightmapIndex)
	if !exists {
		return 0, exists
	}
	return heightmap.getHeightAt(indexToRight)
}

func (heightmap heightmap) getHeightmapIndexToRightOf(heightmapIndex int) (indexToRight int, exists bool) {
	indexToRight = heightmapIndex + 1
	if indexToRight%heightmap.width != 0 && indexToRight >= 0 && indexToRight < len(heightmap.heights) {
		exists = true
	}
	return
}

func (heightmap heightmap) isLowPoint(heightmapIndex int) bool {
	point, pointExists := heightmap.getHeightAt(heightmapIndex)
	if !pointExists {
		return false
	}

	pointAbove, pointAboveExists := heightmap.getHeightAbove(heightmapIndex)
	if pointAboveExists && pointAbove <= point {
		return false
	}

	pointBelow, pointBelowExists := heightmap.getHeightBelow(heightmapIndex)
	if pointBelowExists && pointBelow <= point {
		return false
	}

	pointToLeft, pointToLeftExists := heightmap.getHeightToLeft(heightmapIndex)
	if pointToLeftExists && pointToLeft <= point {
		return false
	}

	pointToRight, pointToRightExists := heightmap.getHeightToRight(heightmapIndex)
	if pointToRightExists && pointToRight <= point {
		return false
	}

	return true
}

func (heightmap heightmap) getBasins() (basins [][]int) {
	unassignedPointsWithinBasin := heightmap.getPointsWithinAnyBasin()

	for len(unassignedPointsWithinBasin) > 0 {
		basin, u := heightmap.getIndexOfPointsInBasinFromIndex(unassignedPointsWithinBasin[0], unassignedPointsWithinBasin)
		basins = append(basins, basin)
		unassignedPointsWithinBasin = u
	}
	return basins
}

func (heightmap heightmap) getIndexOfPointsInBasinFromIndex(heightmapIndex int, remainingPointsWithinAnyBasin []int) ([]int, []int) {
	var points []int
	points = append(points, heightmapIndex)
	remainingPointsWithinAnyBasin = removeFromSlice(remainingPointsWithinAnyBasin, heightmapIndex)

	p := heightmap.getIndexesOfPointsInBasinDirectlyAround(heightmapIndex, remainingPointsWithinAnyBasin)
	for _, point := range p {
		remainingPointsWithinAnyBasin = removeFromSlice(remainingPointsWithinAnyBasin, point)
		//points = append(points, point)
	}
	for _, point := range p {
		var p2 []int
		p2, remainingPointsWithinAnyBasin = heightmap.getIndexOfPointsInBasinFromIndex(point, remainingPointsWithinAnyBasin)
		points = append(points, p2...)
	}

	return points, remainingPointsWithinAnyBasin
}

func (heightmap heightmap) getIndexesOfPointsInBasinDirectlyAround(heightmapIndex int, pointsWithinAnyBasin []int) (points []int) {
	indexAbove, aboveExists := heightmap.getHeightmapIndexAbove(heightmapIndex)
	if aboveExists && sliceContains(pointsWithinAnyBasin, indexAbove) {
		points = append(points, indexAbove)
	}

	indexBelow, belowExists := heightmap.getHeightmapIndexBelow(heightmapIndex)
	if belowExists && sliceContains(pointsWithinAnyBasin, indexBelow) {
		points = append(points, indexBelow)
	}

	indexToLeft, toLeftExists := heightmap.getHeightmapIndexToLeftOf(heightmapIndex)
	if toLeftExists && sliceContains(pointsWithinAnyBasin, indexToLeft) {
		points = append(points, indexToLeft)
	}

	indexToRight, toRightExists := heightmap.getHeightmapIndexToRightOf(heightmapIndex)
	if toRightExists && sliceContains(pointsWithinAnyBasin, indexToRight) {
		points = append(points, indexToRight)
	}

	return
}

func (heightmap heightmap) getPointsWithinAnyBasin() (pointsWithinAnyBasin []int) {
	for index, height := range heightmap.heights {
		if height != 9 {
			pointsWithinAnyBasin = append(pointsWithinAnyBasin, index)
		}
	}
	return
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

func sliceContains(slice []int, value int) bool {
	_, found := sliceIndexOf(slice, value)
	return found
}

func sliceIndexOf(slice []int, value int) (indexOfElement int, found bool) {
	for index, element := range slice {
		if element == value {
			indexOfElement = index
			found = true
			return
		}
	}
	return
}

func removeFromSlice(slice []int, value int) []int {
	index, found := sliceIndexOf(slice, value)
	if found {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
