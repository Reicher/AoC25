package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Reicher/AoC25/helper"
)

func part1(input []string) (int, error) {
	ranges := strings.Split(input[0], ",")
	invalid_ids := []int{}
	for _, r := range ranges {
		debug_txt := fmt.Sprintf("- %s :", r)
		from, err := strconv.Atoi(strings.Split(r, "-")[0])
		if err != nil {
			fmt.Println("Error converting from value:", err)
			return 0, err
		}
		to, err := strconv.Atoi(strings.Split(r, "-")[1])
		if err != nil {
			fmt.Println("Error converting to value:", err)
			return 0, err
		}
		for i := from; i <= to; i++ {
			str_num := strconv.Itoa(i)
			str_len := len(str_num)

			// skip numbers with odd length
			if str_len%2 != 0 {
				continue
			}
			if str_num[:str_len/2] == str_num[str_len/2:] {

				invalid_ids = append(invalid_ids, i)
				debug_txt += fmt.Sprintf(" %d", i)
			}
		}
		fmt.Println(debug_txt)
	}
	// sum invalid ids
	sum := 0
	for _, id := range invalid_ids {
		sum += id
	}
	return sum, nil
}

func part2(input []string) (int, error) {
	ranges := strings.Split(input[0], ",")
	invalid_ids := []int{}
	for _, r := range ranges {
		debug_txt := fmt.Sprintf("- %s :", r)
		from, err := strconv.Atoi(strings.Split(r, "-")[0])
		if err != nil {
			fmt.Println("Error converting from value:", err)
			return 0, err
		}
		to, err := strconv.Atoi(strings.Split(r, "-")[1])
		if err != nil {
			fmt.Println("Error converting to value:", err)
			return 0, err
		}
		for i := from; i <= to; i++ {
			str_num := strconv.Itoa(i)
			str_len := len(str_num)

			// loop from 1 to len/2 and repeating pattern check
			for j := 1; j <= str_len/2; j++ {
				occurrences := strings.Count(str_num, str_num[:j])
				if occurrences*j == str_len {
					invalid_ids = append(invalid_ids, i)
					debug_txt += fmt.Sprintf(" %d", i)
					break // this number is invalid, no need to check further
				}
			}
		}
		fmt.Println(debug_txt)
	}

	// sum invalid ids
	sum := 0
	for _, id := range invalid_ids {
		sum += id
	}
	return sum, nil
}

func main() {
	fmt.Println("Day 2!")

	input, err := helper.ReadFileToInput("day2/input")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	sum, err := part2(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sum of invalid IDs:", sum)
}
