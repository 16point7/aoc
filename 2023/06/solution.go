package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
