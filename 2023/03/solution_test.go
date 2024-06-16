package main

import "testing"

const input1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
	want := 4361
	got := part1(input1)

	if want != got {
		t.Fatalf("Invalid result for input \n%s. got %d, want %d", input1, got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input1)
	}
}

func TestPart2(t *testing.T) {
	want := 467835
	got := part2(input1)

	if want != got {
		t.Fatalf("Invalid result for input \n%s. got %d, want %d", input1, got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input1)
	}
}
