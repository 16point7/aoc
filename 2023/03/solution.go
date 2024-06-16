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
	fmt.Println(part2(input))
}

func part1(input string) int {
	matrix := strings.Split(input, "\n")

	res := 0
	for j, row := range matrix {
		part, isPart := 0, false
		for i := 0; i < len(row); i++ {
			if c := row[i]; isDigit(c) {
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

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

const gear byte = '*'

func part2(input string) int {
	matrix := strings.Split(input, "\n")

	res := 0
	for j, row := range matrix {
		for i := 0; i < len(row); i++ {
			if row[i] == gear {
				gearRatio, isGear := inspectSurrounding(matrix, j, i)
				if isGear {
					res += gearRatio
				}
			}
		}
	}
	return res
}

func inspectSurrounding(matrix []string, j, i int) (gearRatio int, isGear bool) {
	count, gearRatio := 0, 1

	left := i - 1
	safeLeft := left > -1
	if safeLeft && isDigit(matrix[j][left]) {
		count++
		partNum := getPartNum(matrix[j], left)
		gearRatio *= partNum
	}
	right := i + 1
	safeRight := right < len(matrix[j])
	if safeRight && isDigit(matrix[j][right]) {
		count++
		partNum := getPartNum(matrix[j], right)
		gearRatio *= partNum
		if count == 2 {
			isGear = true
			return
		}
	}

	if up := j - 1; up > -1 {
		if isDigit(matrix[up][i]) {
			count++
			partNum := getPartNum(matrix[up], i)
			gearRatio *= partNum
			if count == 2 {
				isGear = true
				return
			}
		} else {
			if safeLeft && isDigit(matrix[up][left]) {
				count++
				partNum := getPartNum(matrix[up], left)
				gearRatio *= partNum
				if count == 2 {
					isGear = true
					return
				}
			}
			if safeRight && isDigit(matrix[up][right]) {
				count++
				partNum := getPartNum(matrix[up], right)
				gearRatio *= partNum
				if count == 2 {
					isGear = true
					return
				}
			}
		}
	}

	if down := j + 1; down > -1 {
		if isDigit(matrix[down][i]) {
			count++
			partNum := getPartNum(matrix[down], i)
			gearRatio *= partNum
			if count == 2 {
				isGear = true
				return
			}
		} else {
			if safeLeft && isDigit(matrix[down][left]) {
				count++
				partNum := getPartNum(matrix[down], left)
				gearRatio *= partNum
				if count == 2 {
					isGear = true
					return
				}
			}
			if safeRight && isDigit(matrix[down][right]) {
				count++
				partNum := getPartNum(matrix[down], right)
				gearRatio *= partNum
				if count == 2 {
					isGear = true
					return
				}
			}
		}
	}
	return
}

func getPartNum(row string, i int) int {
	for i > -1 && isDigit(row[i]) {
		i--
	}
	i++
	res := 0
	for i < len(row) && isDigit(row[i]) {
		res = res*10 + int(row[i]-'0')
		i++
	}
	return res
}
