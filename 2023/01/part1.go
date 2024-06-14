package main

func part1(input []byte) int {
	var res int
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			res += part1line(input[:i])
			input = input[i+1:]
			i = 0
		}
	}
	res += part1line(input)
	return res
}

func part1line(line []byte) int {
	var l int
	for i := 0; i < len(line); i++ {
		if c := line[i]; c >= '0' && c <= '9' {
			l = int(c - '0')
			break
		}
	}
	var r int
	for i := len(line) - 1; i > -1; i-- {
		if c := line[i]; c >= '0' && c <= '9' {
			r = int(c - '0')
			break
		}
	}
	return l*10 + r
}
