package main

import (
	"fmt"
	"sync"
)

type lineStruct struct {
	lineIndex        int
	hasBeenProcessed bool
}

func processLine(line string, lineIndices *[]lineStruct, lineIndex int) {
	hits := hitsOnLine(line)
	if hits == 0 {
		return
	}

	newIndices := make([]lineStruct, hits)
	for i := range hits {
		newIndices[i] = lineStruct{lineIndex: lineIndex + (i + 1), hasBeenProcessed: false}
	}

	*lineIndices = append(*lineIndices, newIndices...)
}

func checkIfDone(lineIndices []lineStruct) bool {
	for _, line := range lineIndices {
		if !line.hasBeenProcessed {
			return false
		}
	}
	return true
}

func part2(lineArray []string) {
	var lineIndices []lineStruct
	var mu sync.Mutex // Mutex to prevent race conditions
	for index := range lineArray {
		lineIndices = append(lineIndices, lineStruct{lineIndex: index, hasBeenProcessed: false})
	}
	var wg sync.WaitGroup
	for {
		for iteration := 0; iteration < len(lineIndices); iteration++ {
			if lineIndices[iteration].hasBeenProcessed {
				continue
			}
			wg.Add(1)
			go func(iteration int) { // Use a function literal to capture the current value of iteration
				defer wg.Done()
				mu.Lock() // Lock the mutex before calling processLine
				processLine(lineArray[lineIndices[iteration].lineIndex], &lineIndices, lineIndices[iteration].lineIndex)
				lineIndices[iteration].hasBeenProcessed = true
				mu.Unlock() // Unlock the mutex after processLine returns
			}(iteration)
		}
		wg.Wait() // Wait for all goroutines to finish before checking the condition
		if checkIfDone(lineIndices) {
			break
		}
	}
	fmt.Println("Amount of cards: ", len(lineIndices))
}

// func _part2Sync(lineArray []string) {
// 	var lineIndices []lineStruct
// 	for index := range lineArray {
// 		lineIndices = append(lineIndices, lineStruct{lineIndex: index, hasBeenProcessed: false})
// 	}
// 	for {
// 		for iteration := 0; iteration < len(lineIndices); iteration++ {
// 			if lineIndices[iteration].hasBeenProcessed {
// 				continue
// 			}
// 			processLine(lineArray[lineIndices[iteration].lineIndex], &lineIndices, lineIndices[iteration].lineIndex)
// 			lineIndices[iteration].hasBeenProcessed = true
// 		}
// 		if checkIfDone(lineIndices) {
// 			break
// 		}
// 	}
// 	fmt.Println("Amount of cards: ", len(lineIndices))
// }
