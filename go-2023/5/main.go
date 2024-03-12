package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func nextMap(scanner *bufio.Scanner) [][3]int {
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
		nextMap := nextMap(scanner)
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

type Seed struct {
	Start  int
	Length int
}

type Range struct {
	Start  int
	Length int
}

type Conversion struct {
	Dest   int
	Src    int
	Length int
}

type ConvertMap struct {
	From        string
	To          string
	Conversions []Conversion
}

func startsWithLetter(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] >= byte('A') && s[0] <= byte('Z') {
		return true
	}
	if s[0] >= byte('a') && s[0] <= byte('z') {
		return true
	}
	return false
}

func nextConvertMap(scanner *bufio.Scanner) *ConvertMap {
	res := &ConvertMap{}
	text := ""
	for scanner.Scan() {
		text := scanner.Text()
		if startsWithLetter(text) {
			title := strings.Fields(text)[0]
			parsedTitle := strings.Split(title, "-")
			res.From, res.To = parsedTitle[0], parsedTitle[2]
			break
		}
	}
	scanner.Scan()
	text = scanner.Text()
	for len(text) > 0 {
		fields := strings.Fields(text)
		dest, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		src, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(fields[2])
		if err != nil {
			panic(err)
		}
		row := Conversion{dest, src, length}
		res.Conversions = append(res.Conversions, row)

		if !scanner.Scan() {
			break
		}
		text = scanner.Text()
	}
	if len(res.Conversions) == 0 {
		return nil
	}
	return res
}

func parseSeeds(line string) []Seed {
	var res []Seed
	fields := strings.Fields(line)[1:]
	for i := 0; i < len(fields); i += 2 {
		start, err := strconv.Atoi(fields[i])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(fields[i+1])
		if err != nil {
			panic(err)
		}
		res = append(res, Seed{start, length})
	}
	return res
}

func (s Seed) StartsIn(r Range) bool {
	return s.Start >= r.Start && s.Start < r.Start+r.Length
}

func (s Seed) EndsIn(r Range) bool {
	return s.Start+s.Length-1 >= r.Start && s.Start+s.Length <= r.Start+r.Length
}

func (s Seed) Surrounds(r Range) bool {
	return s.Start < r.Start && s.Start+s.Length > r.Start+r.Length
}

func cut(seed Seed, with Range) (removed Seed, leftover []Seed) {
	if seed.StartsIn(with) && seed.EndsIn(with) {
		removed = seed
	} else if seed.StartsIn(with) && !seed.EndsIn(with) {
		removed = Seed{seed.Start, with.Start + with.Length - seed.Start}
		leftover = []Seed{{with.Start + with.Length, seed.Start + seed.Length - with.Start - with.Length}}
	} else if seed.EndsIn(with) {
		removed = Seed{with.Start, seed.Start + seed.Length - with.Start}
		leftover = []Seed{{seed.Start, seed.Length - removed.Length}}
	} else if seed.Surrounds(with) {
		removed = Seed{with.Start, with.Length}
		leftover = []Seed{
            {seed.Start, with.Start - seed.Start},
            {with.Start + with.Length, seed.Start + seed.Length - with.Start - with.Length},
        }
	} else {
		leftover = []Seed{seed}
	}
	return
}

func part2() {
	fmt.Print("Part 2: ")
	fmt.Println()
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	seeds := parseSeeds(text)
	var nextSeeds []Seed

	nextMap := nextConvertMap(scanner)
	for nextMap != nil {
		fmt.Println(nextMap.From, nextMap.To)
		for _, conv := range nextMap.Conversions {
			seedCount := len(seeds)
			for i := 0; i < seedCount; i++ {
				seed := seeds[i]
				removed, leftover := cut(seed, Range{conv.Src, conv.Length})
				if removed.Length > 0 {
					offset := removed.Start - conv.Src
					newSeed := Seed{conv.Dest + offset, removed.Length}
					nextSeeds = append(nextSeeds, newSeed)
				}
				validLeftover := []Seed{}
				for _, l := range leftover {
					if l.Length > 0 {
						validLeftover = append(validLeftover, l)
					}
				}
				if len(validLeftover) == 0 {
					seeds = slices.Concat(seeds[:i], seeds[i+1:])
					seedCount--
					i--
				} else if len(validLeftover) == 1 {
					seeds[i] = validLeftover[0]
				} else if len(validLeftover) == 2 {
					seeds = slices.Concat(seeds[:i], validLeftover, seeds[i+1:])
					seedCount++
					i++
				}
			}
		}
		seeds = slices.Concat(seeds, nextSeeds)
		nextSeeds = []Seed{}
		nextMap = nextConvertMap(scanner)
	}

	minSeed := seeds[0]
	for _, s := range seeds {
		if s.Start < minSeed.Start {
			minSeed = s
		}
	}

	fmt.Println(minSeed.Start)
}

func main() {
	fmt.Println("Day 5")
	part1()
	part2()
}
