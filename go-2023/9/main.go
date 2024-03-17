package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLines(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func allZeroes(v []int) bool {
	for _, n := range v {
		if n != 0 {
			return false
		}
	}
	return true
}

func extrapolate(line string) int {
	strValues := strings.Fields(line)
	var values []int
	for _, v := range strValues {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		values = append(values, n)
	}

	rLine := []int{}
	currLine := values
	for !allZeroes(currLine) {
		rLine = append(rLine, currLine[len(currLine)-1])
		prevLine := []int{}

		for i := 1; i < len(currLine); i++ {
			prevLine = append(prevLine, currLine[i]-currLine[i-1])
		}
		currLine = prevLine

	}

	sum := 0
	for _, n := range rLine {
		sum += n
	}

	return sum
}

func part1() {
	lines := getLines("input.txt")
	sum := 0
	for _, line := range lines {
		sum += extrapolate(line)
	}
	fmt.Println("Part 1:", sum)
}

func extrapolateBack(line string) int {
	strValues := strings.Fields(line)
	var values []int
	for _, v := range strValues {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		values = append(values, n)
	}

	lLine := []int{}
	currLine := values
	for !allZeroes(currLine) {
		lLine = append([]int{currLine[0]}, lLine...)
		prevLine := []int{}

		for i := 1; i < len(currLine); i++ {
			prevLine = append(prevLine, currLine[i]-currLine[i-1])
		}
		currLine = prevLine

	}

	res := 0
	for _, n := range lLine {
		res = n - res
	}

	return res
}

func part2() {
	lines := getLines("input.txt")
	sum := 0
	for _, line := range lines {
		eb := extrapolateBack(line)
		sum += eb
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	fmt.Println("Day 9")
	part1()
	part2()
}
