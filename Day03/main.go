package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileContent := readFile("input")
	sum := 0
	do := true
	for i := 0; i < len(fileContent)-4; i++ {
		if fileContent[i:i+7] == "don't()" {
			i = i + 7
			do = false
		}
		if fileContent[i:i+4] == "do()" {
			i = i + 4
			do = true
		}
		if fileContent[i:i+4] == "mul(" {
			i = i + 4
			if !do {
				continue
			}
			int1 := ""
			int2 := ""
			for fileContent[i] >= 48 && fileContent[i] <= 57 {
				int1 += string(fileContent[i])
				i++
				if fileContent[i] == ',' {
					break
				}
			}
			if fileContent[i] != ',' {
				continue
			}
			i++
			for fileContent[i] >= 48 && fileContent[i] <= 57 {
				int2 += string(fileContent[i])
				i++
				if fileContent[i] == ')' {
					break
				}
			}
			if fileContent[i] != ')' {
				continue
			}
			if int1 == "" || int2 == "" {
				continue
			}
			num1, _ := strconv.Atoi(int1)
			num2, _ := strconv.Atoi(int2)

			fmt.Println(num1)
			fmt.Println(num2)

			sum += num1 * num2

		}
	}
	fmt.Println("sum is : ", sum)
}

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
