package main

import (
	"fmt"
	"strconv"

	"github.com/Reicher/AoC25/helper"
)

func main() {
	fmt.Println("Day 1!")

	input, err := helper.ReadFileToInput("day1/input")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	dial := 50
	zeroPos := 0

	fmt.Println("The dial starts by pointing at 50.")

	for _, line := range input {
		dir := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing steps:", err)
			return
		}

		var change int
		switch dir {
		case 'L':
			change = -1
		case 'R':
			change = 1
		}

		// Loop over each step to check for zero crossing
		for i := 0; i < steps; i++ {
			dial = (dial + change + 100) % 100
			if dial == 0 {
				zeroPos += 1
			}
		}
	}

	fmt.Println("Final dial position:", dial)
	fmt.Println("Number of times at zero position:", zeroPos)

}
