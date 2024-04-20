package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getScanner(data []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	return scanner
}

func numbersInString(line string) []int {
	s := strings.Split(line, " ")
	var numbers []int
	for _, r := range s {
		n, err := strconv.Atoi(string(r))
		if err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
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

	var seeds []int
	var seedToSoil []SeedToSoil
	var soilToFertilizer []SoilToFertilizer
	var fertilizerToWater []FertilizerToWater
	var waterToLight []WaterToLight
	var lightToTemperature []LightToTemperature
	var temperatureToHumidity []TemperatureToHumidity
	var humidityToLocation []HumidityToLocation

	currentMap := ""
	for _, line := range lineArray {
		//consolidate the if satements into a switch statement
		switch {
		case strings.Contains(line, "seeds"):
			seeds = numbersInString(line)
		case strings.Contains(line, "seed-to-soil"):
			currentMap = "seed-to-soil"
		case strings.Contains(line, "soil-to-fertilizer"):
			currentMap = "soil-to-fertilizer"
		case strings.Contains(line, "fertilizer-to-water"):
			currentMap = "fertilizer-to-water"
		case strings.Contains(line, "water-to-light"):
			currentMap = "water-to-light"
		case strings.Contains(line, "light-to-temperature"):
			currentMap = "light-to-temperature"
		case strings.Contains(line, "temperature-to-humidity"):
			currentMap = "temperature-to-humidity"
		case strings.Contains(line, "humidity-to-location"):
			currentMap = "humidity-to-location"
		}
		switch currentMap {
		case "seed-to-soil":

		}
	}
}
