package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hitsOnLine(line string) int {
	split := strings.Split(line, "|")
	winningNumbers := strings.Split(strings.TrimSpace(split[0]), " ")
	yourNumbers := strings.Split(strings.TrimSpace(split[1]), " ")

	wNumbersMap := make(map[int]bool)
	for _, number := range winningNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		wNumbersMap[num] = true
	}

	hits := 0
	for _, number := range yourNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		if wNumbersMap[num] {
			hits++
		}
	}

	return hits
}

func _hitsOnLine(line string) int {
	split := strings.Split(line, "|")
	winningNumbers := strings.Split(strings.TrimSpace(split[0]), " ")
	yourNumbers := strings.Split(strings.TrimSpace(split[1]), " ")
	var wNumbersArray []int
	var yNumbersArray []int
	for _, number := range winningNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		wNumbersArray = append(wNumbersArray, num)
	}
	for _, number := range yourNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		yNumbersArray = append(yNumbersArray, num)
	}
	hits := 0
	for _, wNumber := range wNumbersArray {
		for _, yNumber := range yNumbersArray {
			if wNumber == yNumber {
				hits++
			}
		}
	}
	return hits
}

func pointsOfLine(line string) int {
	hits := hitsOnLine(line)
	if hits > 0 {
		points := 1
		for range hits - 1 {
			points = points * 2
		}
		return points
	}
	return 0
}

func part1(lineArray []string) {
	points := 0
	for _, line := range lineArray {
		points += pointsOfLine(line)
	}
	fmt.Println("Points: ", points)
}
