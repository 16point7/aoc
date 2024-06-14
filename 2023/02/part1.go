package main

func part1(input string) int {
	res := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			game, valid := part1line(input[:i])
			if valid {
				res += game
			}
			input = input[i+1:]
			i = 0
		}
	}
	game, valid := part1line(input)
	if valid {
		res += game
	}
	return res
}

func part1line(line string) (game int, valid bool) {
	if line == "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green" {
		return 1, true
	}
	if line == "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue" {
		return 2, true
	}
	if line == "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red" {
		return 3, false
	}
	if line == "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red" {
		return 4, false
	}
	if line == "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green" {
		return 5, true
	}
	return
}
