package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Reicher/AoC25/helper"
)

type FreshList struct {
	ranges []FreshRange
}

func (fl *FreshList) newRange(rangeStr string) error {
	from, err := strconv.Atoi(strings.Split(rangeStr, "-")[0])
	if err != nil {
		fmt.Println("Error converting from value:", err)
		return err
	}
	to, err := strconv.Atoi(strings.Split(rangeStr, "-")[1])
	if err != nil {
		fmt.Println("Error converting to value:", err)
		return err
	}

	fl.ranges = append(fl.ranges, FreshRange{from: from, to: to})
	return nil
}

type FreshRange struct {
	from int
	to   int
}

// combine overlapping ranges
func (fl *FreshList) optimizeRanges() {
	for {
		newRanges := []FreshRange{}
		improved := false
		for _, r := range fl.ranges {
			merged := false
			for i, or := range newRanges {
				if or.contains(r.from) || or.contains(r.to) || r.contains(or.from) || r.contains(or.to) {
					newFrom := int(math.Min(float64(or.from), float64(r.from)))
					newTo := int(math.Max(float64(or.to), float64(r.to)))
					newRanges[i] = FreshRange{from: newFrom, to: newTo}
					merged = true
					improved = true
					break
				}
			}
			if !merged {
				newRanges = append(newRanges, r)
			}
		}
		fl.ranges = newRanges
		if !improved {
			break
		}
	}
}

// Count total number of fresh ingredients in all ranges
func (fl *FreshList) countFreshIds() int {
	total := 0
	for _, r := range fl.ranges {
		total += r.to - r.from + 1
	}
	return total
}

// Print all ranges
func (fl *FreshList) printRanges() {
	for _, r := range fl.ranges {
		fmt.Printf("%d-%d\n", r.from, r.to)
	}
}

func (i *FreshRange) contains(ingridient int) bool {
	return ingridient >= i.from && ingridient <= i.to
}

func (fl *FreshList) anyRangeContains(ingredient int) bool {
	for _, r := range fl.ranges {
		if r.contains(ingredient) {
			return true
		}
	}
	return false
}

func part1(input []string) (int, error) {
	freshList := FreshList{}
	freshListComplete := false
	freshIngredients := 0

	for _, line := range input {
		if line == "\n" || line == "" {
			freshListComplete = true
			//fmt.Println("Ranges: ", len(freshList))
			continue
		}

		if !freshListComplete {
			if freshList.newRange(line) != nil {
				return 0, fmt.Errorf("error adding range: %s", line)
			}

		} else {
			ingredient, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error converting ingredient value:", err)
				return 0, err
			}
			if freshList.anyRangeContains(ingredient) {
				freshIngredients++
			}
		}
	}

	return freshIngredients, nil
}

func part2(input []string) (int, error) {
	freshList := FreshList{}

	for _, line := range input {
		if line == "\n" || line == "" {
			break
		}

		if freshList.newRange(line) != nil {
			return 0, fmt.Errorf("error adding range: %s", line)
		}
	}

	freshList.printRanges()
	fmt.Println("Fresh Ids before: ", freshList.countFreshIds(), "\n")
	freshList.optimizeRanges()
	freshList.printRanges()
	fmt.Println("Fresh Ids after:  ", freshList.countFreshIds())

	return freshList.countFreshIds(), nil
}

func main() {
	fmt.Println("Day 5!")

	input, err := helper.ReadFileToInput("day5/input")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	result, err := part2(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result: ", result)
}
