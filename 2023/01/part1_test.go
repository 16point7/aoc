package main

import "testing"

const input1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
	want := 142
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
