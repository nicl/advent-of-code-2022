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

type Edge struct {
	from, to Point
}

func (e Edge) String() string {
	return fmt.Sprintf("%s -> %s", e.from, e.to)
}

type Grid [][]byte

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	_, end, grid := parseGrid(trimmed)

	startingPoints := findStartingPoints(grid)

	minSteps := 0
	for _, start := range startingPoints {
		got, ok := steps(grid, start, end)
		if ok && (minSteps == 0 || got < minSteps) {
			minSteps = got
		}
	}

	println(minSteps)
}

// This is simpler than Djikstra as all edges are the same length. We take
// advantage of this fact throughout.
func steps(grid Grid, start, end Point) (int, bool) {
	links := map[Point]Point{} // point->previous, the optimal path
	priorityQueue := []Edge{{start, start}}

	found := false

	for {
		if len(priorityQueue) == 0 {
			break
		}

		for _, edge := range priorityQueue {
			priorityQueue = priorityQueue[1:] // a 'stack'

			if edge.to.Is(end) {
				links[edge.to] = edge.from
				found = true
				break
			}

			_, ok := links[edge.to]
			if !ok {
				links[edge.to] = edge.from
				es := filterUnseen(links, filterEdges(grid, edge.to))
				priorityQueue = append(priorityQueue, es...)
			}
		}
	}

	return traceSteps(links, end, 0), found
}

func filterEdges(grid Grid, from Point) []Edge {
	es := []Edge{}
	height := len(grid)
	width := len(grid[0])
	offsets := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, offset := range offsets {
		newX := from.X + offset.X
		newY := from.Y + offset.Y

		if (newX >= 0 && newX < width) && (newY >= 0 && newY < height) && allowedMove(grid[from.Y][from.X], grid[newY][newX]) {
			es = append(es, Edge{from, Point{newX, newY}})
		}
	}

	return es
}

func findStartingPoints(grid Grid) []Point {
	found := []Point{}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'a' || grid[y][x] == 'S' {
				found = append(found, Point{x, y})
			}
		}
	}

	return found
}

func filterUnseen(links map[Point]Point, edges []Edge) []Edge {
	es := []Edge{}

	for _, edge := range edges {
		_, ok := links[edge.to]
		if !ok {
			es = append(es, edge)
		}
	}

	return es
}

func traceSteps(links map[Point]Point, point Point, stepsSoFar int) int {
	prev := links[point]
	if prev.Is(point) {
		return stepsSoFar
	}

	return traceSteps(links, prev, stepsSoFar+1)
}

func allowedMove(from, to byte) bool {
	if from == 'S' {
		from = 'a'
	}

	if to == 'E' {
		to = 'z'
	}

	return to <= from+1
}

func parseGrid(input string) (Point, Point, Grid) {
	lines := strings.Split(input, "\n")

	var start, end Point
	grid := make([][]byte, len(lines))
	for y, line := range lines {
		for x, b := range []byte(line) {
			grid[y] = append(grid[y], b)
			if b == 'S' {
				start = Point{x, y}
			}
			if b == 'E' {
				end = Point{x, y}
			}
		}
	}

	return start, end, grid
}
