package main

import "testing"

func TestPart1(t *testing.T) {
	input := []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	want := 142
	got := part1(input)

	if got != want {
		t.Fatalf("Invalid result for input %s. got %d, want %d", input, got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	for i := 0; i < b.N; i++ {
		part1(input)
	}
}
