package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lanternFish struct {
	daysUntilOffspringDue int
}

func (f *lanternFish) advanceDay() *lanternFish {
	if f.daysUntilOffspringDue == 0 {
		f.daysUntilOffspringDue = 6
		offspring := lanternFish{
			daysUntilOffspringDue: 8,
		}
		return &offspring
	}
	f.daysUntilOffspringDue--
	return nil
}

func Part1() string {
	fish, err := getInitialFish()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 80; i++ {
		var offspring []*lanternFish
		for _, f := range *fish {
			fOffspring := f.advanceDay()
			if fOffspring != nil {
				offspring = append(offspring, fOffspring)
			}
		}
		newShoal := append(*fish, offspring...)
		fish = &newShoal
	}

	return fmt.Sprintf("Number of fish after 80 days: %d", len(*fish))
}

func getInitialFish() (*[]*lanternFish, error) {
	file, err := os.Open("day6/lantern_fish.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return createAllLanternFish(scanner.Text())
}

func createAllLanternFish(input string) (*[]*lanternFish, error) {
	rawInputs := strings.Split(input, ",")
	fish := make([]*lanternFish, len(rawInputs))
	for index, input := range rawInputs {
		f, err := createSingleLanternFish(input)
		if err != nil {
			return nil, err
		}
		fish[index] = f
	}
	return &fish, nil
}

func createSingleLanternFish(input string) (*lanternFish, error) {
	daysUntilOffspringDue, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	fish := lanternFish{
		daysUntilOffspringDue: daysUntilOffspringDue,
	}
	return &fish, nil
}
