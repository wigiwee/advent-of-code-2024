package main

import (
	"bufio"
	"fmt"
	"os"
)

var cityMap []string = readInput("input")
var antennas map[string][][]int = getAntennasMap()
var antiNodes [][]int = getAllAntinodes()
var xMax int = len(cityMap[0])
var yMax int = len(cityMap)

func main() {
	// for _, val := range antennas {
	// 	fmt.Println(val)
	// }
	fmt.Println(len(antiNodes))
	an := getAllAntinodesP2()
	fmt.Println(len(an))
	// for _, a := range an {
	// 	fmt.Println(a)
	// }

}

func getAntennasMap() map[string][][]int {
	antennaMap := make(map[string][][]int)
	for y, Y := range cityMap {
		for x := range Y {
			if cityMap[y][x] != '.' {
				antennaMap[string(cityMap[y][x])] = append(antennaMap[string(cityMap[y][x])], []int{x, y})
			}
		}
	}
	return antennaMap
}

func isPresent(slice [][]int, target []int) bool {
	for _, node := range slice {
		if node[0] == target[0] && node[1] == target[1] {
			return true
		}
	}
	return false
}

func getAllAntinodes() [][]int {
	keys := make([]string, 0, len(antennas))
	for k := range antennas {
		keys = append(keys, k)
	}
	allAntiNodes := make([][]int, 0)
	for _, key := range keys {
		allAntiNodes = append(allAntiNodes, getAntinodes(key)...)
	}

	uniqueAntiNodes := make([][]int, 0)
	uniqueAntiNodes = append(uniqueAntiNodes, allAntiNodes[0])

	for _, node := range allAntiNodes {
		if !isPresent(uniqueAntiNodes, node) {
			uniqueAntiNodes = append(uniqueAntiNodes, node)
		}
	}
	return uniqueAntiNodes
}
func getAllAntinodesP2() [][]int {
	keys := make([]string, 0, len(antennas))
	for k := range antennas {
		keys = append(keys, k)
	}
	allAntiNodes := make([][]int, 0)
	for _, key := range keys {
		allAntiNodes = append(allAntiNodes, getAntinodesP2(key)...)
	}

	uniqueAntiNodes := make([][]int, 0)
	uniqueAntiNodes = append(uniqueAntiNodes, allAntiNodes[0])

	for _, node := range allAntiNodes {
		if !isPresent(uniqueAntiNodes, node) {
			uniqueAntiNodes = append(uniqueAntiNodes, node)
		}
	}
	return uniqueAntiNodes
}

func getAntinodes(freq string) [][]int {
	antinodes := make([][]int, 0)
	antennasPos := antennas[freq]

	for idx, Apos := range antennasPos {
		for i := idx + 1; i < len(antennasPos); i++ {
			Bpos := antennasPos[i]

			nodeA := []int{Apos[0] + (Apos[0] - Bpos[0]), Apos[1] + (Apos[1] - Bpos[1])}
			nodeB := []int{Bpos[0] + (Bpos[0] - Apos[0]), Bpos[1] + (Bpos[1] - Apos[1])}
			if nodeA[0] >= 0 && nodeA[0] < xMax && nodeA[1] >= 0 && nodeA[1] < yMax {
				antinodes = append(antinodes, nodeA)
			}
			if nodeB[0] >= 0 && nodeB[0] < xMax && nodeB[1] >= 0 && nodeB[1] < yMax {
				antinodes = append(antinodes, nodeB)
			}
		}
	}
	return antinodes
}
func getAntinodesP2(freq string) [][]int {
	antinodes := make([][]int, 0)
	antennasPos := antennas[freq]

	for idx, Apos := range antennasPos {
		for i := idx + 1; i < len(antennasPos); i++ {

			Bpos := make([]int, 2)
			copy(Bpos, antennasPos[i])

			// adding antennas intianally
			Aposcopy := make([]int, 2)
			Bposcopy := make([]int, 2)
			copy(Aposcopy, Apos)
			copy(Bposcopy, Bpos)
			antinodes = append(antinodes, Aposcopy)
			antinodes = append(antinodes, Bposcopy)

			nodeA := []int{Apos[0] + (Apos[0] - Bpos[0]), Apos[1] + (Apos[1] - Bpos[1])}
			nodeB := []int{Bpos[0] + (Bpos[0] - Apos[0]), Bpos[1] + (Bpos[1] - Apos[1])}

			if nodeA[0] >= 0 && nodeA[0] < xMax && nodeA[1] >= 0 && nodeA[1] < yMax {
				nodeAcopy := make([]int, 2)
				copy(nodeAcopy, nodeA)
				antinodes = append(antinodes, nodeAcopy)
				nodeA[0] += Apos[0] - Bpos[0]
				nodeA[1] += Apos[1] - Bpos[1]
			}
			if nodeB[0] >= 0 && nodeB[0] < xMax && nodeB[1] >= 0 && nodeB[1] < yMax {
				nodeBcopy := make([]int, 2)
				copy(nodeBcopy, nodeB)
				antinodes = append(antinodes, nodeBcopy)
				nodeB[0] += Bpos[0] - Apos[0]
				nodeB[1] += Bpos[1] - Apos[1]
			}
			for nodeA[0] >= 0 && nodeA[0] < xMax && nodeA[1] >= 0 && nodeA[1] < yMax {
				copyNodeA := make([]int, 2)
				copy(copyNodeA, nodeA)
				antinodes = append(antinodes, copyNodeA)
				nodeA[0] += Apos[0] - Bpos[0]
				nodeA[1] += Apos[1] - Bpos[1]
			}
			for nodeB[0] >= 0 && nodeB[0] < xMax && nodeB[1] >= 0 && nodeB[1] < yMax {
				copyNodeB := make([]int, 2)
				copy(copyNodeB, nodeB)
				antinodes = append(antinodes, copyNodeB)
				nodeB[0] += Bpos[0] - Apos[0]
				nodeB[1] += Bpos[1] - Apos[1]
			}
		}
	}
	return antinodes
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
