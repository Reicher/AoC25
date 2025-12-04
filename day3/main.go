package main

import (
	"fmt"
	"strconv"

	"github.com/Reicher/AoC25/helper"
)

func part1(input []string) (int, error) {
	total_joltage := 0
	for _, bank := range input {

		first := 0
		first_index := 0
		for i := 0; i < len(bank)-1; i++ {
			num := int(bank[i] - '0')
			if first < num {
				first = num
				first_index = i
			}
		}
		second := 0
		for i := first_index + 1; i < len(bank); i++ {
			num := int(bank[i] - '0')
			if second < num {
				second = num
			}
		}
		bank_jolt, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
		if err != nil {
			fmt.Println("Error converting bank joltage:", err)
			return 0, err
		}
		fmt.Println("Bank joltage:", bank_jolt)
		total_joltage += bank_jolt
	}

	return total_joltage, nil
}

func part2(input []string) (int, error) {
	total_joltage := 0
	for _, bank := range input {
		var joltages [12]int
		next_index := 0

		for batt := 0; batt < len(joltages); batt++ {
			//fmt.Println("Finding joltage for battery", batt)
			for i := next_index; i <= len(bank)-len(joltages)+batt; i++ {
				//fmt.Println("trying...index", i)
				num := int(bank[i] - '0')
				if joltages[batt] < num {
					joltages[batt] = num
					next_index = i + 1
					//fmt.Println("Set joltage for", batt, "to", num)
				}
			}

		}

		bank_str_jolts := ""
		for _, val := range joltages {
			bank_str_jolts += strconv.Itoa(val)
		}

		bank_jolt, err := strconv.Atoi(bank_str_jolts)
		if err != nil {
			fmt.Println("Error converting bank joltage:", err)
			return 0, err
		}
		fmt.Println("Bank joltage:", bank_jolt)
		total_joltage += bank_jolt
	}

	return total_joltage, nil
}

func main() {
	fmt.Println("Day 3!")

	input, err := helper.ReadFileToInput("day3/input")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	joltage, err := part2(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Total output joltage:", joltage)
}
