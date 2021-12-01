package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1() string {
	measurements := getMeasurements()
	return fmt.Sprintf("Increases: %d", numberOfIncreases(measurements))
}

func Part2() string {
	measurements := getMeasurements()
	return fmt.Sprintf("Increases: %d", numberOfIncreasesInSlidingWindow(measurements))
}

func getMeasurements() *[]int {
	file, err := os.Open("day1/measurements.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var measurements []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		measurements = append(measurements, value)
	}

	return &measurements
}

func numberOfIncreases(measurements *[]int) int {
	increases := 0
	var previous *int = nil
	for i, measurement := range *measurements {
		if previous != nil && measurement > *previous {
			increases++
		}
		previous = &(*measurements)[i]
	}
	return increases
}

func numberOfIncreasesInSlidingWindow(measurements *[]int) int {
	increases := 0
	for i := 0; i < len(*measurements)-3; i++ {
		if (*measurements)[i] < (*measurements)[i+3] {
			increases++
		}
	}
	return increases
}