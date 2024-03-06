package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	totalPoints := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.SplitAfter(line, ":")
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
		input := strings.Fields(inputOutput[0])
		output := strings.Fields(inputOutput[1])

		totalPoints += getPoints(input, output)
	}

	fmt.Println("Part 1:", totalPoints)
}

type card struct {
    input []string
    output []string
    count int
}


func ([]card) increment(cardNum int, n int) {
    
}

func part2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

    totalCards := 0
    var cards []
    var q [][]string
	scanner := bufio.NewScanner(file)
    
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.SplitAfter(line, ":")
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
		input := strings.Fields(inputOutput[0])
		output := strings.Fields(inputOutput[1])
	}

}

func main() {
	fmt.Println("Day 4")
	part1()
	part2()
}
