package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

type PartNumber struct {
	number        int
	startPosition int
	endPosition   int
}

func getScanner(data []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	return scanner
}

func getPartNumbers(line string) []PartNumber {
	// get numbers then return them and their position
	var numbers []PartNumber
	var lastAddition int = -99
	var currentNumber int
	for index := 0; index < len(line); index++ {
		num, err := strconv.Atoi(string(line[index]))
		if err != nil {
			continue
		}
		if lastAddition == index-1 && numbers[currentNumber].endPosition == index-1 {
			partNumber, err := strconv.Atoi(fmt.Sprint(numbers[currentNumber].number) + fmt.Sprint(num))
			if err != nil {
				log.Fatalf("Error converting string to int: %s", err)
			}
			numbers[currentNumber].number = partNumber
			numbers[currentNumber].endPosition = index
			lastAddition = index
			continue
		}
		numbers = append(numbers, PartNumber{number: num, startPosition: index, endPosition: index})
		lastAddition = index
		currentNumber = len(numbers) - 1
	}
	return numbers
}

func isSymbol(char rune) bool {
	if string(char) != "." && !unicode.IsDigit(char) {
		return true
	}
	return false
}

func getRealPartNumber(partNumber PartNumber, lineToCheck string) int {
	start := partNumber.startPosition
	if (partNumber.startPosition - 1) > 0 {
		start = partNumber.startPosition - 1
	}
	end := partNumber.endPosition
	if (partNumber.endPosition + 1) < len(lineToCheck)-1 {
		end = partNumber.endPosition + 1
	}
	for index := start; index <= end; index++ {
		potentialSymbol := rune(lineToCheck[index])
		isItTho := isSymbol(potentialSymbol)
		if isItTho {
			return partNumber.number
		}
	}
	return 0
}
