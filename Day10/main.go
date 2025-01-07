package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var uphillMap [][]int = readInput("input")
var visited [][]int = make([][]int, 0)

func main() {
	fmt.Println("Sum of score of trails is :", getAllTrailScores())
	fmt.Println("Sum of score of trails (part 2) is :", getAllTrailScoresP2())
}
func isVisited(x int, y int) bool {
	for _, val := range visited {
		if val[0] == x && val[1] == y {
			return true
		}
	}
	return false
}
func getTrailScore(x int, y int, currLoc int) int {

	if currLoc == 9 {
		if isVisited(x, y) {
			return 0
		}
		visited = append(visited, []int{x, y})
		return 1
	}
	trailScore := 0
	if x-1 >= 0 {
		if uphillMap[y][x-1] == currLoc+1 {
			trailScore += getTrailScore(x-1, y, currLoc+1)
		}
	}
	if x+1 < len(uphillMap[y]) {
		if uphillMap[y][x+1] == currLoc+1 {
			trailScore += getTrailScore(x+1, y, currLoc+1)
		}
	}
	if y-1 >= 0 {
		if uphillMap[y-1][x] == currLoc+1 {
			trailScore += getTrailScore(x, y-1, currLoc+1)
		}
	}
	if y+1 < len(uphillMap) {
		if uphillMap[y+1][x] == currLoc+1 {
			trailScore += getTrailScore(x, y+1, currLoc+1)
		}
	}

	return trailScore
}

func getTrailScoreP2(x int, y int, currLoc int) int {

	if currLoc == 9 {
		return 1
	}
	trailScore := 0
	if x-1 >= 0 {
		if uphillMap[y][x-1] == currLoc+1 {
			trailScore += getTrailScoreP2(x-1, y, currLoc+1)
		}
	}
	if x+1 < len(uphillMap[y]) {
		if uphillMap[y][x+1] == currLoc+1 {
			trailScore += getTrailScoreP2(x+1, y, currLoc+1)
		}
	}
	if y-1 >= 0 {
		if uphillMap[y-1][x] == currLoc+1 {
			trailScore += getTrailScoreP2(x, y-1, currLoc+1)
		}
	}
	if y+1 < len(uphillMap) {
		if uphillMap[y+1][x] == currLoc+1 {
			trailScore += getTrailScoreP2(x, y+1, currLoc+1)
		}
	}

	return trailScore
}

func getAllTrailScoresP2() int {
	totalScore := 0
	for y, Y := range uphillMap {
		for x, X := range Y {

			if X == 0 {
				totalScore += getTrailScoreP2(x, y, 0)
			}
		}
	}
	return totalScore
}

func getAllTrailScores() int {
	totalScore := 0
	for y, Y := range uphillMap {
		for x, X := range Y {

			if X == 0 {
				for len(visited) != 0 {
					visited = visited[:len(visited)-1]
				}
				totalScore += getTrailScore(x, y, 0)
			}
		}
	}
	return totalScore
}

func readInput(filename string) [][]int {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	input := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		intLine := make([]int, 0)
		for idx := range len(line) {
			num, _ := strconv.Atoi(string(line[idx]))
			intLine = append(intLine, num)
		}
		input = append(input, intLine)
	}
	return input
}
