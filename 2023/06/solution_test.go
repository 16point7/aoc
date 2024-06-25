package main

import "testing"

const input1 = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	want := 288
	got := part1(input1)

	if got != want {
		t.Fatalf("Invalid result for input %s. got %d, want %d", input1, got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

const input2 = `Time:      7  15   30
Distance:  9  40  200`

func TestPart2(t *testing.T) {
	want := 71503
	got := part2(input2)

	if got != want {
		t.Fatalf("Invalid result for input %s. got %d, want %d", input2, got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
