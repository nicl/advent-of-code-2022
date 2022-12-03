package main

import (
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	batches := [][]string{}
	for i := 0; i < len(lines); i += 3 {
		batches = append(batches, lines[i:i+3])
	}

	sum := 0
	for _, batch := range batches {
		first := batch[0]
		second := batch[1]
		third := batch[2]

		for _, c := range first {
			if strings.ContainsRune(second, c) && strings.ContainsRune(third, c) {
				sum += score(c)
				break
			}
		}
	}

	println(sum)
}

func part1() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	sum := 0
	for _, line := range lines {
		first := line[:len(line)/2]
		second := line[len(line)/2:]

		// find item that exists in both
		for _, c := range first {
			if strings.ContainsRune(second, c) {
				sum += score(c)
				break
			}
		}
	}

	println(sum)
}

func score(c rune) int {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	pos := strings.IndexRune(letters, c)
	if pos == -1 {
		panic("Unexpected pos!")
	}

	return pos + 1
}
