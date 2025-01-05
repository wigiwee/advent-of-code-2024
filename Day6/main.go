package main

import (
	"bufio"
	"fmt"
	"os"
)

var mappedArea []string = readInput("input")
var visitedPos [][]int = [][]int{}

func main() {
	mapArea := make([]string, len(mappedArea))
	copy(mapArea, mappedArea)
	patrol(mapArea)
	fmt.Println("unique positions visited : ", countDistinctPos(mapArea))

	//remove initial position from visitedPos
	x, y := getGuardPosition(mappedArea)

	for idx, val := range visitedPos {
		if val[0] == x && val[1] == y {
			visitedPos = append(visitedPos[:idx], visitedPos[idx+1:]...)
		}
	}
	didLoop := 0
	for _, val := range visitedPos {
		mapCopy := make([]string, len(mappedArea))
		copy(mapCopy, mappedArea)
		addObstacles(mapCopy, val[0], val[1])
		if patrolP2(mapCopy) {
			didLoop++
		}
	}
	fmt.Println("Obstacle could be placed on", didLoop, "locations")

}

func addObstacles(mapArea []string, x int, y int) {
	mapArea[y] = mapArea[y][:x] + "#" + mapArea[y][x+1:]
}

func getGuardPosition(mapArea []string) (int, int) {
	for y, Y := range mapArea {
		for x, X := range Y {
			if X == '^' {
				return x, y
			}
		}
	}
	return -1, -1

}

func markAsVisited(mapArea []string, x int, y int) {
	mapArea[y] = mapArea[y][:x] + "X" + mapArea[y][x+1:]
}

func changeDir(mapArea []string, x int, y int, direction string) {
	mapArea[y] = mapArea[y][:x] + direction + mapArea[y][x+1:]
}

func countDistinctPos(mapArea []string) int {
	count := 0
	for y, Y := range mapArea {
		for x, X := range Y {
			if X == 'v' || X == '>' || X == '<' || X == '^' {
				markAsVisited(mapArea, x, y)
				count++
				visitedPos = append(visitedPos, []int{x, y})
			}
			if X == 'X' {
				visitedPos = append(visitedPos, []int{x, y})
				count++
			}
		}
	}
	return count
}

type turn struct {
	x         int
	y         int
	direction string
}

func isTurnAlreadyTake(turnsTaken []turn, t turn) bool {
	for _, val := range turnsTaken {
		if val == t {
			return true
		}
	}
	return false
}
func patrolP2(mapArea []string) bool {

	x, y := getGuardPosition(mapArea)
	turnsTaken := []turn{}
	for x < len(mapArea[y])-1 && y < len(mapArea)-1 && x > 0 && y > 0 {
		if string(mapArea[y][x]) == "^" {
			if string(mapArea[y-1][x]) == "#" {
				if isTurnAlreadyTake(turnsTaken, turn{x, y, "^"}) {
					return true
				}
				turnsTaken = append(turnsTaken, turn{x, y, "^"})
				changeDir(mapArea, x, y, ">")
			} else {
				markAsVisited(mapArea, x, y)
				y--
				changeDir(mapArea, x, y, "^")
			}

		} else if string(mapArea[y][x]) == ">" {
			if string(mapArea[y][x+1]) == "#" {
				if isTurnAlreadyTake(turnsTaken, turn{x, y, ">"}) {
					return true
				}
				turnsTaken = append(turnsTaken, turn{x, y, ">"})
				changeDir(mapArea, x, y, "v")
			} else {
				markAsVisited(mapArea, x, y)
				x++
				changeDir(mapArea, x, y, ">")
			}
		} else if string(mapArea[y][x]) == "<" {
			if string(mapArea[y][x-1]) == "#" {
				if isTurnAlreadyTake(turnsTaken, turn{x, y, "<"}) {
					return true
				}
				turnsTaken = append(turnsTaken, turn{x, y, "<"})
				changeDir(mapArea, x, y, "^")
			} else {
				markAsVisited(mapArea, x, y)
				x--
				changeDir(mapArea, x, y, "<")

			}
		} else if string(mapArea[y][x]) == "v" {
			if string(mapArea[y+1][x]) == "#" {
				if isTurnAlreadyTake(turnsTaken, turn{x, y, "v"}) {
					return true
				}
				turnsTaken = append(turnsTaken, turn{x, y, "v"})
				changeDir(mapArea, x, y, "<")
			} else {
				markAsVisited(mapArea, x, y)
				y++
				changeDir(mapArea, x, y, "v")
			}
		}
	}
	return false
}

func patrol(mapArea []string) {
	x, y := getGuardPosition(mapArea)

	for x < len(mapArea[y])-1 && y < len(mapArea)-1 && x > 0 && y > 0 {
		if string(mapArea[y][x]) == "^" {
			if string(mapArea[y-1][x]) == "#" {
				changeDir(mapArea, x, y, ">")
			} else {
				markAsVisited(mapArea, x, y)
				y--
				changeDir(mapArea, x, y, "^")
			}
		} else if string(mapArea[y][x]) == ">" {
			if string(mapArea[y][x+1]) == "#" {
				changeDir(mapArea, x, y, "v")
			} else {
				markAsVisited(mapArea, x, y)
				x++
				changeDir(mapArea, x, y, ">")
			}
		} else if string(mapArea[y][x]) == "<" {
			if string(mapArea[y][x-1]) == "#" {
				changeDir(mapArea, x, y, "^")
			} else {
				markAsVisited(mapArea, x, y)
				x--
				changeDir(mapArea, x, y, "<")

			}
		} else if string(mapArea[y][x]) == "v" {
			if string(mapArea[y+1][x]) == "#" {
				changeDir(mapArea, x, y, "<")
			} else {
				markAsVisited(mapArea, x, y)
				y++
				changeDir(mapArea, x, y, "v")
			}
		}
	}

}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
