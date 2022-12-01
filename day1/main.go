package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	data, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(data)

	scores := []int{}
	count := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			scores = append(scores, count)
			count = 0
			continue
		}

		n, err := strconv.Atoi(line)
		check(err)

		count += n
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	println(scores[0] + scores[1] + scores[2])
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
