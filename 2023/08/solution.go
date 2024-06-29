package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}

const (
	begin = "AAA"
	end   = "ZZZ"
	left  = 'L'
)

func part1(input string) int {
	moves, after, _ := strings.Cut(input, "\n")
	input = after[1:]

	m := make(map[string][2]string)
	for {
		before, after, found := strings.Cut(input, "\n")
		key, left, right := parseLine(before)
		m[key] = [2]string{left, right}
		if !found {
			break
		}
		input = after
	}

	count := 0
	for i, cur := 0, begin; cur != end; i = (i + 1) % len(moves) {
		if moves[i] == left {
			cur = m[cur][0]
		} else {
			cur = m[cur][1]
		}
		count++
	}
	return count
}

func parseLine(line string) (key, left, right string) {
	key, after, _ := strings.Cut(line, " = ")
	l, r, _ := strings.Cut(after, ", ")
	left, right = l[1:], r[:len(r)-1]
	return
}

func part2(input string) int {
	return 0
}
