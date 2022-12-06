package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	stacksInput, _ := os.ReadFile("stacks.txt")
	stacksP1 := parseStacks(string(stacksInput))
	stacksP2 := parseStacks(string(stacksInput))

	commandsInput, _ := os.ReadFile("commands.txt")
	commands := parseCommands(string(commandsInput))

	for _, cmd := range commands {
		stacksP1.move(cmd)
		stacksP2.move9001(cmd)
	}

	var p1 string
	for _, stack := range stacksP1 {
		p1 += stack[len(stack)-1]
	}

	var p2 string
	for _, stack := range stacksP2 {
		p2 += stack[len(stack)-1]
	}

	println(p1)
	println(p2)
}

type Stack []string

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	if l == 0 {
		panic("cannot pop empty stack!")
	}

	return s[:l-1], s[l-1]
}

func (s Stack) Push(item string) Stack {
	return append(s, item)
}

type Stacks []Stack

type Cmd struct {
	n, from, to int
}

func (stacks Stacks) move9001(cmd Cmd) {
	stackFrom := stacks[cmd.from]
	stackTo := stacks[cmd.to]

	stacks[cmd.from] = stackFrom[:len(stackFrom)-cmd.n]
	stacks[cmd.to] = append(stackTo, stackFrom[len(stackFrom)-cmd.n:]...)
}

func (stacks Stacks) move(cmd Cmd) {
	for i := 0; i < cmd.n; i++ {
		var crate string
		stacks[cmd.from], crate = stacks[cmd.from].Pop()
		stacks[cmd.to] = stacks[cmd.to].Push(crate)
	}
}

func parseCommands(input string) []Cmd {
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	commands := []Cmd{}
	for _, line := range lines {
		var n, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to)
		commands = append(commands, Cmd{n: n, from: from - 1, to: to - 1})
	}

	return commands
}

func parseStacks(input string) Stacks {
	lines := strings.Split(input, "\n")
	reverse(lines) // note reverse to go bottom up!

	var stacks Stacks
	for i, r := range lines[0] {
		if !unicode.IsUpper(r) {
			continue
		}

		stack := []string{string(r)}
		for _, line := range lines[1:] {
			if string(line[i]) != " " {
				stack = append(stack, string(line[i]))
			}
		}

		stacks = append(stacks, stack)
	}

	return stacks
}

func reverse[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
