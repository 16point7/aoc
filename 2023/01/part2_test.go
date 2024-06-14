package main

import "testing"

const input2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPart2(t *testing.T) {
	want := 281
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
