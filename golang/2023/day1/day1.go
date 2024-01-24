package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, runeValue := range scanner.Text() {
			if runeValue >= 0 && runeValue <= '9' {
				fmt.Printf("%#U", runeValue)
				fmt.Printf("\n")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
