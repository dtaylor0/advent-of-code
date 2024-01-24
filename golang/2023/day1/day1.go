package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
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

func runeIsInt(in rune) bool {
	return in >= '0' && in <= '9'
}

func findFirstInt(s *string) int {
	for idx, runeVal := range *s {
		if runeIsInt(runeVal) {
			return idx
		}
	}
	return -1
}

func findLastInt(s *string) int {
	i := -1
	for idx, runeVal := range *s {
		if runeIsInt(runeVal) {
			i = idx
		}
	}
	return i
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalSum := 0
	// digitMap := map[string]int{
	// 	"zero":  0,
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// 	"four":  4,
	// 	"five":  5,
	// 	"six":   6,
	// 	"seven": 7,
	// 	"eight": 8,
	// 	"nine":  9,
	// }

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		first, _ := utf8.DecodeRuneInString(text)
		var firstValue int
		last, _ := utf8.DecodeLastRuneInString(text)
		var lastValue int

		if runeIsInt(first) {
			firstValue = int(first - '0')
		} else {
			print("else")
		}
		if runeIsInt(last) {
			lastValue = int(last - '0')
		} else {
			print("else")
		}
		totalSum += 10*firstValue + lastValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}

func main() {
	partOne()
	inputs := []string{"djfkdlsjf4jkfldsjdf", "3jfkd", "kdjsfl8", "jdkfls"}
	for i := 0; i < len(inputs); i++ {
		fmt.Printf("%s: %d\n", inputs[i], findFirstInt(&inputs[i]))
	}
}
