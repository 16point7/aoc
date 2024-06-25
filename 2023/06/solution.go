package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}

type data struct {
	time, dist float64
}

func part1(input string) int {
	data := getData(input)

	res := 1
	for _, d := range data {
		sqrtb2minus4ac := math.Sqrt(d.time*d.time - 4*d.dist)
		r1 := int(math.Ceil(((d.time - sqrtb2minus4ac) / 2)))
		r2 := int(math.Floor(((d.time + sqrtb2minus4ac) / 2)))
		res *= r2 - r1 + 1
	}
	return res
}

func getData(input string) []data {
	first, second, _ := strings.Cut(input, "\n")
	_, timeline, _ := strings.Cut(first, ":")
	_, distline, _ := strings.Cut(second, ":")

	res := make([]data, 0, 4)

	t, d := 0, 0
	for t < len(timeline) && d < len(distline) {
		timenum, nextT := nextNum(timeline, t)
		distnum, nextD := nextNum(distline, d)
		res = append(res, data{time: float64(timenum), dist: float64(distnum) + 0.01})
		t, d = nextT, nextD
	}

	return res
}

func nextNum(line string, i int) (num, lastI int) {
	for line[i] == ' ' {
		i++
	}

	for lastI = i; lastI < len(line) && line[lastI] != ' '; lastI++ {
		if c := line[lastI]; c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
	}

	return
}

func part2(input string) int {
	first, second, _ := strings.Cut(input, "\n")
	time := getData2(first)
	dist := getData2(second) + 0.01
	sqrtb2minus4ac := math.Sqrt(time*time - 4*dist)
	r1 := int(math.Ceil(((time - sqrtb2minus4ac) / 2)))
	r2 := int(math.Floor(((time + sqrtb2minus4ac) / 2)))
	return r2 - r1 + 1
}

func getData2(input string) float64 {
	num := 0
	for i := 0; i < len(input); i++ {
		if c := input[i]; c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
	}
	return float64(num)
}
