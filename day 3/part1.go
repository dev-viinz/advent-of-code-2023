package main

import (
	"fmt"
)

func part1(lineArray []string) {
	// sum := 0
	var realPartNumbers []int
	for key, line := range lineArray {
		currentPartNumbers := getPartNumbers(line)
		for _, partNumber := range currentPartNumbers {
			if key != 0 {
				//check if line has neighbor with previous line
				realPartNumber := getRealPartNumber(partNumber, lineArray[key-1])
				if realPartNumber != 0 {
					realPartNumbers = append(realPartNumbers, realPartNumber)
					continue
				}
			}
			// check if line has neighbors in itself
			realPartNumber := getRealPartNumber(partNumber, line)
			if realPartNumber != 0 {
				realPartNumbers = append(realPartNumbers, realPartNumber)
				continue
			}
			if key != len(lineArray)-1 {
				// check if line has neighbor with next line
				realPartNumber := getRealPartNumber(partNumber, lineArray[key+1])
				if realPartNumber != 0 {
					realPartNumbers = append(realPartNumbers, realPartNumber)
					continue
				}
			}
		}
	}
	sum := 0
	for _, num := range realPartNumbers {
		sum += num
	}
	fmt.Println("Sum of part numbers: ", sum)
}
