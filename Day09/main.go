package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	diskmap := readInput("input")
	diskmapCopy := make([]int, len(diskmap))
	copy(diskmapCopy, diskmap)
	fmt.Println("Checksum of the optimized disk storage is :", calculateChecksum(diskmapCopy))
	fmt.Println("Checksum of the optimized disk storage(part 2 ) is :", calculateChecksumP2(diskmap))

}
func diskReprArr(diskmap []int) [][]int {
	diskmapRepr := make([][]int, 0)

	for i := 0; i < len(diskmap); i++ {
		file := make([]int, diskmap[i])
		for j := range diskmap[i] {
			file[j] = getFileid(i)
		}
		diskmapRepr = append(diskmapRepr, file)
	}
	return diskmapRepr
}
func calculateChecksumP2(diskmap []int) int {
	diskmapRepr := diskReprArr(diskmap)
	ic := 0
	if (len(diskmap)-1)%2 == 0 {
		ic = len(diskmap) - 1
	} else {

		ic = len(diskmap) - 2
	}
	// i -> file
	// j -> free space
	for i := ic; i > 0; i -= 2 {
		for j := 1; j < i; j += 2 {
			jlen := 0
			for idx := range len(diskmapRepr[j]) {
				if diskmapRepr[j][idx] == -1 {
					jlen++
				}
			}
			if jlen >= len(diskmapRepr[i]) {
				if diskmapRepr[j][0] == -1 {
					copy(diskmapRepr[j], diskmapRepr[i])
				} else {
					for idx := range diskmapRepr[j] {
						if diskmapRepr[j][idx] != -1 {
							continue
						} else {
							for z := range len(diskmapRepr[i]) {
								diskmapRepr[j][idx+z] = diskmapRepr[i][0]
							}
							break

						}
					}
				}
				for idx := range diskmapRepr[i] {
					diskmapRepr[i][idx] = -1
				}
				break
			}

		}
	}
	index := 0
	checksum := 0
	for _, fileblock := range diskmapRepr {
		for _, file := range fileblock {
			if file != -1 {
				checksum = checksum + (file * index)
			}
			index++
		}
	}
	return checksum
}

func readInput(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	diskmap := make([]int, 0)

	stringData := strings.Trim(string(data), " ")
	for _, digit := range stringData {
		digit, _ := strconv.Atoi(string(digit))
		diskmap = append(diskmap, digit)
	}
	return diskmap

}
func getFileid(originalIdx int) int {
	if originalIdx%2 == 0 {
		return originalIdx / 2
	}
	return -1
}
func calculateChecksum(diskmap []int) int {
	checksum := 0
	i := 0
	j := 0
	if (len(diskmap)-1)%2 == 0 {
		j = len(diskmap) - 1
	} else {

		j = len(diskmap) - 2
	}
	index := 0
	for diskmap[i] != 0 {
		for diskmap[i] != 0 {
			checksum += getFileid(i) * index
			index++
			diskmap[i]--
		}
		if i == j {
			break
		}
		i++
		for diskmap[i] != 0 {
			checksum += getFileid(j) * index
			diskmap[i]--
			index++
			diskmap[j]--
			if diskmap[j] == 0 {
				j = j - 2
			}
		}
		i++

	}
	return checksum

}
