package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Reicher/AoC25/helper"
)

func operation(series []int, op string) int {
	switch op {
	case "+":
		sum := 0
		for _, num := range series {
			sum += num
		}
		return sum
	case "*":
		prod := 1
		for _, num := range series {
			prod *= num
		}
		return prod
	default:
		return 0
	}
}

func part1(input []string) (int, error) {
	series := [][]int{}
	answers := []int{}

	for i, line := range input {
		elements := strings.Fields(line)
		for j, el := range elements {
			if el == "*" || el == "+" {
				answers = append(answers, operation(series[j], el))
				continue
			}
			num, err := strconv.Atoi(el)
			if err != nil {
				fmt.Printf("Error converting element at line %d, position %d: %v\n", i, j, err)
				return 0, err
			}
			if len(series) <= j {
				series = append(series, []int{})
			}
			series[j] = append(series[j], num)
		}
	}

	fmt.Println(answers)

	return operation(answers, "+"), nil
}

func part2(input []string) (int, error) {
	answers := []int{}
	nums := []int{}

	for col := len(input[0]) - 1; col >= 0; col-- {
		num_str := ""

		for row := 0; row < len(input); row++ {
			char := string(input[row][col])

			if row == len(input)-1 && num_str != "" {
				num, err := strconv.Atoi(num_str)
				if err != nil {
					fmt.Printf("Error converting cephalopod math string at column %d: %v\n", col, err)
					return 0, err
				}
				nums = append(nums, num)
			}

			if char == "*" || char == "+" {
				answers = append(answers, operation(nums, char))
				nums = []int{}
			} else if char != " " {
				num_str += char
			}
		}
	}

	return operation(answers, "+"), nil
}

func main() {
	fmt.Println("Day 6!")

	input, err := helper.ReadFileToInput("input")
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
