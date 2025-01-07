package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	initialStones := readInput("input")

	initialStonesCopy := make([]int, len(initialStones))
	copy(initialStonesCopy, initialStones)

	// fmt.Println(initialStonesCopy)
	// initialStonesCopy = blink25Times(initialStonesCopy)
	// // fmt.Println(initialStonesCopy)
	// fmt.Println("Number of stones after blinking 25 times", len(initialStonesCopy))
	// initialStonesCopy = blink25Times(initialStonesCopy)
	// initialStonesCopy = blink25Times(initialStonesCopy)
	// fmt.Println("Number of stones after blinking 75 times (part 2)", len(initialStonesCopy))
	fmt.Println(stonesToHashmap(initialStones))

	stonesMap := stonesToHashmap(initialStonesCopy)
	for range 75 {
		stonesMap = blinkImproved(stonesMap)
	}
	totalStone := 0
	for _, val := range stonesMap {
		totalStone += val
	}
	fmt.Println(stonesMap)
	fmt.Println(totalStone)

}
func stonesToHashmap(stones []int) map[int]int {
	stonesHashMap := make(map[int]int, 0)
	for _, stone := range stones {
		stonesHashMap[stone] += 1
	}
	return stonesHashMap
}
func blinkImproved(stonesHashMap map[int]int) map[int]int {
	newStonesHashMap := make(map[int]int)

	for key, val := range stonesHashMap {
		numStr := strconv.Itoa(key)
		if key == 0 {
			newStonesHashMap[1] += val
		} else if len(numStr)%2 == 1 {
			newStonesHashMap[key*2024] += val
		} else {
			a, b := splitNum(key)
			newStonesHashMap[a] += val
			newStonesHashMap[b] += val
		}
	}

	return newStonesHashMap
}

func splitNum(num int) (int, int) {
	numStr := strconv.Itoa(num)
	num1Str := numStr[:len(numStr)/2]
	num2Str := numStr[len(numStr)/2:]
	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2Str)
	return num1, num2
}

func blink(stones []int) []int {
	// fmt.Println(stones)
	j := len(stones)
	for i := 0; i < j; i++ {
		numStr := strconv.Itoa(stones[i])
		if stones[i] == 0 {
			stones[i] = 1
		} else if len(numStr)%2 == 1 {
			stones[i] = stones[i] * 2024
		} else if len(numStr)%2 == 0 {
			a, b := splitNum(stones[i])
			part1 := append([]int{}, stones[:i]...)
			part2 := append([]int{}, stones[i+1:]...)
			part1 = append(part1, a)
			part1 = append(part1, b)
			stones = append(part1, part2...)
			i++
			j++
		}
	}
	return stones
}
func blink25Times(stones []int) []int {

	for i := 0; i < 25; i++ {
		fmt.Println(i)
		stones = blink(stones)
	}
	return stones
}

func readInput(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	numList := strings.Split(string(data), " ")
	intList := make([]int, 0)
	for _, val := range numList {
		n, _ := strconv.Atoi(val)
		intList = append(intList, n)
	}
	return intList
}
