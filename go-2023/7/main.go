package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(d byte) bool {
	return d > '0' && d <= '9'
}

type Card struct {
	Value byte
}

type Hand struct {
	Cards   []Card
	RawHand string
	Value   byte
	Bid     int
}

func (a Card) gt(b Card) bool {
	values := "23456789TJQKA"
	return strings.IndexByte(values, a.Value) > strings.IndexByte(values, b.Value)
}

func (a Card) lt(b Card) bool {
	values := "23456789TJQKA"
	return strings.IndexByte(values, a.Value) < strings.IndexByte(values, b.Value)
}

func (a Card) eq(b Card) bool {
	values := "23456789TJQKA"
	return strings.IndexByte(values, a.Value) == strings.IndexByte(values, b.Value)
}

func CalculateHand(h Hand) byte {
	cardCounts := make(map[byte]int)
	for _, c := range h.Cards {
		cardCounts[c.Value]++
	}

	if len(cardCounts) == 1 {
		return '5'
	} else if len(cardCounts) == 2 {
		for _, v := range cardCounts {
			if v == 3 {
				return 'F'
			}
			if v == 4 {
				return '4'
			}
		}
	} else if len(cardCounts) == 3 {
		for _, v := range cardCounts {
			if v == 3 {
				return '3'
			} else if v == 2 {
				return '2'
			}
		}
	} else if len(cardCounts) == 4 {
		return '1'
	} else {
		return '0'
	}

	return '2'
}

func (a Hand) gt(b Hand) bool {
	hands := "0123F45"
	aIdx := strings.IndexByte(hands, a.Value)
	bIdx := strings.IndexByte(hands, b.Value)
	if aIdx != bIdx {
		return aIdx > bIdx
	}
	for i := 0; i < len(a.Cards); i++ {
		if a.Cards[i].gt(b.Cards[i]) {
			return true
		} else if a.Cards[i].lt(b.Cards[i]) {
			return false
		}
	}
	return false
}

func parseHands(filename string) []Hand {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)
		rawHand := fields[0]
		rawBid := fields[1]
		bid, err := strconv.Atoi(rawBid)
		if err != nil {
			panic(err)
		}

		cards := []Card{}
		for i := 0; i < len(rawHand); i++ {
			cards = append(cards, Card{rawHand[i]})
		}
		hand := Hand{cards, rawHand, 'a', bid}
		hand.Value = CalculateHand(hand)

		inserted := false
		for i := 0; i < len(hands); i++ {
			if hands[i].gt(hand) {
				hands = append(hands[:i], append([]Hand{hand}, hands[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			hands = append(hands, hand)
		}
	}
	return hands
}

func part1() {
	fmt.Print("Part 1:")
	hands := parseHands("input.txt")
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.Bid
	}
	fmt.Println(total)
}

func part2() {
	fmt.Print("Part 2:")
}

func main() {
	fmt.Println("Day 7")
	part1()
	part2()
}
