package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type assembler struct {
	A, B, C uint
	ip      int
	program []uint
	output  []uint
	isHalt  bool
}

func main() {
	assembler := readInput("input")
	fmt.Println("input: ", assembler.program)
	assembler.executeProgram()
	fmt.Println("ouput: ", assembler.output)
	// fmt.Println(assembler.getOutputStr())	//solution to part 1

	// assembler.analyse()
	Areq := 0
	for i := 35184372088832; i > 0; i++ {
		fmt.Println(i)
		assembler.ip = 0
		assembler.A = uint(i)
		assembler.output = []uint{}
		assembler.isHalt = false
		assembler.executeProgram()
		// if assembler.output[len(assembler.output)-1] != assembler.program[len(assembler.program)-1] {
		// 	i = i + int(math.Pow(8, 14))
		// }
		if compareSlices(assembler.output, assembler.program) {
			Areq = i
		}
	}
	fmt.Println(Areq)
	fmt.Println(assembler.program)
	fmt.Println(assembler.output)

}

func (a *assembler) analyse() {
	prevLen := 0
	for i := 0; i >= 0; i = i + 1 {
		a.A = uint(i)
		a.ip = 0
		a.isHalt = false
		a.output = []uint{}
		a.executeProgram()
		if len(a.output) > prevLen {
			fmt.Println(len(a.output), i)
			prevLen = len(a.output)
		}
		// if a.output[len(a.output)-1] == a.program[len(a.program)-1] {
		// 	fmt.Println(i)
		// }
	}
}
func (a *assembler) executeProgram() {
	for !a.isHalt {
		a.performNextInstruction()
	}
}

func (a *assembler) performNextInstruction() {

	performInstruction := a.getInstructionMethod()
	performInstruction()
	a.ip += 2
	if a.ip >= len(a.program) {
		a.isHalt = true
	}

}

func (a *assembler) getInstructionMethod() func() {

	switch a.program[a.ip] {

	case 0:
		return func() { a.A = uint(a.A / uint(math.Pow(float64(2), float64(a.getComboOperand())))) }
	case 1:
		return func() { a.B = a.B ^ a.program[a.ip+1] }
	case 2:
		return func() { a.B = a.getComboOperand() % 8 }
	case 3:
		if a.A != 0 {
			return func() { a.ip = int(a.program[a.ip+1]) - 2 }
		} else {
			return func() {}
		}
	case 4:
		return func() { a.B = a.B ^ a.C }
	case 5:
		return func() { a.output = append(a.output, (a.getComboOperand())%8) }
	case 6:
		return func() { a.B = uint(a.A / uint(math.Pow(float64(2), float64(a.getComboOperand())))) }
	case 7:
		return func() { a.C = uint(a.A / uint(math.Pow(float64(2), float64(a.getComboOperand())))) }
	}
	return nil
}

func (a *assembler) getComboOperand() uint {
	if a.program[a.ip+1] >= 0 && a.program[a.ip+1] <= 3 {
		return a.program[a.ip+1]
	} else if a.program[a.ip+1] == 4 {
		return a.A
	} else if a.program[a.ip+1] == 5 {
		return a.B
	} else if a.program[a.ip+1] == 6 {
		return a.C
	} else {
		log.Fatal("comboOperand: ", a.program[a.ip+1])
	}
	return 0
}

func (a *assembler) getOutputStr() string {
	outputStr := ""
	for _, val := range a.output {
		valStr := strconv.Itoa(int(val))
		outputStr += valStr + ","
	}
	return outputStr[:len(outputStr)-1]
}

func compareSlices(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range len(a) {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func readInput(filename string) *assembler {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	lineArr := strings.Split(scanner.Text(), " ")
	regA, err := strconv.Atoi(lineArr[len(lineArr)-1])
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	lineArr = strings.Split(scanner.Text(), " ")
	regB, err := strconv.Atoi(lineArr[len(lineArr)-1])
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	lineArr = strings.Split(scanner.Text(), " ")
	regC, err := strconv.Atoi(lineArr[len(lineArr)-1])
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	scanner.Scan()

	programStr := strings.Split(strings.Split(scanner.Text(), " ")[1], ",")
	program := make([]uint, 0)
	for _, val := range programStr {
		valNum, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, uint(valNum))
	}
	return &assembler{
		A:       uint(regA),
		B:       uint(regB),
		C:       uint(regC),
		program: program,
		ip:      0,
		output:  []uint{},
		isHalt:  false,
	}
}
