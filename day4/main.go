package main

import (
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from, to int
}

func (r Range) contains(b Range) bool {
	return r.from <= b.from && r.to >= b.to
}

func (r Range) intersects(b Range) bool {
	return (r.from >= b.from && r.from <= b.to) || (r.to >= b.from && r.to <= b.to) || (r.from < b.from && r.to > b.to)
}

func main() {
	part1()
	part2()
}

func part1() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	count := 0
	for _, line := range lines {
		a, b := getRanges(line)
		if a.contains(b) || b.contains(a) {
			count++
		}
	}

	println(count)
}

func part2() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

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
	parts := strings.Split(line, ",")
	return getRange(parts[0]), getRange(parts[1])
}

func getRange(part string) Range {
	parts := strings.Split(part, "-")
	from, _ := strconv.Atoi(parts[0])
	to, _ := strconv.Atoi(parts[1])
	return Range{from, to}
}
