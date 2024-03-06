package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (cardNum int, input []string, output []string) {
	game := strings.Split(line, ":")
	if len(game) != 2 {
		for _, val := range game {
			fmt.Println(val)
		}
		log.Fatal("Bad split on ':' - ", game)
	}
	inputOutput := strings.Split(game[1], "|")
	if len(inputOutput) != 2 {
		log.Fatal("Bad split on '|' - ", inputOutput)
	}
	cardNum, err := strconv.Atoi(strings.Fields(game[0])[1])
	if err != nil {
		panic(err)
	}
	input = strings.Fields(inputOutput[0])
	output = strings.Fields(inputOutput[1])
	return
}

func getPoints(input []string, output []string) int {
	linePoints := 0
	for _, valueIn := range input {
		for _, valueOut := range output {
			if valueIn == valueOut {
				if linePoints == 0 {
					linePoints++
				} else {
					linePoints *= 2
				}
			}
		}
	}
	return linePoints
}

func getMatches(input []string, output []string) int {
	matches := 0
	for _, valueIn := range input {
		for _, valueOut := range output {
			if valueIn == valueOut {
				matches++
			}
		}
	}
	return matches
}

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	totalPoints := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, input, output := parseLine(scanner.Text())
		totalPoints += getPoints(input, output)
	}

	fmt.Println("Part 1:", totalPoints)
}

type card struct {
	count   int
	input   []string
	output  []string
	matches int
}

type cards struct {
	tracker map[int]*card
	maxCard int
}

func part2() {
	fmt.Print("Part 2: ")
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	var ct cards
	ct.tracker = make(map[int]*card)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cardNum, input, output := parseLine(scanner.Text())
		ct.maxCard = cardNum
		ptr, ok := ct.tracker[cardNum]
		if ok {
			(*ptr).count += 1
		} else {
			matches := getMatches(input, output)
			ct.tracker[cardNum] = &card{1, input, output, matches}
		}
	}

	for i := range ct.maxCard {
		cardNum := i + 1
		c := (*ct.tracker[cardNum])
		matches := c.matches
		for matches > 0 {
			currCard, ok := ct.tracker[cardNum+matches]
			if ok {
				(*currCard).count += c.count
			} else {
				log.Fatal("No card found:", cardNum, matches)
			}
			matches--
		}
	}

	total := 0
	for _, v := range ct.tracker {
		total += (*v).count
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("Day 4")
	part1()
	part2()
}
