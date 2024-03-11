package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Races struct {
	Times     []int
	Distances []int
}

func GetRaces(filename string) *Races {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	races := Races{}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()
	strTimes := strings.Fields(text)[1:]
	for _, strTime := range strTimes {
		time, err := strconv.Atoi(strTime)
		if err != nil {
			panic(err)
		}
		races.Times = append(races.Times, time)
	}

	scanner.Scan()
	text = scanner.Text()
	strDistances := strings.Fields(text)[1:]
	for _, strDistance := range strDistances {
		distance, err := strconv.Atoi(strDistance)
		if err != nil {
			panic(err)
		}
		races.Distances = append(races.Distances, distance)
	}
	return &races
}

func part1() {
	fmt.Print("Part 1:")
	races := GetRaces("input.txt")
	var counts []int
	for i := 0; i < len(races.Times); i++ {
		raceTime := races.Times[i]
		raceDist := races.Distances[i]

		count := 0
		// brute force time
		for holdTime := 0; holdTime <= raceTime; holdTime++ {
			travelTime := raceTime - holdTime
			travelDist := travelTime * holdTime
			if travelDist > raceDist {
				count++
			}
		}
		counts = append(counts, count)
	}

	res := 1
	for _, count := range counts {
		res *= count
	}
	fmt.Println(res)
}

type Race struct {
	Time     int
	Distance int
}

func GetRaceIgnoreSpace(filename string) Race {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	race := Race{}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()
	digits := strings.Fields(text)[1:]
	allDigits := ""
	for _, d := range digits {
		allDigits += d
	}
	timeValue, err := strconv.Atoi(allDigits)
	if err != nil {
		panic(err)
	}
	race.Time = timeValue

	scanner.Scan()
	text = scanner.Text()
	digits = strings.Fields(text)[1:]
	allDigits = ""
	for _, d := range digits {
		allDigits += d
	}
	distValue, err := strconv.Atoi(allDigits)
	if err != nil {
		panic(err)
	}
	race.Distance = distValue

	return race
}

func distance(raceTime int, holdTime int) int {
	travelTime := raceTime - holdTime
	return travelTime * holdTime
}

func search(race Race, minHoldTime int, maxHoldTime int) int {
	lo := minHoldTime
	hi := maxHoldTime
	middle := (hi + lo) / 2
	if distance(race.Time, middle) > race.Distance {
		if distance(race.Time, middle-1) <= race.Distance {
			return middle
		}
		return search(race, lo, middle)
	} else {
		if distance(race.Time, middle+1) > race.Distance {
			return middle + 1
		}
		return search(race, middle+1, hi)
	}
}

func part2() {
	fmt.Print("Part 2:")
	race := GetRaceIgnoreSpace("input.txt")

	firstWinner := search(race, 0, race.Time/2)
	fmt.Println(race.Time - (firstWinner-1)*2 - 1)
}

func main() {
	fmt.Println("Day 6")
	part1()
	part2()
}
