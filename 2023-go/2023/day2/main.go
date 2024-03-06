package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GetNum(s *string, lastIdx int) int {
	var firstIdx int
	for i := lastIdx - 1; i > -1; i-- {
		if (*s)[i] == ' ' {
			firstIdx = i + 1
			break
		}
	}
	digit := (*s)[firstIdx:lastIdx]
	num, err := strconv.Atoi(digit)
	if err != nil {
		panic(err)
	}
	return num
}

func GameIsPossible(game *string, redCubes int, greenCubes int, blueCubes int) bool {
	for i, c := range *game {
		if c == 'r' {
			if (*game)[i+1] == 'e' && (*game)[i+2] == 'd' {
				numRedCubes := GetNum(game, i-1)
				if numRedCubes > redCubes {
					return false
				}
			}
		}
		if c == 'g' {
			if (*game)[i+1] == 'r' && (*game)[i+2] == 'e' && (*game)[i+3] == 'e' && (*game)[i+4] == 'n' {
				numGreenCubes := GetNum(game, i-1)
				if numGreenCubes > greenCubes {
					return false
				}
			}
		}
		if c == 'b' {
			if (*game)[i+1] == 'l' && (*game)[i+2] == 'u' && (*game)[i+3] == 'e' {
				numBlueCubes := GetNum(game, i-1)
				if numBlueCubes > blueCubes {
					return false
				}
			}
		}
	}
	return true
}

func GetGameID(game *string) int {
	idx := strings.Index(*game, ":")
	var firstIdx int
	for i := idx - 1; i > -1; i-- {
		if (*game)[i] == ' ' {
			firstIdx = i + 1
			break
		}
	}
	digit := (*game)[firstIdx:idx]
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
		if GameIsPossible(&line, 12, 13, 14) {
			gameId := GetGameID(&line)
			totalSum += gameId
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", totalSum)
}

func GetAllValues(s *string, color string) []int {
	currIdx := strings.Index((*s), color)
	values := []int{}
	for currIdx > 0 && currIdx < len((*s)) {
		values = append(values, GetNum(s, currIdx-1))
		nextIdx := strings.Index((*s)[currIdx+1:], color)
		if nextIdx == -1 {
			break
		}
		currIdx += strings.Index((*s)[currIdx+1:], color) + 1
	}
	return values
}

func PartTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pwr := 1

		redValues := GetAllValues(&line, "red")
		greenValues := GetAllValues(&line, "green")
		blueValues := GetAllValues(&line, "blue")

		if len(redValues) > 0 {
			pwr *= slices.Max(redValues)
		}
		if len(greenValues) > 0 {
			pwr *= slices.Max(greenValues)
		}
		if len(blueValues) > 0 {
			pwr *= slices.Max(blueValues)
		}

		totalSum += pwr
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part Two: %d\n", totalSum)
}

func main() {
	PartOne()
	PartTwo()
}
