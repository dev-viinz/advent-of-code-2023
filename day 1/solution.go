package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type digit struct {
	str string
	num int
}

var digits = []digit{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

func getScanner(data []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	return scanner
}

// this reverse function is copy pasted from stackoverflow
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func findDigit(line string, lastDigit bool) int {
	if lastDigit {
		line = reverse(line)
	}
	first := -1
	lastIndex := math.MaxInt64
	for _, digit := range digits {
		var result int
		var stringDigit string
		if lastDigit {
			stringDigit = reverse(digit.str)
		} else {
			stringDigit = digit.str
		}
		resultString := strings.Index(line, stringDigit)
		resultInt := strings.Index(line, fmt.Sprint(digit.num))

		if resultString == -1 {
			resultString = math.MaxInt64
		}
		if resultInt == -1 {
			resultInt = math.MaxInt64
		}

		if resultInt == resultString {
			continue
		}

		if resultString < resultInt {
			result = resultString
		} else {
			result = resultInt
		}
		if result != -1 && result < lastIndex {
			lastIndex = result
			first = digit.num
		}
	}
	return first
}

func main() {
	fmt.Println("Reading input file...")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	fmt.Println("Input file read successfully!")
	count := 0
	scanner := getScanner(input)

	counter := 0
	for scanner.Scan() {
		counter++
		line := scanner.Text()
		firstDigit := findDigit(line, false)
		lastDigit := findDigit(line, true)

		stringRepresentation := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		intRepresentation, err := strconv.Atoi(stringRepresentation)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}

		count += intRepresentation
	}
	fmt.Println("Answer:", count)
}
