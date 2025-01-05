package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pageOrderingRules, updates := readInput("input")
	mappedPageRules := mappingPageOrder(pageOrderingRules)
	middlePageNoSum := 0
	correctedUpdateMiddlePageSum := 0
	for _, update := range updates {
		isValid := isUpdateCorrect(mappedPageRules, update)
		if isValid {
			middlePageNoSum += update[(len(update)-1)/2]
		} else {
			correctedUpdateMiddlePageSum += updateCorrection(mappedPageRules, update)
		}

	}
	fmt.Println("middle page no. sum for correct updates : ", middlePageNoSum)
	fmt.Println("middle page no. sum for corrected updates : ", correctedUpdateMiddlePageSum)
}

func isUpdateCorrect(mappedPageRules map[int][]int, update []int) bool {
	isValid := true
	for idx, pageNo := range update {
		beforeList := update[:idx]
		afterList := update[idx+1:]
		for _, val := range beforeList {
			isValid = isValid && doesSliceContain(mappedPageRules[val], pageNo)
		}
		for _, val := range afterList {
			isValid = isValid && doesSliceContain(mappedPageRules[pageNo], val)
		}
	}
	return isValid

}

func updateCorrection(mappedPageRules map[int][]int, update []int) int {

	correctedSlice := make([]int, 0)
	remainingSlice := make([]int, 0)
	remainingSlice = append(remainingSlice, update...)
	for i := 0; i < len(update); i++ {
		for idx, val := range remainingSlice {
			temp := make([]int, len(remainingSlice)-1)
			//making a copy because go fucks up the original value
			//at memory address if we remove an element by appendmethod,
			//so now the copy will be fucked and the original instace will be ok
			// why this happens? idk, took me 4 hours to figure this out

			remainingSliceCopy := make([]int, len(remainingSlice))
			copy(remainingSliceCopy, remainingSlice)
			copy(temp, append(remainingSliceCopy[:idx], remainingSliceCopy[idx+1:]...))
			if doesContainAll(mappedPageRules, val, temp) {
				correctedSlice = append(correctedSlice, val)
				remainingSlice = temp
				break
			}
		}
	}
	return correctedSlice[(len(correctedSlice)-1)/2]
}

func doesContainAll(mappedPageRules map[int][]int, key int, values []int) bool {
	doesContainAll := true
	for _, val := range values {
		doesContainAll = doesContainAll && doesSliceContain(mappedPageRules[key], val)
	}
	return doesContainAll

}

func doesSliceContain(slice []int, target int) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

func mappingPageOrder(pageOrderingRules [][]int) map[int][]int {
	mappedRules := make(map[int][]int, 0)
	for _, val := range pageOrderingRules {
		afterList, doesExist := mappedRules[val[0]]
		if doesExist {
			mappedRules[val[0]] = append(afterList, val[1])
		} else {
			mappedRules[val[0]] = []int{val[1]}
		}
	}
	return mappedRules
}

func readInput(filename string) ([][]int, [][]int) {
	pageOrderingRules := make([][]int, 0)
	updates := make([][]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	isOrder := true
	for scanner.Scan() {
		if isOrder {
			line := scanner.Text()
			if line == "" {
				isOrder = false
			} else {
				rule := make([]int, 2)
				nums := strings.Split(line, "|")
				rule[0], _ = strconv.Atoi(nums[0])
				rule[1], _ = strconv.Atoi(nums[1])
				pageOrderingRules = append(pageOrderingRules, rule)
			}
		} else {
			line := scanner.Text()
			numOrder := strings.Split(line, ",")
			intOrder := make([]int, 0)
			for _, val := range numOrder {
				a, _ := strconv.Atoi(val)
				intOrder = append(intOrder, a)
			}
			updates = append(updates, intOrder)
		}
	}
	return pageOrderingRules, updates
}
