package main

func part2(input []byte) int {
	var res int
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			res += part2line(input[:i])
			input = input[i+1:]
			i = 0
		}
	}
	res += part2line(input)
	return res
}

var numbers = []struct {
	txt []byte
	val int
}{
	{[]byte("one"), 1},
	{[]byte("two"), 2},
	{[]byte("three"), 3},
	{[]byte("four"), 4},
	{[]byte("five"), 5},
	{[]byte("six"), 6},
	{[]byte("seven"), 7},
	{[]byte("eight"), 8},
	{[]byte("nine"), 9},
}

func part2line(line []byte) int {
	var l int
outer1:
	for i := 0; i < len(line); i++ {
		if c := line[i]; c >= '0' && c <= '9' {
			l = int(c - '0')
			break
		}
		for _, n := range numbers {
			if matchLeft(line[i:], n.txt) {
				l = n.val
				break outer1
			}
		}

	}
	var r int
outer2:
	for i := len(line) - 1; i > -1; i-- {
		if c := line[i]; c >= '0' && c <= '9' {
			r = int(c - '0')
			break
		}
		for _, n := range numbers {
			if matchRight(line[:i+1], n.txt) {
				r = n.val
				break outer2
			}
		}
	}
	return l*10 + r
}

func matchLeft(a, b []byte) bool {
	if len(b) > len(a) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func matchRight(a, b []byte) bool {
	if len(b) > len(a) {
		return false
	}
	for i, j := len(b)-1, len(a)-1; i > -1; i, j = i-1, j-1 {
		if a[j] != b[i] {
			return false
		}
	}
	return true
}
