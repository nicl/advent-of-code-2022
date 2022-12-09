package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

func (p Point) Is(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func sign(n int) int {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	default:
		return 0
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func moveTail(head, tail Point) Point {
	dX := head.X - tail.X
	dY := head.Y - tail.Y

	switch {
	case abs(dX) <= 1 && abs(dY) <= 1:
		return tail
	default:
		return Point{tail.X + sign(dX), tail.Y + sign(dY)}
	}
}

func find[A any](items []A, fn func(item A) bool) (int, bool) {
	for i, item := range items {
		if fn(item) {
			return i, true
		}
	}

	return 0, false
}

func printGrid(knots []Point, width int) {
	for y := width; y >= 0; y-- {
		for x := 0; x < width; x++ {
			match, ok := find(knots, func(knot Point) bool {
				return knot.Is(Point{x, y})
			})

			if ok {
				print(match)
			} else {
				print(".")
			}
		}

		println()
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	knots := [10]Point{} // change to 2 for P1
	seen := map[Point]bool{}

	for _, cmd := range lines {
		parts := strings.Split(cmd, " ")
		steps, _ := strconv.Atoi(parts[1])
		dir := parts[0]

		//println(cmd)

		for i := 0; i < steps; i++ {
			head := knots[0]

			switch dir {
			case "L":
				knots[0] = Point{head.X - 1, head.Y}
			case "R":
				knots[0] = Point{head.X + 1, head.Y}
			case "U":
				knots[0] = Point{head.X, head.Y + 1}
			case "D":
				knots[0] = Point{head.X, head.Y - 1}
			}

			for i, knot := range knots {
				if i == 0 {
					continue // skip head
				}

				updated := moveTail(knots[i-1], knot)
				knots[i] = updated

				if i == len(knots)-1 {
					seen[updated] = true
				}
			}

			fmt.Println(knots)
			printGrid(knots[:], 10)
			//fmt.Scanln()
		}
	}

	println(len(seen))
}
