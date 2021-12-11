package day11

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	north = iota
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
)

type direction int

var allDirections = []direction{
	north,
	northEast,
	east,
	southEast,
	south,
	southWest,
	west,
	northWest,
}

type octopus struct {
	energyLevel int
	flashed     bool
}

func (octopus *octopus) incrementEnergyLevel() (flashed bool) {
	octopus.energyLevel++
	if octopus.energyLevel > 9 && !octopus.flashed {
		octopus.flashed = true
		flashed = true
	}
	return
}

func (octopus *octopus) resetIfFlashed() (wasReset bool) {
	if octopus.flashed {
		octopus.energyLevel = 0
		octopus.flashed = false
		wasReset = true
	}
	return
}

func newOctopus(energyLevel int) octopus {
	return octopus{
		energyLevel: energyLevel,
		flashed:     false,
	}
}

type octopusGrid struct {
	octopuses []*octopus
	width     int
}

func (grid octopusGrid) getLocationOfOctopusFromNeighbourInDirection(startLocation int, direction direction) (location int, exists bool) {
	exists = true
	switch direction {
	case north:
		location = startLocation - grid.width
	case northEast:
		location = startLocation - grid.width + 1
		exists = grid.positionToLeftExists(location)
	case east:
		location = startLocation + 1
		exists = grid.positionToLeftExists(location)
	case southEast:
		location = startLocation + grid.width + 1
		exists = grid.positionToLeftExists(location)
	case south:
		location = startLocation + grid.width
	case southWest:
		exists = grid.positionToLeftExists(startLocation)
		location = startLocation + grid.width - 1
	case west:
		exists = grid.positionToLeftExists(startLocation)
		location = startLocation - 1
	case northWest:
		exists = grid.positionToLeftExists(startLocation)
		location = startLocation - grid.width - 1
	default:
		exists = false
	}
	exists = exists && location >= 0 && location < len(grid.octopuses)
	return
}

func (grid octopusGrid) positionToLeftExists(startLocation int) bool {
	return startLocation%grid.width != 0
}

func (grid *octopusGrid) step() (flashes int) {
	for i := 0; i < len(grid.octopuses); i++ {
		grid.incrementEnergyAtLocation(i)
	}
	for _, octopus := range grid.octopuses {
		if octopus.resetIfFlashed() {
			flashes++
		}
	}
	return
}

func (grid *octopusGrid) incrementEnergyAtLocation(location int) {
	octopus := grid.octopuses[location]
	flashed := octopus.incrementEnergyLevel()
	if !flashed {
		return
	}

	for _, direction := range allDirections {
		neighbourLocation, neighbourExists := grid.getLocationOfOctopusFromNeighbourInDirection(location, direction)
		if neighbourExists {
			grid.incrementEnergyAtLocation(neighbourLocation)
		}
	}

	return
}

func Part1(filePath string) (string, error) {
	grid, err := readOctopuses(filePath)
	if err != nil {
		return "", fmt.Errorf("could not retrieve octopuses: %w", err)
	}

	numberOfFlashes := 0
	for i := 0; i < 100; i++ {
		numberOfFlashes += grid.step()
	}

	return fmt.Sprintf("Total number of flashes: %d", numberOfFlashes), nil
}

func Part2(filePath string) (string, error) {
	grid, err := readOctopuses(filePath)
	if err != nil {
		return "", fmt.Errorf("could not retrieve octopuses: %w", err)
	}

	numberOfSteps := 1
	for {
		if grid.step() == len(grid.octopuses) {
			break
		}
		numberOfSteps++
	}

	return fmt.Sprintf("Total number of steps required to synchronise: %d", numberOfSteps), nil
}

func readOctopuses(filePath string) (octopusGrid, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return octopusGrid{}, fmt.Errorf("could not open octopus grid file at %q: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	width := -1
	var octopuses []*octopus

	for scanner.Scan() {
		line := scanner.Text()
		if width == -1 {
			width = len(line)
		} else if width != len(line) {
			return octopusGrid{}, errors.New("octopuses do not form a grid")
		}

		for _, character := range line {
			energyLevel, err := convertToInt(character)
			if err != nil {
				return octopusGrid{}, err
			}
			octopus := newOctopus(energyLevel)
			octopuses = append(octopuses, &octopus)
		}
	}

	return octopusGrid{octopuses: octopuses, width: width}, nil
}

func convertToInt(character int32) (int, error) {
	value := character - 48
	if value < 0 || value > 9 {
		return 0, fmt.Errorf("value %c is not numeric", character)
	}
	return int(value), nil
}
