package main

import "testing"

func TestPart2(t *testing.T) {
	input := []byte(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	want := 281
	got := part2(input)

	if got != want {
		t.Fatalf("Invalid result for input %s. got %d, want %d", input, got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := []byte(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
