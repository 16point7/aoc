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

func part1(input string) int {
	res := 0
	nums := make([]int, 0, 64)
	for {
		before, after, found := strings.Cut(input, "\n")
		res += processLine(before, nums, true)
		if !found {
			break
		}
		input = after
	}
	return res
}

func part2(input string) int {
	res := 0
	nums := make([]int, 0, 64)
	for {
		before, after, found := strings.Cut(input, "\n")
		res += processLine(before, nums, false)
		if !found {
			break
		}
		input = after
	}
	return res
}

func processLine(line string, nums []int, forward bool) int {
	for len(line) > 0 {
		before, after, _ := strings.Cut(line, " ")
		nums = append(nums, atoi(before))
		line = after
	}

	if forward {
		end := len(nums) - 1
		for {
			for i := 0; i < end; i++ {
				nums[i] = nums[i+1] - nums[i]
			}
			if allZeroes(nums[:end]) {
				break
			}
			end--
		}
	} else {
		start := 0
		for {
			for i := len(nums) - 1; i > start; i-- {
				nums[i] = nums[i-1] - nums[i]
			}
			if allZeroes(nums[start+1:]) {
				break
			}
			start++
		}
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func atoi(input string) int {
	i := 0

	sign := 1
	if input[0] == '-' {
		sign = -1
		i++
	}

	res := 0
	for i < len(input) {
		res = res*10 + int(input[i]-'0')*sign
		i++
	}

	return res
}

func allZeroes(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			return false
		}
	}
	return true
}
