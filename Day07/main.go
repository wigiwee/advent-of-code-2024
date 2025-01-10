package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var equations [][]int = readInput("input")

func main() {

	sumOfTotal := 0
	equationsCopy := make([][]int, len(equations))
	afterConcat := 0
	copy(equationsCopy, equations)
	for _, eq := range equationsCopy {
		if calculate(append([]int{}, eq[2:]...), eq[1], eq[0]) {
			sumOfTotal += eq[0]
		} else if calculateP2(append([]int{}, eq[2:]...), eq[1], eq[0]) {
			afterConcat += eq[0]
		}
	}
	fmt.Println("part 1 : Sum of Total :", sumOfTotal)
	fmt.Println("part 2 : Sum of Total :", sumOfTotal+afterConcat)
}

// part 1
func calculate(eq []int, outcome int, req int) bool {
	if len(eq) == 1 {
		if outcome*eq[0] == req {
			return true
		} else if outcome+eq[0] == req {
			return true
		} else {
			return false
		}
	}
	if calculate(append([]int{}, eq[1:]...), outcome*eq[0], req) {
		return true
	} else if calculate(append([]int{}, eq[1:]...), outcome+eq[0], req) {
		return true
	} else {
		return false
	}
}

// part 2
func calculateP2(eq []int, outcome int, req int) bool {
	if len(eq) == 1 {
		if outcome*eq[0] == req {
			return true
		} else if outcome+eq[0] == req {
			return true
		} else if concat(outcome, eq[0]) == req {
			return true
		} else {
			return false
		}
	}
	if calculateP2(append([]int{}, eq[1:]...), outcome*eq[0], req) {
		return true
	} else if calculateP2(append([]int{}, eq[1:]...), outcome+eq[0], req) {
		return true
	} else if calculateP2(append([]int{}, eq[1:]...), concat(outcome, eq[0]), req) {
		return true
	} else {
		return false
	}
}

func concat(a int, b int) int {
	astr := strconv.Itoa(a)
	bstr := strconv.Itoa(b)
	result, _ := strconv.Atoi(astr + bstr)
	return result

}

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	input := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		eq := make([]int, 0)
		nums := strings.Split(scanner.Text(), " ")
		firstNum, _ := strconv.Atoi(nums[0][:(len(nums[0]) - 1)])
		eq = append(eq, firstNum)
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			eq = append(eq, num)
		}
		input = append(input, eq)
	}
	return input
}
