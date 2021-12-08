package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1(filePath string) (string, error) {
	measurements, err := getMeasurements(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Increases: %d", numberOfIncreases(measurements)), nil
}

func Part2(filePath string) (string, error) {
	measurements, err := getMeasurements(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Increases: %d", numberOfIncreasesInSlidingWindow(measurements)), nil
}

func getMeasurements(filePath string) (*[]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var measurements []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to convert measurement to integer. %w", err)
		}
		measurements = append(measurements, value)
	}

	return &measurements, nil
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
