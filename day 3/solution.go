package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

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
		lineArray = append(lineArray, scanner.Text())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go part1(&wg, lineArray)
	go part2(&wg, lineArray)
	wg.Wait()
}
