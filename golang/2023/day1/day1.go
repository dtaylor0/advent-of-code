package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first rune
		var last rune
		for _, runeValue := range scanner.Text() {
			if runeValue >= '0' && runeValue <= '9' {
				if first != 0 {
					last = runeValue
				} else {
					first = runeValue
				}
			}
		}
		if last == 0 {
			last = first
		}
		totalSum += 10*int(first-'0') + int(last-'0')
		first = 0
		last = 0
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}

func stringToInt(s string) int {
	digitsMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	return digitsMap[s]
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		firstIdx := -1
		lastIdx := -1
		var firstVal string
		var lastVal string

		for _, digit := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			currIdxFirst := strings.Index(text, digit)
			if currIdxFirst > -1 && currIdxFirst < firstIdx {
				firstIdx = currIdxFirst
				firstVal = string(digit)
			}
			currIdxLast := strings.LastIndex(text, digit)
			if currIdxLast > -1 && currIdxLast > lastIdx {
				lastIdx = currIdxLast
				lastVal = string(digit)
			}
		}

		for _, digit := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
			currIdxFirst := strings.Index(text, digit)
			if currIdxFirst > -1 && currIdxFirst < firstIdx {
				firstIdx = currIdxFirst
				firstVal = string(digit)
			}
			currIdxLast := strings.LastIndex(text, digit)
			if currIdxLast > -1 && currIdxLast > lastIdx {
				lastIdx = currIdxLast
				lastVal = string(digit)
			}
		}

		firstValue := stringToInt(firstVal)
		lastValue := stringToInt(lastVal)

		totalSum += 10*firstValue + lastValue

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total Sum Part 2: %d\n", totalSum)
}

func main() {
	partOne()
	partTwo()
}
