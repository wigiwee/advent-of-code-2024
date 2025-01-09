package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Green = "\033[32m"

const lobbyX int = 101
const lobbyY int = 103

// var lobbyArea [lobbyY][lobbyX]int
var robotsPos [][2][2]int

// var robotsPosAfter100sec = make([][2]int, 0)

func main() {

	readInput("input")
	robotsPosAfter100sec := simulateNsec(100)
	q1, q2, q3, q4 := getRobotsInQ(robotsPosAfter100sec)
	fmt.Println(q1 * q2 * q3 * q4)

	minSafteyFactor := 999999999999999999
	minSafteyFactorTime := 0
	for t := range lobbyX * lobbyY {
		q1, q2, q3, q4 := getRobotsInQ(simulateNsec(t))
		safteyFactor := q1 * q2 * q3 * q4
		if safteyFactor < minSafteyFactor {
			minSafteyFactor = safteyFactor
			minSafteyFactorTime = t
		}

	}
	fmt.Println(minSafteyFactorTime, minSafteyFactor)
	printRobotsPosition(simulateNsec(minSafteyFactorTime))

}

func printRobotsPosition(robotsPosition [][2]int) {
	var bathroomFloor [103][101]int
	for x := range 101 {
		for y := range 103 {
			bathroomFloor[y][x] = 0
		}
	}
	for _, val := range robotsPosition {
		bathroomFloor[val[1]][val[0]] += 1
	}
	// for _, val := range bathroomFloor {
	// 	fmt.Println(val)
	// }
	for _, Y := range bathroomFloor {
		line := ""
		for _, X := range Y {
			if X == 0 {
				line += "  "
			} else {
				line += Green + " 0"
			}
		}
		fmt.Println(line)
	}
}
func getRobotsInQ(robotsPosition [][2]int) (int, int, int, int) {
	xMid := int(lobbyX / 2)
	yMid := int(lobbyY / 2)
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, val := range robotsPosition {
		if val[0] < xMid {
			if val[1] < yMid {
				q1++
			} else if val[1] > yMid {
				q3++
			}
		} else if val[0] > xMid {
			if val[1] < yMid {
				q2++
			} else if val[1] > yMid {
				q4++
			}
		}
	}
	return q1, q2, q3, q4
}
func simulateNsec(n int) [][2]int {
	robotsPosition := make([][2]int, 0)
	for _, val := range robotsPos {
		x, y := val[0][0]+n*val[1][0], val[0][1]+n*val[1][1]
		x = x % lobbyX
		y = y % lobbyY
		if x < 0 {
			x = lobbyX + x
		}
		if y < 0 {
			y = lobbyY + y
		}
		robotsPosition = append(robotsPosition, [2]int{x, y})
	}
	return robotsPosition
}

func readInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var robot [2][2]int
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		position := strings.Split(splitLine[0][2:], ",")
		x, _ := strconv.Atoi(position[0])
		y, _ := strconv.Atoi(position[1])
		robot[0][0] = x
		robot[0][1] = y
		velocity := strings.Split(splitLine[1][2:], ",")
		x, _ = strconv.Atoi(velocity[0])
		y, _ = strconv.Atoi(velocity[1])
		robot[1][0] = x
		robot[1][1] = y
		robotsPos = append(robotsPos, robot)
	}

}
