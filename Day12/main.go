package main

import (
	"bufio"
	"fmt"
	"os"
)

var visited map[[2]int]bool = make(map[[2]int]bool)
var xMax int
var yMax int
var garden [][]rune

func main() {
	garden = readInput("input")
	xMax = len(garden[0]) - 1
	yMax = len(garden) - 1
	priceMapP1, priceMapP2 := visitAll(garden)

	totalPrice := 0
	for _, val := range priceMapP1 {
		totalPrice += val
	}
	fmt.Println("Total price of fence is :", totalPrice)
	totalPrice = 0
	for _, val := range priceMapP2 {
		totalPrice += val
	}
	fmt.Println("Total price of fence (part2 ) is  :", totalPrice)

}
func isEqual(x int, y int, region rune) bool {
	if x == -1 {
		return false
	}
	if x == xMax+1 {
		return false
	}
	if y == -1 {
		return false
	}
	if y == xMax+1 {
		return false
	}

	return garden[y][x] == region
}

func getCorners(x int, y int) int {
	corners := 0
	region := garden[y][x]
	//first corner ie. top left corner
	if !isEqual(x-1, y, region) && !isEqual(x-1, y-1, region) && !isEqual(x, y-1, region) {
		corners++
	} else if isEqual(x-1, y, region) && isEqual(x, y-1, region) && !isEqual(x-1, y-1, region) {
		corners++
	} else if !isEqual(x-1, y, region) && !isEqual(x, y-1, region) && isEqual(x-1, y-1, region) {
		corners++
	}

	//second corner ie. top right corner
	if !isEqual(x+1, y, region) && !isEqual(x+1, y-1, region) && !isEqual(x, y-1, region) {
		corners++
	} else if isEqual(x+1, y, region) && isEqual(x, y-1, region) && !isEqual(x+1, y-1, region) {
		corners++
	} else if !isEqual(x+1, y, region) && !isEqual(x, y-1, region) && isEqual(x+1, y-1, region) {
		corners++
	}

	//third corner ie. bottom right corner
	if !isEqual(x+1, y, region) && !isEqual(x+1, y+1, region) && !isEqual(x, y+1, region) {
		corners++
	} else if isEqual(x+1, y, region) && isEqual(x, y+1, region) && !isEqual(x+1, y+1, region) {
		corners++
	} else if !isEqual(x+1, y, region) && !isEqual(x, y+1, region) && isEqual(x+1, y+1, region) {
		corners++
	}

	//forth corner ie. bottom left corner
	if !isEqual(x-1, y, region) && !isEqual(x-1, y+1, region) && !isEqual(x, y+1, region) {
		corners++
	} else if isEqual(x-1, y, region) && isEqual(x, y+1, region) && !isEqual(x-1, y+1, region) {
		corners++
	} else if !isEqual(x-1, y, region) && !isEqual(x, y+1, region) && isEqual(x-1, y+1, region) {
		corners++
	}
	return corners
}
func getAreaAndPeremter(x int, y int, region rune) (int, int, int) {
	area := 0
	peremeter := 0
	visited[[2]int{x, y}] = true
	area++
	corner := getCorners(x, y)
	if x == 0 || x == xMax {
		peremeter++
	}
	if y == 0 || y == yMax {
		peremeter++
	}
	if x-1 >= 0 {
		if garden[y][x-1] == region && !visited[[2]int{x - 1, y}] {
			a, p, c := getAreaAndPeremter(x-1, y, region)
			area += a
			peremeter += p
			corner += c
		} else if garden[y][x-1] != region {
			peremeter++
		}
	}
	if x+1 <= xMax {
		if garden[y][x+1] == region && !visited[[2]int{x + 1, y}] {
			a, p, c := getAreaAndPeremter(x+1, y, region)
			area += a
			peremeter += p
			corner += c
		} else if garden[y][x+1] != region {
			peremeter++
		}
	}
	if y-1 >= 0 {
		if garden[y-1][x] == region && !visited[[2]int{x, y - 1}] {
			a, p, c := getAreaAndPeremter(x, y-1, region)
			area += a
			peremeter += p
			corner += c
		} else if garden[y-1][x] != region {
			peremeter++
		}
	}
	if y+1 <= yMax {
		if garden[y+1][x] == region && !visited[[2]int{x, y + 1}] {
			a, p, c := getAreaAndPeremter(x, y+1, region)
			area += a
			peremeter += p
			corner += c
		} else if garden[y+1][x] != region {
			peremeter++
		}
	}
	return area, peremeter, corner

}

func visitAll(garden [][]rune) (map[rune]int, map[rune]int) {
	priceMapP1 := make(map[rune]int)
	priceMapP2 := make(map[rune]int)
	for y, Y := range garden {
		for x, X := range Y {
			if !visited[[2]int{x, y}] {
				area, peremeter, corners := getAreaAndPeremter(x, y, X)
				priceMapP1[X] += area * peremeter
				priceMapP2[X] += area * corners
			}
		}
	}
	return priceMapP1, priceMapP2
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	output := make([][]rune, 0)
	for scanner.Scan() {
		lineString := scanner.Text()
		output = append(output, []rune(lineString))
	}
	return output
}
