package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input := readInput("input")

	totalTokens := 0
	for _, a := range input {
		A := mat.NewDense(2, 2, []float64{a[0][0], a[0][1], a[1][0], a[1][1]})
		b := mat.NewVecDense(2, []float64{a[0][2], a[1][2]})
		var x mat.VecDense
		if err := x.SolveVec(A, b); err != nil {
			fmt.Println(err)
		}

		combination := x.RawVector().Data
		if !isInteger(combination[0]) || combination[0] < 0 || !isInteger(combination[1]) || combination[1] < 0 {
			continue
		}

		totalTokens += convertToInt(combination[0])*3 + convertToInt(combination[1])
	}
	fmt.Println("Tokens required to win as many prices as possible ", totalTokens)

	totalTokensP2 := 0
	for _, a := range input {
		A2 := mat.NewDense(2, 2, []float64{a[0][0], a[0][1], a[1][0], a[1][1]})
		b2 := mat.NewVecDense(2, []float64{a[0][2] + 10000000000000, a[1][2] + 10000000000000})

		var p2 mat.VecDense
		if err := p2.SolveVec(A2, b2); err != nil {
			fmt.Println(err)
		}
		combination := p2.RawVector().Data
		if !isInteger(combination[0]) || combination[0] < 0 || !isInteger(combination[1]) || combination[1] < 0 {
			continue
		}
		totalTokensP2 += convertToInt(combination[0])*3 + convertToInt(combination[1])
	}
	fmt.Println(totalTokensP2)
}

func convertToInt(n float64) int {
	if n-float64(int(n)) <= 0.0005 {
		return int(n)
	} else if n-float64(int(n)) >= 0.9999 {
		return int(n) + 1
	}
	return 0
}
func isInteger(n float64) bool {
	if n-float64(int(n)) <= 0.0005 {
		return true
	} else if n-float64(int(n)) >= 0.9999 {
		return true
	}
	return false
}

func readInput(filename string) [][2][3]float64 {
	input := make([][2][3]float64, 0)
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		line1Arr := strings.Split(line, " ")
		Ax, _ := strconv.Atoi(line1Arr[2][2 : len(line1Arr[2])-1])
		Ay, _ := strconv.Atoi(line1Arr[3][2:])

		scanner.Scan()
		line = scanner.Text()
		line2Arr := strings.Split(line, " ")
		Bx, _ := strconv.Atoi(line2Arr[2][2 : len(line2Arr[2])-1])
		By, _ := strconv.Atoi(line2Arr[3][2:])

		scanner.Scan()
		line = scanner.Text()
		line3Arr := strings.Split(line, " ")
		Px, _ := strconv.Atoi(line3Arr[1][2 : len(line3Arr[1])-1])
		Py, _ := strconv.Atoi(line3Arr[2][2:])

		scanner.Scan()

		a := [2][3]float64{{float64(Ax), float64(Bx), float64(Px)}, {float64(Ay), float64(By), float64(Py)}}
		input = append(input, a)
	}

	return input

}
