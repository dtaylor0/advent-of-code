package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getLines(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type Node struct {
	left  string
	right string
}

func createNetwork(lines []string) *map[string]Node {
	res := make(map[string]Node)
	for _, line := range lines {
		fields := strings.Fields(line)
		value := fields[0]
		left := fields[2][1:4]
		right := fields[3][:3]

		res[value] = Node{left, right}
	}
	return &res
}

func part1() {
	lines := getLines("input.txt")
	dirs := lines[0]
	lines = lines[2:]
	network := *createNetwork(lines)

	steps := 0
	curr := "AAA"

	i := 0
	for true {
		dir := dirs[i]
		if dir == 'L' {
			curr = network[curr].left
		} else if dir == 'R' {
			curr = network[curr].right
		}
		steps++

		if curr == "ZZZ" {
			break
		}

		i++
		if i == len(dirs) {
			i = 0
		}
	}

	fmt.Println("Part 1: ", steps)
}

type ZTracker struct {
	zCount int
	steps  int
	zDist  int
}

func getLCM(arr []int) int {
	multiplesAgg := make(map[int]int)

	for _, n := range arr {
		multiples := make(map[int]int)
		d := n - 1
		for d > 1 {
			if n%d == 0 {
				multiples[d]++
				n = n / d
			}
			d--
		}
		for k, v := range multiples {
			if v > multiplesAgg[k] {
				multiplesAgg[k] = v
			}
		}
	}
	lcm := 1
	for k, v := range multiplesAgg {
		lcm *= k * v
	}
	return lcm
}

func part2() {
	lines := getLines("input.txt")
	dirs := lines[0]
	lines = lines[2:]
	network := *createNetwork(lines)

	steps := 0
	var locs []string
	for k := range network {
		if k[2] == 'A' {
			locs = append(locs, k)
		}
	}

	tracker := make(map[int]*ZTracker)
	for i := range locs {
		tracker[i] = &ZTracker{}
	}

	i := 0
	for true {
		dir := dirs[i]
		steps++

		for j := 0; j < len(locs); j++ {
			if dir == 'L' {
				locs[j] = network[locs[j]].left
			} else if dir == 'R' {
				locs[j] = network[locs[j]].right
			}

			if locs[j][2] == 'Z' {
				tracker[j].zCount++
				if tracker[j].zCount%2 == 1 {
					tracker[j].steps = steps
				} else if tracker[j].zCount%2 == 0 {
					tracker[j].zDist = steps - tracker[j].steps
				}
			}
		}

		done := true
		for _, v := range tracker {
			if v.zDist == 0 {
				done = false
			}
		}
		if done {
			break
		}

		i++
		if i == len(dirs) {
			i = 0
		}
	}

	lcm := []int{}
	for _, v := range tracker {
		lcm = append(lcm, v.zDist)
	}

	fmt.Println("Part 2: ", getLCM(lcm))
}

func main() {
	fmt.Println("Day 8")
	part1()
	part2()
}
