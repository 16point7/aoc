package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1(input))
}

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
		if isPart {
			res += part
		}

	}
	return res
}

func hasSymbol(matrix []string, j, i int) bool {
	left := i - 1
	safeLeft := left > -1
	if safeLeft && isSymbol(matrix[j][left]) {
		return true
	}
	right := i + 1
	safeRight := right < len(matrix[j])
	if safeRight && isSymbol(matrix[j][right]) {
		return true
	}

	if up := j - 1; up > -1 {
		if isSymbol(matrix[up][i]) || safeLeft && isSymbol(matrix[up][left]) || safeRight && isSymbol(matrix[up][right]) {
			return true
		}
	}

	if down := j + 1; down < len(matrix) {
		if isSymbol(matrix[down][i]) || safeLeft && isSymbol(matrix[down][left]) || safeRight && isSymbol(matrix[down][right]) {
			return true
		}
	}

	return false
}

func isSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}
