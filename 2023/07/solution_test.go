package main

import "testing"

const input1 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1(t *testing.T) {
	want := 6440
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
