package main

import (
	"os"
	"strings"
)

// A, X = Rock (1)
// B, Y = Paper (2)
// C, Z = Scissors (3)
// W=6, D=3, L=0

const (
	rock     = 1
	paper    = 2
	scissors = 3

	win  = 6
	draw = 3
	loss = 0
)

type Hand int
type Outcome int

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	total := 0
	for _, line := range lines {
		them, me := parseLine(line)
		s := score(them, me)
		total += int(s)
	}

	println(total)
}

func score(them Hand, requiredOutcome Outcome) Hand {
	me := requiredHand(them, requiredOutcome)

	switch {
	case me == weakness(them):
		return loss + me
	case me == counter(them):
		return win + me
	default:
		return draw + me
	}
}

func requiredHand(them Hand, desiredOutcome Outcome) Hand {
	switch desiredOutcome {
	case win:
		return counter(them)
	case loss:
		return weakness(them)
	default: // draw
		return them

	}
}

func counter(hand Hand) Hand {
	switch hand {
	case rock:
		return paper
	case paper:
		return scissors
	default: // scissors
		return rock
	}
}

func weakness(hand Hand) Hand {
	switch hand {
	case rock:
		return scissors
	case paper:
		return rock
	default: // scissors
		return paper
	}
}

func parseLine(line string) (Hand, Outcome) {
	parts := strings.Split(line, " ")
	return parseHand(parts[0]), parseOutcome(parts[1])
}

func parseHand(letter string) Hand {
	switch letter {
	case "A":
		return rock
	case "B":
		return paper
	default:
		return scissors
	}
}

func parseOutcome(letter string) Outcome {
	switch letter {
	case "X":
		return loss
	case "Y":
		return draw
	default:
		return win
	}
}
