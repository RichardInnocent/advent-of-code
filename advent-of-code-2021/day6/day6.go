package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	birthFrequencyDays = 7
	daysUntilMaturity  = 2
)

type shoal struct {
	groups *[]*shoalGroup
}

func newShoal() *shoal {
	var groups []*shoalGroup
	shoal := shoal{
		groups: &groups,
	}
	return &shoal
}

func (shoal *shoal) addSingleFish(daysUntilOffspringDue int) {
	shoal.addMultipleFish(daysUntilOffspringDue, 1)
}

func (shoal *shoal) addMultipleFish(daysUntilOffspringDue int, number int64) {
	for _, group := range *shoal.groups {
		if group.daysUntilOffspringDue == daysUntilOffspringDue {
			group.add(number)
			return
		}
	}
	newGroup := shoalGroup{
		daysUntilOffspringDue: daysUntilOffspringDue,
		number:                number,
	}
	newGroups := append(*shoal.groups, &newGroup)
	shoal.groups = &newGroups
}

func (shoal *shoal) advanceDay() {
	var newbornFish int64 = 0
	for _, group := range *shoal.groups {
		newbornFish += group.advanceDay()
	}
	shoal.addMultipleFish(daysUntilMaturity+birthFrequencyDays-1, newbornFish)
}

func (shoal *shoal) getNumberOfFish() int64 {
	var number int64 = 0
	for _, group := range *shoal.groups {
		number += group.number
	}
	return number
}

// A group is a collection of fish within the shoal that have the same number of days until their offspring is due. I
// originally stored each fish as its own entity but this consumes a needlessly large amount of RAM for lots of
// identical objects (and took a thousand years).
type shoalGroup struct {
	daysUntilOffspringDue int
	number                int64
}

func (group *shoalGroup) add(number int64) {
	group.number += number
}

func (group *shoalGroup) advanceDay() int64 {
	if group.daysUntilOffspringDue == 0 {
		group.daysUntilOffspringDue = birthFrequencyDays - 1
		return group.number
	}
	group.daysUntilOffspringDue -= 1
	return 0
}

func Part1() string {
	fish, err := getInitialFish()
	if err != nil {
		panic(err)
	}

	days := 80
	return fmt.Sprintf("Number of fish after %d days: %d", days, getNumberOfLanternfishAfterDays(fish, days))
}

func Part2() string {
	fish, err := getInitialFish()
	if err != nil {
		panic(err)
	}

	days := 256
	return fmt.Sprintf("Number of fish after %d days: %d", days, getNumberOfLanternfishAfterDays(fish, days))
}

func getNumberOfLanternfishAfterDays(shoal *shoal, days int) int64 {
	for i := 0; i < days; i++ {
		shoal.advanceDay()
	}
	return shoal.getNumberOfFish()
}

func getInitialFish() (*shoal, error) {
	file, err := os.Open("day6/lanternfish.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return createShoal(scanner.Text())
}

func createShoal(input string) (*shoal, error) {
	rawInputs := strings.Split(input, ",")
	shoal := newShoal()

	for _, input := range rawInputs {
		daysUntilOffspringDue, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		shoal.addSingleFish(daysUntilOffspringDue)
	}
	return shoal, nil
}
