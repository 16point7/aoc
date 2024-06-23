package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}

func part1(input string) int {
	seeds, after := getSeeds(input)
	mappings := getMappings(after)
	return getMinLoc(seeds, mappings)
}

func getSeeds(input string) ([]int, string) {
	seeds := make([]int, 0, 32)

	i := 0

	for input[i] != ':' {
		i++
	}

	i += 2

	for num := 0; true; i++ {
		if c := input[i]; c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else {
			seeds = append(seeds, num)
			if c == '\n' {
				break
			}
			num = 0
		}
	}

	return seeds, input[i:]
}

type mapping struct {
	startI, startF int
	endI           int
}

func getMappings(input string) [7][]mapping {
	mappings := [7][]mapping{}
	maps := strings.SplitN(input, ":\n", 8)
	for j := 0; j < len(mappings); j++ {
		i, m := 0, maps[j+1]
		for i < len(m) && m[i] != '\n' {
			startI, startF, endI := 0, 0, 0
			num, numCount := 0, 0
			for i < len(m) && m[i] != '\n' {
				if c := m[i]; c >= '0' && c <= '9' {
					num = num*10 + int(c-'0')
				} else {
					switch numCount {
					case 0:
						endI = num
					case 1:
						startI = num
					}
					num = 0
					numCount++
				}
				i++
			}
			startF = startI + num
			mappings[j] = append(mappings[j], mapping{startI: startI, startF: startF, endI: endI})
			i++
		}
	}
	return mappings
}

func getMinLoc(seeds []int, mappings [7][]mapping) int {
	for _, m := range mappings {
		slices.SortFunc(m, func(a, b mapping) int {
			return a.startI - b.startI
		})
	}

	minLoc := math.MaxInt
	for _, seed := range seeds {
		cur := seed
		for _, mapping := range mappings {
			cur = findNext(cur, mapping)
		}
		minLoc = min(cur, minLoc)
	}
	return minLoc
}

func findNext(cur int, mapping []mapping) int {
	for i := 0; i < len(mapping) && cur >= mapping[i].startI; i++ {
		if cur < mapping[i].startF {
			return cur - mapping[i].startI + mapping[i].endI
		}
	}
	return cur
}

func part2(input string) int {
	return 0
}
