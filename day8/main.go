package main

import (
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	width := strings.IndexByte(trimmed, byte('\n'))
	trees := []byte(strings.Replace(trimmed, "\n", "", -1))

	var count, topScenicScore int
	for i := range trees {
		if isVisible(trees, i, width) {
			count++
		}

		topScenicScore = max(topScenicScore, scenicScore(trees, i, width))
	}

	println(count)
	println(topScenicScore)
}

func isVisible(input []byte, pos int, width int) bool {
	val := input[pos]
	left, right, up, down := getNeighbours(input, pos, width)
	return every(left, isLt(val)) || every(right, isLt(val)) || every(up, isLt(val)) || every(down, isLt(val))
}

func scenicScore(input []byte, pos int, width int) int {
	val := input[pos]
	left, right, up, down := getNeighbours(input, pos, width)
	return viewingDistance(reverse(left), val) * viewingDistance(right, val) * viewingDistance(reverse(up), val) * viewingDistance(down, val)
}

func getNeighbours(trees []byte, pos int, width int) (left, right, up, down []byte) {
	rem := pos % width
	left = trees[pos-rem : pos]
	right = trees[pos+1 : pos+(width-rem)]
	up = getCol(trees[:pos-rem], rem, width)
	down = getCol(trees[pos+(width-rem):], rem, width)

	return left, right, up, down
}

func isLt(val byte) func(item byte) bool {
	return func(item byte) bool {
		return item < val
	}
}

func viewingDistance(trees []byte, val byte) int {
	count := 0
	for _, tree := range trees {
		count++

		if !isLt(val)(tree) {
			return count
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getCol(input []byte, rowPos int, width int) []byte {
	items := []byte{}
	for i := rowPos; i < len(input); i += width {
		items = append(items, input[i])
	}

	return items
}

func every[A any](items []A, fn func(item A) bool) bool {
	for _, item := range items {
		if !fn(item) {
			return false
		}
	}

	return true
}

func reverse[A any](s []A) []A {
	items := []A{}
	for i := len(s) - 1; i >= 0; i -= 1 {
		items = append(items, s[i])
	}

	return items
}
