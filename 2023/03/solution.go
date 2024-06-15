package main

import (
	"strings"
)

func part1(input string) int {
	matrix := strings.Split(input, "\n")

	res := 0
	for j, row := range matrix {
		part, isPart := 0, false
		for i, c := range row {
			if c >= '0' && c <= '9' {
				part = part*10 + int(c-'0')
				isPart = isPart || hasSymbol(matrix, j, i)
				continue
			}
			if isPart {
				res += part
			}
			part, isPart = 0, false
		}
	}
	return res
}

var directions = []struct{ j, i int }{
	{1, 0},
	{1, -1},
	{1, 1},
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{0, -1},
	{0, 1},
}

func hasSymbol(matrix []string, j, i int) bool {
	for _, d := range directions {
		nextJ, nextI := j+d.j, i+d.i
		if nextJ < 0 || nextI < 0 || nextJ >= len(matrix) || nextI >= len(matrix[nextJ]) {
			continue
		}
		if c := matrix[nextJ][nextI]; c != '.' && (c < '0' || c > '9') {
			return true
		}
	}
	return false
}
