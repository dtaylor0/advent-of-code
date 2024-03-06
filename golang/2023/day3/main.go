package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func charIsNumber(b byte) bool {
	if b < byte('0') || b > byte('9') {
		return false
	}
	return true
}

func isDot(b byte) bool {
	if b == byte('.') {
		return true
	}
	return false
}

func get(lines *[]string, i int, j int, y int, x int) (byte, error) {
	if i+y < 0 || i+y >= len(*lines) {
		return 0, errors.New("IndexOutOfBounds")
	}
	if j+x < 0 || j+x >= len((*lines)[i]) {
		return 0, errors.New("IndexOutOfBounds")
	}
	return (*lines)[i+y][j+x], nil
}

func main() {
	fmt.Println("Day 3")
	dirs := [][]int{
		{1, -1}, {1, 0}, {1, 1},
		{0, -1}, {0, 1},
		{-1, -1}, {-1, 0}, {-1, 1},
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		isNumber := false
		checkIfPart := true
		currNumber := ""

		for j := 0; j < len(line); j++ {
			char, _ := get(&lines, i, j, 0, 0)
			isNumber = charIsNumber(char)

			if !isNumber && !checkIfPart {
				val, err := strconv.Atoi(currNumber)
				if err != nil {
					log.Fatal(err)
				}

				sum += val
				checkIfPart = true
			}

			if !isNumber {
				currNumber = ""
			}

			if isNumber && checkIfPart {
				isPart := false
				for _, dir := range dirs {
					dirChar, err := get(&lines, i, j, dir[0], dir[1])
					if err != nil {
						continue
					}
					if !isDot(dirChar) && !charIsNumber(dirChar) {
						isPart = true
						break
					}
				}
				if isPart {
					checkIfPart = false
				}
			}

			if isNumber {
				currNumber += string([]byte{char})
			}
		}

		if isNumber && !checkIfPart {
			val, err := strconv.Atoi(currNumber)
			if err != nil {
				log.Fatal(err)
			}

			sum += val
			checkIfPart = true
		}
	}

	fmt.Println("Part 1:", sum)
}
