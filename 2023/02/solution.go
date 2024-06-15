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

const (
	rlimit = 12
	glimit = 13
	blimit = 14
)

func part1line(line string) (int, bool) {
	before, after, _ := strings.Cut(line[5:], ": ")
	game, _ := strconv.Atoi(before)

	for _, set := range strings.Split(after, "; ") {
		for _, marble := range strings.Split(set, ", ") {
			before, after, _ := strings.Cut(marble, " ")
			count, _ := strconv.Atoi(before)
			r, g, b := 0, 0, 0
			switch after {
			case "red":
				r = count
			case "green":
				g = count
			case "blue":
				b = count
			}
			if r > rlimit || g > glimit || b > blimit {
				return game, false
			}
		}
	}
	return game, true
}
