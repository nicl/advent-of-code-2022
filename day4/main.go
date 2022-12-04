package main

import (
	"fmt"
	"os"
	"strings"
)

type Range struct {
	from, to int
}

func (r Range) contains(b Range) bool {
	return r.from <= b.from && r.to >= b.to
}

func (r Range) intersects(b Range) bool {
	return r.from <= b.to && b.from <= r.to
}

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	count := 0
	for _, line := range lines {
		a, b := getRanges(line)
		if a.contains(b) || b.contains(a) {
			count++
		}
	}

	println(count)
}

func part2(lines []string) {
	count := 0
	for _, line := range lines {
		a, b := getRanges(line)
		if a.intersects(b) {
			count++
		}
	}

	println(count)
}

func getRanges(line string) (Range, Range) {
	var aFrom, aTo, bFrom, bTo int
	fmt.Sscanf(line, "%d-%d,%d-%d", &aFrom, &aTo, &bFrom, &bTo)
	return Range{aFrom, aTo}, Range{bFrom, bTo}
}
