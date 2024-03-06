package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convert(value int, convMap [][3]int) int {
	for _, line := range convMap {
		dest := line[0]
		src := line[1]
		r := line[2]

		if value >= src && value < src+r {
			return dest + value - src
		}
	}
	return value
}

func toInts(sl []string) []int {
	var res []int
	for _, s := range sl {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Cannot convert to int: ", s)
		}
		res = append(res, i)
	}
	return res
}

func startsWithDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] >= byte('0') && s[0] <= byte('9') {
		return true
	}
	return false
}

func getNextMap(scanner *bufio.Scanner) [][3]int {
	var res [][3]int
	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		if startsWithDigit(text) {
			break
		}
	}

	for startsWithDigit(text) {
		fields := strings.Fields(text)
		first, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		third, err := strconv.Atoi(fields[2])
		if err != nil {
			panic(err)
		}
		row := [3]int{first, second, third}
		res = append(res, row)

		if !scanner.Scan() {
			break
		}
		text = scanner.Text()
	}
	return res
}

func part1() {
	fmt.Print("Part 1: ")
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	src := toInts(strings.Fields(text)[1:])
	var dest []int
	for scanner.Scan() {
		nextMap := getNextMap(scanner)
		for _, i := range src {
			dest = append(dest, convert(i, nextMap))
		}
		src = dest
		dest = []int{}
	}

	minDest := src[0]
	for _, v := range src {
		if v < minDest {
			minDest = v
		}
	}
	fmt.Println(minDest)
}

func main() {
	fmt.Println("Day 5")
	part1()
}
