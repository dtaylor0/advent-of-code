package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GameIsPossible(game string, redCubes int, greenCubes int, blueCubes int) bool {
	redIdx := strings.Index(game, "red")
	greenIdx := strings.Index(game, "green")
	blueIdx := strings.Index(game, "blue")

	if redIdx > -1 {
		var lastIdx int
		for i := redIdx - 2; i > -1; i-- {
			if game[i] == ' ' {
				lastIdx = i+1
				break
			}
		}
		digit := game[lastIdx : redIdx-1]
		numRedCubes, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}

		if numRedCubes > redCubes {
			return false
		}
	}
	if greenIdx > -1 {
		var lastIdx int
		for i := greenIdx - 2; i > -1; i-- {
			if game[i] == ' ' {
				lastIdx = i+1
				break
			}
		}
		digit := game[lastIdx : greenIdx-1]
		numGreenCubes, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}

		if numGreenCubes > greenCubes {
			return false
		}
	}
	if blueIdx > -1 {
		var lastIdx int
		for i := blueIdx - 2; i > -1; i-- {
			if game[i] == ' ' {
				lastIdx = i+1
				break
			}
		}
		digit := game[lastIdx : blueIdx-1]
		numBlueCubes, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}

		if numBlueCubes > blueCubes {
			return false
		}
	}
	return true
}

func GetGameID(game string) int {
	idx := strings.Index(game, ":")
	var firstIdx int
	for i := idx - 1; i > -1; i-- {
		if game[i] == ' ' {
			firstIdx = i+1
			break
		}
	}
	digit := game[firstIdx:idx]
	gid, err := strconv.Atoi(digit)
	if err != nil {
		panic(err)
	}

	return gid
}

func PartOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if GameIsPossible(line, 12, 13, 14) {
			totalSum += GetGameID(line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", totalSum)
}

func main() {
	PartOne()
}
