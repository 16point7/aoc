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

const start byte = 'S'

func part1(input string) int {
	m, s := parse(input)
	j, i := find(m, start)
	length := getLength(m, s, j, i)
	return length / 2
}

func parse(input string) ([]string, [][]bool) {
	cols := strings.IndexByte(input, '\n')
	rows := (len(input) + 1) / (cols + 1)

	m := make([]string, rows)
	buf := make([]bool, cols*rows)
	s := make([][]bool, rows)
	for j := 0; j < rows-1; j++ {
		m[j], input = input[:cols], input[cols+1:]
		s[j], buf = buf[:cols], buf[cols:]
	}
	m[rows-1] = input[:cols]
	s[rows-1] = buf[:cols]

	return m, s
}

func find(m []string, c byte) (j, i int) {
	for j := 0; j < len(m); j++ {
		for i := 0; i < len(m[j]); i++ {
			if m[j][i] == c {
				return j, i
			}
		}
	}
	return -1, -1
}

type coord struct {
	j, i int
}

func getLength(m []string, s [][]bool, j, i int) int {
	s[j][i] = true
	length := 1

	stack := make([]coord, 0, 128)

	next := getNext(m, j, i)
	stack = append(stack, coord{next[0].j, next[0].i})
	stack = append(stack, coord{next[1].j, next[1].i})

	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		s[c.j][c.i] = true
		length++

		next := getNext(m, c.j, c.i)

		if !s[next[0].j][next[0].i] {
			stack = append(stack, coord{next[0].j, next[0].i})
		}
		if !s[next[1].j][next[1].i] {
			stack = append(stack, coord{next[1].j, next[1].i})
		}
	}

	return length
}

func getNext(m []string, j, i int) [2]coord {
	switch m[j][i] {
	case '|':
		return [2]coord{{j + 1, i}, {j - 1, i}}
	case '-':
		return [2]coord{{j, i - 1}, {j, i + 1}}
	case 'L':
		return [2]coord{{j - 1, i}, {j, i + 1}}
	case 'J':
		return [2]coord{{j - 1, i}, {j, i - 1}}
	case '7':
		return [2]coord{{j, i - 1}, {j + 1, i}}
	case 'F':
		return [2]coord{{j + 1, i}, {j, i + 1}}
	case 'S':
		res := [2]coord{}
		r := 0

		if nextJ := j - 1; nextJ > -1 && (m[nextJ][i] == '|' || m[nextJ][i] == '7' || m[nextJ][i] == 'F') {
			res[r] = coord{nextJ, i}
			r++
		}

		if nextJ := j + 1; nextJ < len(m) && (m[nextJ][i] == '|' || m[nextJ][i] == 'L' || m[nextJ][i] == 'J') {
			res[r] = coord{nextJ, i}
			r++
			if r >= len(res) {
				return res
			}
		}

		if nextI := i - 1; nextI > -1 && (m[j][nextI] == '-' || m[j][nextI] == 'L' || m[j][nextI] == 'F') {
			res[r] = coord{j, nextI}
			r++
			if r >= len(res) {
				return res
			}
		}

		if nextI := i + 1; nextI < len(m[j]) && (m[j][nextI] == '-' || m[j][nextI] == 'L' || m[j][nextI] == '7') {
			res[r] = coord{j, nextI}
			r++
		}

		return res
	}
	return [2]coord{}
}

func part2(input string) int {
	return 0
}
