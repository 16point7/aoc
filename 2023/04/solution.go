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
	res := 0
	for _, line := range strings.Split(input, "\n") {
		before, after1, _ := strings.Cut(line, " | ")
		_, after2, _ := strings.Cut(before, ": ")
		winningNums := getWinningNums(after2)
		numMatch := countMatches(winningNums, after1)
		if numMatch > 0 {
			res += 1 << (numMatch - 1)
		}
	}
	return res
}

func getWinningNums(s string) [100]bool {
	res := [100]bool{}
	num, numFound := 0, false
	for i := 0; i < len(s); i++ {
		if c := s[i]; c >= '0' && c <= '9' {
			num, numFound = num*10+int(c-'0'), true
		} else if numFound {
			res[num] = true
			num, numFound = 0, false
		}
	}
	if numFound {
		res[num] = true
	}
	return res
}

func countMatches(winningNumbers [100]bool, s string) int {
	res := 0
	num, numSet := 0, false
	for i := 0; i < len(s); i++ {
		if c := s[i]; c >= '0' && c <= '9' {
			num, numSet = num*10+int(c-'0'), true
		} else {
			if numSet && winningNumbers[num] {
				res++
			}
			num, numSet = 0, false
		}
	}
	if numSet && winningNumbers[num] {
		res++
	}
	return res
}
