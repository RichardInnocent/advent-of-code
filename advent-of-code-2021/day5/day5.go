package day5

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	x, y int
}

type vent struct {
	start, end *coordinates
}

type gradient struct {
	xIncrease, yIncrease int
}

func newGradient(xDiff, yDiff int) (*gradient, error) {
	if xDiff != 0 && yDiff != 0 && mod(xDiff) != mod(yDiff) {
		return nil, errors.New("gradient is not perfectly vertical, horizontal or diagonal")
	}
	grad := gradient{
		xIncrease: sign(xDiff),
		yIncrease: sign(yDiff),
	}
	return &grad, nil
}

func newHorizontalOrVerticalGradient(xDiff, yDiff int) (*gradient, error) {
	if xDiff != 0 && yDiff != 0 {
		return nil, errors.New("gradient is not perfectly vertical, horizontal or diagonal")
	}
	grad := gradient{
		xIncrease: sign(xDiff),
		yIncrease: sign(yDiff),
	}
	return &grad, nil
}

func (grad *gradient) getNext(current *coordinates) *coordinates {
	nextCoordinates := coordinates{
		x: current.x + grad.xIncrease,
		y: current.y + grad.yIncrease,
	}
	return &nextCoordinates
}

func Part1() string {
	vents, err := getVents()
	if err != nil {
		panic(err)
	}

	coordinateCoverCount := make(map[coordinates]int)
	for _, vent := range *vents {
		coveredCoordinates, err := vent.getHorizontalAndVerticalCoveredCoordinates()
		if err == nil {
			for _, coveredSpot := range *coveredCoordinates {
				count, found := coordinateCoverCount[coveredSpot]
				if !found {
					coordinateCoverCount[coveredSpot] = 1
				} else {
					coordinateCoverCount[coveredSpot] = count + 1
				}
			}
		}
	}

	overlapCount := 0
	for _, value := range coordinateCoverCount {
		if value > 1 {
			overlapCount++
		}
	}

	return fmt.Sprintf("Overlapping points: %d", overlapCount)
}

func getVents() (*[]*vent, error) {
	file, fileErr := os.Open("day5/vent_coordinates.txt")
	if fileErr != nil {
		return nil, fileErr
	}

	var vents []*vent

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vent, ventErr := newVent(scanner.Text())
		if ventErr != nil {
			return nil, ventErr
		}
		vents = append(vents, vent)
	}

	return &vents, nil
}

func newVent(input string) (*vent, error) {
	rawCoordinates := strings.Split(input, " -> ")
	if len(rawCoordinates) != 2 {
		return nil, errors.New("Incorrect number of coordinates supplied in input " + input)
	}
	coordinatesSlice, err := stringSliceToCoordinateSlice(&rawCoordinates)
	if err != nil {
		return nil, err
	}

	vent := vent{
		start: (*coordinatesSlice)[0],
		end:   (*coordinatesSlice)[1],
	}

	return &vent, nil
}

func stringSliceToCoordinateSlice(stringSlice *[]string) (*[]*coordinates, error) {
	result := make([]*coordinates, len(*stringSlice))
	for index, str := range *stringSlice {
		coordinates, err := newCoordinate(str)
		if err != nil {
			return nil, err
		}
		result[index] = coordinates
	}
	return &result, nil
}

func newCoordinate(input string) (*coordinates, error) {
	rawValues := strings.Split(input, ",")
	if len(rawValues) != 2 {
		return nil, errors.New("Incorrect number of elements in coordinates for input " + input)
	}
	intSlice, err := stringSliceToIntSlice(&rawValues)
	if err != nil {
		return nil, err
	}

	coordinates := coordinates{
		x: (*intSlice)[0],
		y: (*intSlice)[1],
	}

	return &coordinates, nil
}

func stringSliceToIntSlice(stringSlice *[]string) (*[]int, error) {
	result := make([]int, len(*stringSlice))
	for index, str := range *stringSlice {
		value, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result[index] = value
	}
	return &result, nil
}

func (vent *vent) getHorizontalAndVerticalCoveredCoordinates() (*[]coordinates, error) {
	xDiff := vent.end.x - vent.start.x
	yDiff := vent.end.y - vent.start.y

	gradient, err := newHorizontalOrVerticalGradient(xDiff, yDiff)
	if err != nil {
		return nil, err
	}

	numberOfCoveredCoords := 1 + max(mod(xDiff), mod(yDiff))

	coveredCoordinates := make([]coordinates, numberOfCoveredCoords)
	coveredCoordinates[0] = *vent.start

	for i := 1; i < numberOfCoveredCoords; i++ {
		coveredCoordinates[i] = *gradient.getNext(&coveredCoordinates[i-1])
	}

	return &coveredCoordinates, nil
}

func max(values ...int) int {
	var currentMax *int
	for _, value := range values {
		if currentMax == nil || *currentMax < value {
			tmpValue := value
			currentMax = &tmpValue
		}
	}
	return *currentMax
}

func mod(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func sign(value int) int {
	if value < 0 {
		return -1
	}
	if value > 0 {
		return 1
	}
	return value
}
