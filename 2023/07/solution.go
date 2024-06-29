package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}

type hand struct {
	bid   int
	rank  int
	cards []int
}

func part1(input string) int {
	hands := parseHands(input)

	slices.SortFunc(hands, func(h1, h2 hand) int {
		if rankDiff := h1.rank - h2.rank; rankDiff != 0 {
			return rankDiff
		}
		for i := 0; i < len(h1.cards); i++ {
			if valDiff := h1.cards[i] - h2.cards[i]; valDiff != 0 {
				return valDiff
			}
		}
		return 0
	})

	res := 0
	for i, hand := range hands {
		res += (i + 1) * hand.bid
	}
	return res
}

func parseHands(input string) []hand {
	hands := make([]hand, 0, 32)
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			hands = append(hands, parseHand(input[:i]))
			input = input[i+1:]
			i = 0
		}
	}
	hands = append(hands, parseHand(input))
	return hands
}

func parseHand(line string) hand {
	before, after, _ := strings.Cut(line, " ")

	bid := 0
	for i := 0; i < len(after); i++ {
		bid = bid*10 + int(after[i]-'0')
	}

	cards := getCards(before)
	rank := getRank(cards)

	return hand{bid: bid, cards: cards, rank: rank}
}

func getCards(input string) []int {
	cards := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		cards[i] = getValue(input[i])
	}
	return cards
}

func getValue(card byte) int {
	switch card {
	case '2':
		return 0
	case '3':
		return 1
	case '4':
		return 2
	case '5':
		return 3
	case '6':
		return 4
	case '7':
		return 5
	case '8':
		return 6
	case '9':
		return 7
	case 'T':
		return 8
	case 'J':
		return 9
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	default:
		return -1
	}
}

const (
	highCard int = 1 + iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func getRank(cards []int) int {
	c1 := [13]int{}
	for _, c := range cards {
		c1[c]++
	}

	c2 := [6]int{}
	for _, c := range c1 {
		c2[c]++
	}

	switch {
	case c2[5] > 0:
		return fiveOfAKind
	case c2[4] > 0:
		return fourOfAKind
	case c2[3] > 0:
		if c2[2] > 0 {
			return fullHouse
		}
		return threeOfAKind
	case c2[2] == 2:
		return twoPair
	case c2[2] == 1:
		return onePair
	default:
		return highCard
	}
}

func part2(input string) int {
	return 0
}
