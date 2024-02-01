package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GameIsPossible(game string, redCubes int, greenCubes int, blueCubes int) bool {
	for i, c := range game {
		if c == 'r' {
			if game[i+1] == 'e' && game[i+2] == 'd' {
				var lastIdx int
				for j := i - 2; j > -1; j-- {
					if game[j] == ' ' {
						lastIdx = j + 1
						break
					}
				}
				digit := game[lastIdx : i-1]
				numRedCubes, err := strconv.Atoi(digit)
				if err != nil {
					panic(err)
				}

				if numRedCubes > redCubes {
					return false
				}
			}
		}
		if c == 'g' {
			if game[i+1] == 'r' && game[i+2] == 'e' && game[i+3] == 'e' && game[i+4] == 'n' {
				var lastIdx int
				for j := i - 2; j > -1; j-- {
					if game[j] == ' ' {
						lastIdx = j + 1
						break
					}
				}
				digit := game[lastIdx : i-1]
				numGreenCubes, err := strconv.Atoi(digit)
				if err != nil {
					panic(err)
				}

				if numGreenCubes > greenCubes {
					return false
				}
			}
		}
		if c == 'b' {
			if game[i+1] == 'l' && game[i+2] == 'u' && game[i+3] == 'e' {
				var lastIdx int
				for j := i - 2; j > -1; j-- {
					if game[j] == ' ' {
						lastIdx = j + 1
						break
					}
				}
				digit := game[lastIdx : i-1]
				numBlueCubes, err := strconv.Atoi(digit)
				if err != nil {
					panic(err)
				}

				if numBlueCubes > blueCubes {
					return false
				}
			}
		}
	}
	return true
}


func GetGameID(game string) int {
	idx := strings.Index(game, ":")
	var firstIdx int
	for i := idx - 1; i > -1; i-- {
		if game[i] == ' ' {
			firstIdx = i + 1
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
	counter := 0
	for scanner.Scan() {
		counter++
		line := scanner.Text()
		if GameIsPossible(line, 12, 13, 14) {
			gameId := GetGameID(line)
			totalSum += gameId
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
