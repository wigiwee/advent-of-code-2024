package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reacterLevels := readReacterLevels("input")

	safeReports := 0
	for level := range reacterLevels {
		if isReportSafe(reacterLevels[level]) {
			safeReports++
		} else if applyProblemDampener(reacterLevels[level]) {
			safeReports++
		}
	}
	fmt.Println("No. of Safe reports : ", safeReports)

}

func applyProblemDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])
		if isReportSafe(newReport) {
			return true
		}
	}
	return false

}

func isReportSafe(report []int) bool {
	if report[0] > report[1] {
		for i := 1; i < len(report); i++ {
			if !(report[i-1]-report[i] <= 3 && report[i-1]-report[i] >= 1) {
				return false
			}
		}
	} else if report[0] < report[1] {
		for i := 1; i < len(report); i++ {
			if !(report[i]-report[i-1] <= 3 && report[i]-report[i-1] >= 1) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func readReacterLevels(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var reacterLevels [][]int = make([][]int, 0)
	for scanner.Scan() {
		levelsInString := strings.Split(scanner.Text(), " ")
		level := make([]int, 0)
		for idx := range levelsInString {
			reading, _ := strconv.Atoi(levelsInString[idx])
			level = append(level, reading)
		}
		reacterLevels = append(reacterLevels, level)
	}
	return reacterLevels
}
