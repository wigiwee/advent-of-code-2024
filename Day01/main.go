package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//getting input
	file, err := os.Open("input")
	if err != nil {
		panic("Error reading the input file")
	}
	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		abc := strings.Split(scanner.Text(), "   ")
		a, _ := strconv.Atoi(strings.TrimSpace(abc[0]))
		b, _ := strconv.Atoi(strings.TrimSpace(abc[1]))
		list1 = append(list1, a)
		list2 = append(list2, b)
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			totalDistance += list1[i] - list2[i]
		} else {
			totalDistance += list2[i] - list1[i]
		}
	}

	fmt.Println("Total distance : ", totalDistance)

	similarityScore := 0
	j := 0
	for i := 0; i < len(list1); i++ {
		for list1[i] > list2[j] {
			if j < len(list2)-1 {
				j++
			} else {
				break
			}
		}
		for list1[i] == list2[j] {
			similarityScore += list1[i]
			if j < len(list2)-1 {
				j++
			} else {
				break
			}
		}
	}
	fmt.Println("Similarity score : ", similarityScore)
}
