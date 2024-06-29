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

	return process(begin, m, moves, func(cur string) bool { return cur == end })
}

func parseLine(line string) (key, left, right string) {
	key, after, _ := strings.Cut(line, " = ")
	l, r, _ := strings.Cut(after, ", ")
	left, right = l[1:], r[:len(r)-1]
	return
}

func process(cur string, m map[string][2]string, moves string, finished func(string) bool) int {
	count := 0
	for i := 0; !finished(cur); i = (i + 1) % len(moves) {
		if moves[i] == left {
			cur = m[cur][0]
		} else {
			cur = m[cur][1]
		}
		count++
	}
	return count
}

func part2(input string) int {
	moves, after, _ := strings.Cut(input, "\n")
	input = after[1:]

	m := make(map[string][2]string)
	curs := make([]string, 0, 8)
	for {
		before, after, found := strings.Cut(input, "\n")

		key, left, right := parseLine(before)
		m[key] = [2]string{left, right}

		if c := key[2]; c == 'A' {
			curs = append(curs, key)
		}

		if !found {
			break
		}

		input = after
	}

	count := process(curs[0], m, moves, trailingZ)
	for _, cur := range curs[1:] {
		count = lcm(count, process(cur, m, moves, trailingZ))
	}
	return count
}

func trailingZ(cur string) bool {
	return cur[2] == 'Z'
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
