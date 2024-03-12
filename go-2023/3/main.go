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

func getNumber(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func part1() {
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
				sum += getNumber(currNumber)
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
			sum += getNumber(currNumber)
			checkIfPart = true
		}
	}

	fmt.Println("Part 1:", sum)
}

func contains(s []int, v int) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

func addToGearTracker(value int, i int, j int, gearTracker map[[2]int][]int) {
	idx := [2]int{i, j}
	if gearTracker[idx] != nil && !contains(gearTracker[idx], value) {
		gearTracker[idx] = append(gearTracker[idx], value)
	} else {
		gearTracker[idx] = []int{value}
	}
}

func part2() {
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

	gearTracker := make(map[[2]int][]int)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		isNumber := false
		checkIfPart := true
		currNumber := ""
		var gears [][2]int

		for j := 0; j < len(line); j++ {
			char, _ := get(&lines, i, j, 0, 0)
			isNumber = charIsNumber(char)

			if !isNumber && !checkIfPart {
				for _, gear := range gears {
					addToGearTracker(getNumber(currNumber), gear[0], gear[1], gearTracker)
				}
				gears = [][2]int{}
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
						gears = append(gears, [2]int{i + dir[0], j + dir[1]})
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
			for _, gear := range gears {
				addToGearTracker(getNumber(currNumber), gear[0], gear[1], gearTracker)
			}
			gears = [][2]int{}
			checkIfPart = true
		}
	}

	sum := 0
	for _, adj := range gearTracker {
		if len(adj) == 2 {
			sum += adj[0] * adj[1]
		}
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	fmt.Println("Day 3")
	part1()
	part2()
}
