package main

import (
	"bufio"
	"fmt"
	"os"
)

func getLines(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func connectsTo(maze *[]string, position []int, dir []int) {

}

func part1() {
}

func part2() {
}

func main() {
	fmt.Println("Day 10")
	part1()
	part2()
}
