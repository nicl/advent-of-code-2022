package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Node struct {
	children []*Node
	name     string
	size     int
}

func (n Node) Size() int {
	sum := n.size
	for _, child := range n.children {
		sum += child.Size()
	}

	return sum
}

func (n *Node) Find(target string) (*Node, bool) {
	if target == "/" {
		return n, true
	}

	parts := strings.Split(target[1:], "/")

	for _, child := range n.children {
		if child.name == parts[0] {
			if len(parts) == 1 {
				return child, true
			}

			return child.Find("/" + strings.Join(parts[1:], "/"))
		}
	}

	return nil, false
}

func (n *Node) Add(target string, children ...*Node) {
	node, ok := n.Find(target)
	if !ok {
		panic("target not found: " + target)
	}

	node.children = append(node.children, children...)
}

func (n *Node) Visit(fn func(n *Node)) {
	fn(n)

	for _, child := range n.children {
		child.Visit(fn)
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))
	lines := strings.Split(trimmed, "\n")

	at := "/"
	root := Node{name: "/"}
	pos := 0

	for {
		if pos >= len(lines) {
			break
		}

		line := lines[pos]

		switch {
		case isCdRoot(line):
			at = "/"
			pos++
		case isCdUp(line):
			at = path.Dir(at)
			pos++
		case isCdDown(line):
			name := parseDir(line)
			at = path.Join(at, name)
			pos++
		case isLs(line):
			nodes, offset := parseLs(lines[pos+1:])
			root.Add(at, nodes...)
			pos += offset
		}
	}

	sum := 0
	root.Visit(func(n *Node) {
		if len(n.children) == 0 {
			return // skip
		}

		if n.Size() <= 100000 {
			sum += n.Size()
		}
	})

	println(sum) // part 1

	total := 70000000
	required := 30000000
	free := total - root.Size()
	extraRequired := required - free

	matchSize := 0
	root.Visit(func(n *Node) {
		if len(n.children) == 0 {
			return // skip
		}

		unset := matchSize == 0
		bigEnough := n.Size() > extraRequired
		smallerThanAlt := n.Size() < matchSize

		if bigEnough && (unset || smallerThanAlt) {
			matchSize = n.Size()
		}
	})
	println(matchSize) // part 2
}

func isCdRoot(line string) bool {
	return line == "$ cd /"
}

func isCdUp(line string) bool {
	return line == "$ cd .."
}

func isCdDown(line string) bool {
	return strings.HasPrefix(line, "$ cd")
}

func isLs(line string) bool {
	return strings.HasPrefix(line, "$ ls")
}

func parseLs(lines []string) ([]*Node, int) {
	nodes := []*Node{}
	offset := 0
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "$") || line == "":
			return nodes, offset + 1
		case strings.HasPrefix(line, "dir"):
			node := parseLsDir(line)
			nodes = append(nodes, &node)
		default:
			node := parseLsFile(line)
			nodes = append(nodes, &node)
		}

		offset++
	}

	return nodes, offset + 1
}

func parseDir(line string) string {
	return strings.TrimPrefix(line, "$ cd ")
}

func parseLsDir(line string) Node {
	name := strings.TrimPrefix(line, "dir ")
	return Node{name: name}
}

func parseLsFile(line string) Node {
	var size int
	var name string
	fmt.Sscanf(line, "%d %s", &size, &name)
	return Node{name: name, size: size}
}
