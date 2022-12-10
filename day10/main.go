package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	cycles := []int{1}
	for _, instruction := range lines {
		cycles = append(cycles, cycle(instruction)...)
		//fmt.Printf("%s %v\n", instruction, cycles)
	}

	part1 := signalStrength(cycles, 20) + signalStrength(cycles, 60) + signalStrength(cycles, 100) + signalStrength(cycles, 140) + signalStrength(cycles, 180) + signalStrength(cycles, 220)
	fmt.Println(part1)

	// Part 2
	spritePosition := cycles[0]
	for i, add := range cycles[1:] {
		if overlapsSprite(i, spritePosition) {
			print("#")
		} else {
			print(".")
		}

		spritePosition += add

		if (i+1)%40 == 0 {
			println()
		}
	}

}

func overlapsSprite(cycle, spritePosition int) bool {
	cyclePos := cycle % 40
	return abs(cyclePos-spritePosition) <= 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func cycle(instruction string) []int {
	switch instruction {
	case "noop":
		return []int{0}
	default:
		var add int
		fmt.Sscanf(instruction, "addx %d", &add)
		return []int{0, add}
	}
}

func signalStrength(cycles []int, n int) int {
	return sum(cycles[:n]) * n
}

func sum(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}

	return total
}
