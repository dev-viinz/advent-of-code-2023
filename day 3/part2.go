package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"sync"
	"unicode"
)

type Vector2 struct {
	x int
	y int
}

type Validgear struct {
	gear Vector2
	num1 Vector2
	num2 Vector2
}

type Grid2D [][]rune

func initGrid(lineArray []string) Grid2D {
	grid := make(Grid2D, len(lineArray))
	for index, line := range lineArray {
		for runePos, currentRune := range line {
			if grid[index] == nil {
				grid[index] = make([]rune, len(line))
			}
			grid[index][runePos] = currentRune
		}
	}
	return grid
}

func findGears(grid Grid2D) []Vector2 {
	var gears []Vector2
	for y, row := range grid {
		for x, cell := range row {
			if cell == '*' {
				gears = append(gears, Vector2{x: x, y: y})
			}
		}
	}
	return gears
}

func getAdjacent(arr Grid2D, coords Vector2) []Vector2 {
	n := len(arr)
	m := len(arr[0])

	var v []Vector2

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				if coords.x+dx >= 0 && coords.x+dx < n && coords.y+dy >= 0 && coords.y+dy < m {
					v = append(v, Vector2{x: coords.x + dx, y: coords.y + dy})
				}
			}
		}
	}

	return v
}

func determineValidgearsWithNumberPositions(gears []Vector2, grid Grid2D) []Validgear {
	var validGears []Validgear
	for _, gear := range gears {
		adjacent := getAdjacent(grid, gear)
		var ratioNumbers []Vector2
		var burntCoords []Vector2
	breakInner:
		for _, adj := range adjacent {
			if unicode.IsDigit(grid[adj.y][adj.x]) {
				for _, burnt := range burntCoords {
					if burnt.y == adj.y {
						if burnt.x+1 == adj.x || burnt.x-1 == adj.x {
							continue breakInner
						}
					}
				}
				burntCoords = append(burntCoords, buildNumberFromPosAsCoords(grid, adj)...)
				ratioNumbers = append(ratioNumbers, adj)
				if len(ratioNumbers) == 2 {
					validGears = append(validGears, Validgear{gear: gear, num1: ratioNumbers[0], num2: ratioNumbers[1]})
					break
				}
			}
		}
	}
	return validGears
}

func buildNumberFromPosAsCoords(grid Grid2D, pos Vector2) []Vector2 {
	var coords []Vector2
	for i := pos.x; i >= 0; i-- {
		if unicode.IsDigit(grid[pos.y][i]) {
			coords = append(coords, Vector2{x: i, y: pos.y})
		} else {
			break
		}
	}
	for i := pos.x + 1; i < len(grid[pos.y]); i++ {
		if unicode.IsDigit(grid[pos.y][i]) {
			coords = append(coords, Vector2{x: i, y: pos.y})
		} else {
			break
		}
	}
	return coords
}

func buildNumberFromPos(grid Grid2D, pos Vector2) int {
	var number string
	coords := buildNumberFromPosAsCoords(grid, pos)
	sort.Slice(coords, func(i, j int) bool {
		return coords[i].x < coords[j].x
	})
	for _, coord := range coords {
		number += string(grid[coord.y][coord.x])
	}
	// for i := pos.x; i >= 0; i-- {
	// 	if unicode.IsDigit(grid[pos.y][i]) {
	// 		number = string(grid[pos.y][i]) + number
	// 	} else {
	// 		break
	// 	}
	// }
	// for i := pos.x + 1; i < len(grid[pos.y]); i++ {
	// 	if unicode.IsDigit(grid[pos.y][i]) {
	// 		number = number + string(grid[pos.y][i])
	// 	} else {
	// 		break
	// 	}
	// }
	result, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("Error converting string to int: %s", err)
	}
	return result
}

func getGearRatio(gear Validgear, grid Grid2D) int {
	number1 := buildNumberFromPos(grid, gear.num1)
	number2 := buildNumberFromPos(grid, gear.num2)
	return number1 * number2
}

func part2(wg *sync.WaitGroup, lineArray []string) {
	defer wg.Done()
	grid := initGrid(lineArray)
	gears := findGears(grid)
	validGears := determineValidgearsWithNumberPositions(gears, grid)
	sumOfRatios := 0
	for _, gear := range validGears {
		ratio := getGearRatio(gear, grid)
		sumOfRatios += ratio
	}
	fmt.Println("Sum of gear ratios: ", sumOfRatios)
}
