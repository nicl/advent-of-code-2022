package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) bool
	targets   map[bool]int
	seenItems int
}

type Troop []Monkey

type Move struct {
	item int
	to   int
}

func main() {
	/* 	troop := Troop{
		{
			items:     []int{79, 98},
			operation: func(i int) int { return i * 19 },
			test:      func(i int) bool { return i%23 == 0 },
			targets:   map[bool]int{true: 2, false: 3},
		},
		{
			items:     []int{54, 65, 75, 74},
			operation: func(i int) int { return i + 6 },
			test:      func(i int) bool { return i%19 == 0 },
			targets:   map[bool]int{true: 2, false: 0},
		},
		{
			items:     []int{79, 60, 97},
			operation: func(i int) int { return i * i },
			test:      func(i int) bool { return i%13 == 0 },
			targets:   map[bool]int{true: 1, false: 3},
		},

		{
			items:     []int{74},
			operation: func(i int) int { return i + 3 },
			test:      func(i int) bool { return i%17 == 0 },
			targets:   map[bool]int{true: 0, false: 1},
		},
	} */

	troop := Troop{
		{
			items:     []int{59, 74, 65, 86},
			operation: func(i int) int { return i * 19 },
			test:      func(i int) bool { return i%7 == 0 },
			targets:   map[bool]int{true: 6, false: 2},
		},
		{
			items:     []int{62, 84, 72, 91, 68, 78, 51},
			operation: func(i int) int { return i + 1 },
			test:      func(i int) bool { return i%2 == 0 },
			targets:   map[bool]int{true: 2, false: 0},
		},
		{
			items:     []int{78, 84, 96},
			operation: func(i int) int { return i + 8 },
			test:      func(i int) bool { return i%19 == 0 },
			targets:   map[bool]int{true: 6, false: 5},
		},

		{
			items:     []int{97, 86},
			operation: func(i int) int { return i * i },
			test:      func(i int) bool { return i%3 == 0 },
			targets:   map[bool]int{true: 1, false: 0},
		},
		{
			items:     []int{50},
			operation: func(i int) int { return i + 6 },
			test:      func(i int) bool { return i%13 == 0 },
			targets:   map[bool]int{true: 3, false: 1},
		},
		{
			items:     []int{73, 65, 69, 65, 51},
			operation: func(i int) int { return i * 17 },
			test:      func(i int) bool { return i%11 == 0 },
			targets:   map[bool]int{true: 4, false: 7},
		},
		{
			items:     []int{69, 82, 97, 93, 82, 84, 58, 63},
			operation: func(i int) int { return i + 5 },
			test:      func(i int) bool { return i%5 == 0 },
			targets:   map[bool]int{true: 5, false: 7},
		},
		{
			items:     []int{81, 78, 82, 76, 79, 80},
			operation: func(i int) int { return i + 3 },
			test:      func(i int) bool { return i%17 == 0 },
			targets:   map[bool]int{true: 3, false: 4},
		},
	}

	roundSize := len(troop) * 10000
	lcm := 7 * 2 * 19 * 3 * 13 * 11 * 5 * 17 // no time for parsing today!

	for i := 0; i < roundSize; i++ {
		n := i % len(troop)
		monkey := troop[n]
		moves := round(monkey, lcm)
		troop[n].items = []int{}
		troop[n].seenItems += len(moves)
		troop = applyMoves(troop, moves)
	}

	fmt.Println(monkeyBusiness(troop))
}

func monkeyBusiness(troop Troop) int {
	sort.Slice(troop, func(i, j int) bool {
		return troop[i].seenItems > troop[j].seenItems
	})

	return troop[0].seenItems * troop[1].seenItems
}

func applyMoves(troop Troop, moves []Move) Troop {
	for _, move := range moves {
		troop[move.to].items = append(troop[move.to].items, move.item)
	}

	return troop
}

func round(monkey Monkey, lcm int) []Move {
	moves := []Move{}

	for _, item := range monkey.items {
		worryLevel := item
		worryLevel = monkey.operation(worryLevel) % lcm
		//worryLevel = worryLevel / 3
		moves = append(moves, Move{worryLevel, monkey.targets[monkey.test(worryLevel)]})
	}

	return moves
}
