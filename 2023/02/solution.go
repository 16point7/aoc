package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	res := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			game, valid := part1line(input[:i])
			if valid {
				res += game
			}
			input = input[i+1:]
			i = 0
		}
	}
	game, valid := part1line(input)
	if valid {
		res += game
	}
	return res
}

func part1line(line string) (int, bool) {
	before, after, _ := strings.Cut(line[5:], ": ")
	game, _ := strconv.Atoi(before)

	for _, set := range strings.Split(after, "; ") {
		for _, marble := range strings.Split(set, ", ") {
			before, after, _ := strings.Cut(marble, " ")
			count, _ := strconv.Atoi(before)
			var limit int
			switch after {
			case "red":
				limit = 12
			case "green":
				limit = 13
			case "blue":
				limit = 14
			}
			if count > limit {
				return game, false
			}
		}
	}
	return game, true
}

func part2(input string) int {
	res := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			res += part2line(input[:i])
			input = input[i+1:]
			i = 0
		}
	}
	res += part2line(input)
	return res
}

func part2line(line string) int {
	_, after, _ := strings.Cut(line, ": ")

	r, g, b := 1, 1, 1
	for _, set := range strings.Split(after, "; ") {
		for _, marble := range strings.Split(set, ", ") {
			before, after, _ := strings.Cut(marble, " ")
			count, _ := strconv.Atoi(before)
			switch after {
			case "red":
				r = max(r, count)
			case "green":
				g = max(g, count)
			case "blue":
				b = max(b, count)
			}
		}
	}
	return r * g * b
}
