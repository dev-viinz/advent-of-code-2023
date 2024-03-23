package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getScanner(data []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	return scanner
}

func getMaxAmountOfColor(line string, color string) int {
	words := strings.Split(line, " ")
	//string array of name results
	var results []int
	for key, word := range words {
		word = strings.TrimRight(word, ",")
		word = strings.TrimRight(word, ";")
		if word == color {
			r, err := strconv.Atoi(words[key-1])
			if err != nil {
				log.Fatalf("Error converting string to int: %s", err)
			}
			results = append(results, r)
		}
	}
	if len(results) == 0 {
		return 0
	}
	return slices.Max(results)
}

func getGameNumber(line string) int {
	words := strings.Split(line, " ")
	for key, word := range words {
		if word == "Game" {
			r, err := strconv.Atoi(strings.TrimRight(words[key+1], ":"))
			if err != nil {
				log.Fatalf("Error converting string to int: %s", err)
			}
			return r
		}
	}
	return 0
}

func part1(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		amountOfRed := getMaxAmountOfColor(line, "red")
		amountOfGreen := getMaxAmountOfColor(line, "green")
		amountOfBlue := getMaxAmountOfColor(line, "blue")
		if amountOfRed <= 12 && amountOfBlue <= 14 && amountOfGreen <= 13 {
			sum += getGameNumber(line)
		}
	}
	fmt.Println("Sum of games that are possible: ", sum)
}

func part2(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		amountOfRed := getMaxAmountOfColor(line, "red")
		amountOfGreen := getMaxAmountOfColor(line, "green")
		amountOfBlue := getMaxAmountOfColor(line, "blue")
		sum += (amountOfBlue * amountOfGreen * amountOfRed)
	}
	fmt.Println("Sum of minimum set of cubes: ", sum)
}

func main() {
	fmt.Println("Reading input file...")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	fmt.Println("Input file read successfully!")
	scanner1 := getScanner(input)
	scanner2 := getScanner(input)
	part1(scanner1)
	part2(scanner2)
}
