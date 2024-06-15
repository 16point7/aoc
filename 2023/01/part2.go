package main

func part2(input string) int {
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
	txt string
	val int
}{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

func part2line(line string) int {
	var l int
l2r:
	for i := 0; i < len(line); i++ {
		if c := line[i]; c >= '0' && c <= '9' {
			l = int(c - '0')
			break
		}
		for _, n := range numbers {
			if matchLeft(line[i:], n.txt) {
				l = n.val
				break l2r
			}
		}

	}
	var r int
r2l:
	for i := len(line) - 1; i > -1; i-- {
		if c := line[i]; c >= '0' && c <= '9' {
			r = int(c - '0')
			break
		}
		for _, n := range numbers {
			if matchRight(line[:i+1], n.txt) {
				r = n.val
				break r2l
			}
		}
	}
	return l*10 + r
}

func matchLeft(a, b string) bool {
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

func matchRight(a, b string) bool {
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
