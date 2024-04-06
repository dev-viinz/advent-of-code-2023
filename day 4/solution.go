package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func getScanner(data []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	return scanner
}

func main() {
	fmt.Println("Reading input file...")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	fmt.Println("Input file read successfully!")
	scanner := getScanner(input)
	var lineArray []string
	for scanner.Scan() {
		line := scanner.Text()
		// remove game number
		line = strings.Split(line, ":")[1]
		// remove leading and trailing whitespaces
		line = strings.TrimSpace(line)
		lineArray = append(lineArray, line)
	}
	part1(lineArray)
	part2(lineArray)
}
