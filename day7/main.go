package main

import (
	"fmt"
	"strings"

	"github.com/Reicher/AoC25/helper"
)

func part1(input []string) (int, error) {

	splits := 0
	manifold := []string{}
	manifold = append(manifold, input[0])

	fmt.Println(manifold[0])

	for row := 1; row < len(input); row++ {
		next_manifoid_layer := []rune(input[row])
		for col := 0; col < len(next_manifoid_layer); col++ {
			char := string(next_manifoid_layer[col])
			char_above := string(manifold[row-1][col])
			if char_above == "S" || char_above == "|" {
				switch char {
				case ".":
					next_manifoid_layer[col] = '|'
				case "^":
					next_manifoid_layer[col-1] = '|'
					next_manifoid_layer[col+1] = '|'
					splits++
				}
			}
		}
		manifold = append(manifold, string(next_manifoid_layer))
		fmt.Println(manifold[row])
	}

	return splits, nil
}

func part2(input []string) (int, error) {

	manifold := [][]int{}

	rays := make([]int, len(input[0]))
	rays[strings.Index(input[0], "S")] = 1

	manifold = append(manifold, rays)

	for row := 1; row < len(input); row++ {
		rays = make([]int, len(input[row]))
		for col := 0; col < len(input[row]); col++ {
			rays_above := manifold[row-1][col]
			current_char := string(input[row][col])
			if rays_above > 0 {
				switch current_char {
				case ".":
					rays[col] += rays_above
				case "^":
					rays[col-1] += rays_above
					rays[col+1] += rays_above
				}
			}
		}
		manifold = append(manifold, rays)
		fmt.Println(manifold[row])
	}

	// Count rays at the bottom
	realities := 0
	for _, val := range manifold[len(manifold)-1] {
		realities += val
	}

	return realities, nil
}

func main() {
	fmt.Println("Day 7!")

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
