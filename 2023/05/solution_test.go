package main

import "testing"

const input1 = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestPart1(t *testing.T) {
	want := 35
	got := part1(input1)

	if got != want {
		t.Fatalf("Invalid result for input %s. got %d, want %d", input1, got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input1)
	}
}

const input2 = ``

func TestPart2(t *testing.T) {
	want := 0
	got := part2(input2)

	if got != want {
		t.Fatalf("Invalid result for input %s. got %d, want %d", input2, got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input2)
	}
}
