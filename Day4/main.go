package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordSearch := readFile("input")
	XMAScount := 0

	/*
		M . S
		. A .
		M . S
	*/
	MAScount := 0
	for y := range wordSearch {
		for x := range wordSearch[y] {
			if string(wordSearch[y][x]) == "X" {
				if checkDown(wordSearch, x, y) {
					XMAScount++
				}
				if checkUp(wordSearch, x, y) {
					XMAScount++
				}
				if checkleft(wordSearch, x, y) {
					XMAScount++
				}
				if checkRight(wordSearch, x, y) {
					XMAScount++
				}
				if checkUpLeft(wordSearch, x, y) {
					XMAScount++
				}
				if checkUpRight(wordSearch, x, y) {
					XMAScount++
				}
				if checkDownLeft(wordSearch, x, y) {
					XMAScount++
				}
				if checkDownRight(wordSearch, x, y) {
					XMAScount++
				}
			} else if string(wordSearch[y][x]) == "A" {
				if checkForMas(wordSearch, x, y) {
					MAScount++
				}
			}
		}
	}
	fmt.Println("XMAS count ", XMAScount)
	fmt.Println("MAS count ", MAScount)

}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

/*
M . S
. A .
M . S
*/
func checkForMas(letters []string, x int, y int) bool {
	if x-1 >= 0 && x+1 < len(letters[y]) && y-1 >= 0 && y+1 < len(letters) {
		firstDiagonal := false
		secondDiagonal := false
		firstDiagonal = string(letters[y-1][x-1]) == "M" && string(letters[y+1][x+1]) == "S" || string(letters[y-1][x-1]) == "S" && string(letters[y+1][x+1]) == "M"
		secondDiagonal = string(letters[y-1][x+1]) == "M" && string(letters[y+1][x-1]) == "S" || string(letters[y-1][x+1]) == "S" && string(letters[y+1][x-1]) == "M"
		return firstDiagonal && secondDiagonal
	}
	return false
}
func checkleft(letters []string, x int, y int) bool {
	if x-3 >= 0 {
		leftLetters := letters[y][(x - 3):x]
		if leftLetters == "SAM" {
			return true
		}
	}
	return false
}

func checkRight(letters []string, x int, y int) bool {
	if x+3 < len(letters[y]) {
		rightLetters := letters[y][x+1 : x+4]
		if rightLetters == "MAS" {
			return true
		}
	}
	return false

}
func checkUp(letters []string, x int, y int) bool {
	upLetters := ""
	if y-3 >= 0 {
		upLetters += string(letters[y-1][x])
		upLetters += string(letters[y-2][x])
		upLetters += string(letters[y-3][x])
		if upLetters == "MAS" {
			return true
		}
	}
	return false
}
func checkDown(letters []string, x int, y int) bool {
	downLetters := ""
	if y+3 < len(letters) {
		downLetters += string(letters[y+1][x])
		downLetters += string(letters[y+2][x])
		downLetters += string(letters[y+3][x])
		if downLetters == "MAS" {
			return true
		}
	}
	return false

}
func checkUpLeft(letters []string, x int, y int) bool {
	upLeftLetters := ""
	if x-3 >= 0 && y-3 >= 0 {
		upLeftLetters += string(letters[y-1][x-1])
		upLeftLetters += string(letters[y-2][x-2])
		upLeftLetters += string(letters[y-3][x-3])
		if upLeftLetters == "MAS" {
			return true
		}
	}
	return false
}
func checkUpRight(letters []string, x int, y int) bool {
	upRightLetters := ""
	if x+3 < len(letters[y]) && y-3 >= 0 {
		upRightLetters += string(letters[y-1][x+1])
		upRightLetters += string(letters[y-2][x+2])
		upRightLetters += string(letters[y-3][x+3])
		if upRightLetters == "MAS" {
			return true
		}
	}
	return false
}
func checkDownLeft(letters []string, x int, y int) bool {
	downLeftLetters := ""
	if x-3 >= 0 && y+3 < len(letters) {
		downLeftLetters += string(letters[y+1][x-1])
		downLeftLetters += string(letters[y+2][x-2])
		downLeftLetters += string(letters[y+3][x-3])
		if downLeftLetters == "MAS" {
			return true
		}
	}
	return false
}
func checkDownRight(letters []string, x int, y int) bool {
	downRightLetters := ""
	if x+3 < len(letters[y]) && y+3 < len(letters) {
		downRightLetters += string(letters[y+1][x+1])
		downRightLetters += string(letters[y+2][x+2])
		downRightLetters += string(letters[y+3][x+3])
		if downRightLetters == "MAS" {
			return true
		}
	}
	return false
}
