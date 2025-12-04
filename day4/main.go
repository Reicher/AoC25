package main

import (
	"fmt"

	"github.com/Reicher/AoC25/helper"
)

// Check all 8 directions for adjacent rolls
func adjecentRolls(row int, col int, input []string) int {
	rolls := 0
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(input) && newCol >= 0 && newCol < len(input[newRow]) {
			if input[newRow][newCol] == '@' {
				rolls++
			}
		}
	}
	return rolls
}

func part1(input []string) (int, error) {
	available := 0
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '@' && adjecentRolls(row, col, input) < 4 {
				available++
			}
		}
	}
	return available, nil
}

func part2(rollMap []string) (int, error) {
	available := -1
	removed := 0

	fmt.Println("Initial state:")

	for available == -1 || available > 0 {
		available = 0

		//Print current roll map
		for _, row := range rollMap {
			fmt.Println(row)
		}

		newRollMap := []string{}
		for row := range rollMap {
			newRow := ""
			for col := range rollMap[row] {
				if rollMap[row][col] == '@' && adjecentRolls(row, col, rollMap) < 4 {
					available++
					removed++
					newRow += "x"
				} else {
					newRow += string(rollMap[row][col])
				}
			}
			newRollMap = append(newRollMap, newRow)
		}
		rollMap = newRollMap
		fmt.Println("\nRemove", available, "rolls of paper:")

	}
	return removed, nil
}

func main() {
	fmt.Println("Day 4!")

	input, err := helper.ReadFileToInput("day4/input")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	rolls, err := part2(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(rolls, "is accesable by the forklift.")
}
