package main

import (
	"bufio"
	"fmt"
	"os"
)

// dir -> [0,1,2,3] ==> [^, >, v, <]
const UP, RIGHT, DOWN, LEFT = 0, 1, 2, 3

var Ex, Ey int

type path struct {
	destX int
	destY int
	cost  int
	dir   int
}

func main() {
	maze := readInput("input")
	Sx, Sy, Ex, Ey := getStartEndPos(maze)
	mazeMap := getMazeMap(maze)
	bestPathScore, lastUpdatedDistanceMap := search([2]int{Sx, Sy}, [2]int{Ex, Ey}, mazeMap)
	fmt.Println(bestPathScore)

	fmt.Println(lastUpdatedDistanceMap)
	fmt.Println(getTilesFromShortestPath([2]int{Sx, Sy}, [2]int{Ex, Ey}, lastUpdatedDistanceMap, mazeMap))
	// for key, val := range lastUpdatedMap {
	// 	lastUpdatedMap[key] = removeDuplicate(val)
	// }
	// // for key, val := range lastUpdatedMap {
	// // 	fmt.Printf("%v %v \n", key, val)
	// // }
	// // _, lastUPdate := search([2]int{Sx, Sy}, [2]int{Ex, Ey}, mazeMap)
	// fmt.Println(lastUpdatedMap)
	// fmt.Println(getTilesFromShortestPath([2]int{Sx, Sy}, [2]int{Ex, Ey}, lastUpdatedMap, mazeMap))
}

func getTilesFromShortestPath(start, end [2]int, lastUpdatedMap map[[2]int][][2]int, mazeMap map[[2]int][]path) int {
	visited := make(map[[2]int]bool)
	stack := make([][2]int, 0)
	totalTiles := 0
	stack = append(stack, end)
	lastUpdatedMap[start] = [][2]int{}

	for len(stack) != 0 {
		a := stack[0]
		stack = append([][2]int{}, stack[1:]...)
		shortestPathfrom := lastUpdatedMap[a]
		for _, next := range shortestPathfrom {

			totalTiles += getPath(a, next, mazeMap).cost
			if !visited[next] {
				stack = append(stack, next)
				visited[next] = true
			}

		}
	}
	fmt.Println(len(visited))
	diff := 0
	for key, _ := range visited {
		fmt.Println(key)
		if len(lastUpdatedMap[key]) > 1 {
			diff++
		}
	}
	return totalTiles - diff + 1
}
func removeDuplicate(slice [][2]int) [][2]int {
	output := make([][2]int, 0)
	output = append(output, slice[0])
	for _, val := range slice {
		a := true
		for _, out := range output {
			if val == out {
				a = false
				break
			}
		}
		if a {
			output = append(output, val)
		}
	}
	return output
}
func getPath(start, end [2]int, mazeMap map[[2]int][]path) path {
	a := mazeMap[start]
	for _, path := range a {
		if path.destX == end[0] && path.destY == end[1] {
			return path
		}
	}
	return path{}
}

func search(S [2]int, E [2]int, mazeMap map[[2]int][]path) (int, map[[2]int][][2]int) {
	lastUpdated := make(map[[2]int][][2]int)

	cost := make(map[[2]int]int)

	//reachable -> {Sx, Sy, Ex, Ey, totalCost, eDir}
	reachable := make([][6]int, 0)
	currentDir := make(map[[2]int]int)

	reachable = append(reachable, [6]int{S[0], S[1], S[0], S[1], 0, RIGHT})
	currentDir[S] = RIGHT

	for len(reachable) != 0 {
		//get nearest neighbour from the reachable slice
		lowestCost := 999999
		lowestCostIdx := 0
		for i, val := range reachable {
			if lowestCost > val[4] {
				lowestCost = val[4]
				lowestCostIdx = i
			}
		}
		nearest := reachable[lowestCostIdx]
		reachable = append(reachable[:lowestCostIdx], reachable[lowestCostIdx+1:]...)
		start := [2]int{nearest[0], nearest[1]}
		end := [2]int{nearest[2], nearest[3]}

		endCost, doesExist := cost[end]

		if doesExist {
			if endCost > cost[start]+nearest[4] {
				cost[end] = cost[start] + nearest[4]
				lastUpdated[end] = append([][2]int{}, start)
			} else if endCost == cost[start]+nearest[4]-1000 || endCost == cost[start]+nearest[4]+1000 {
				lastUpdated[end] = append(lastUpdated[end], start)
				continue
			} else {
				continue

			}
		} else {
			cost[end] = cost[start] + nearest[4]
			lastUpdated[end] = append(lastUpdated[end], start)
		}

		currentDir[end] = nearest[5]

		for _, neighbour := range mazeMap[end] {
			totalCost := 0
			startDir := currentDir[end]
			endDir := neighbour.dir
			totalCost = neighbour.cost
			if startDir == endDir {
				totalCost += 0
			} else if startDir-endDir == 2 || endDir-startDir == 2 {
				continue
			} else {
				totalCost += 1000
			}
			reachable = append(reachable, [6]int{end[0], end[1], neighbour.destX, neighbour.destY, totalCost, neighbour.dir})
		}
	}
	return cost[E], lastUpdated
}

func getMazeMap(maze [][]rune) map[[2]int][]path {

	mazeMap := make(map[[2]int][]path)

	for Y := range maze {
		for X := range maze[Y] {
			if maze[Y][X] == '#' {
				continue
			}
			if !isNode(X, Y, maze) {
				continue
			}
			//check up
			for y := 1; y < len(maze); y++ {
				if maze[Y-y][X] == '#' {
					break
				} else if isNode(X, Y-y, maze) {
					mazeMap[[2]int{X, Y}] = append(mazeMap[[2]int{X, Y}], path{destX: X, destY: Y - y, cost: y, dir: UP})
					break
				}
			}
			//check down
			for y := 1; y < len(maze); y++ {
				if maze[Y+y][X] == '#' {
					break
				} else if isNode(X, Y+y, maze) {
					mazeMap[[2]int{X, Y}] = append(mazeMap[[2]int{X, Y}], path{destX: X, destY: Y + y, cost: y, dir: DOWN})
					break
				}
			}
			//check left
			for x := 1; x < len(maze[0]); x++ {
				if maze[Y][X-x] == '#' {
					break
				} else if isNode(X-x, Y, maze) {
					mazeMap[[2]int{X, Y}] = append(mazeMap[[2]int{X, Y}], path{destX: X - x, destY: Y, cost: x, dir: LEFT})
					break
				}
			}
			//check right
			for x := 1; x < len(maze[0]); x++ {
				if maze[Y][X+x] == '#' {
					break
				} else if isNode(X+x, Y, maze) {

					mazeMap[[2]int{X, Y}] = append(mazeMap[[2]int{X, Y}], path{destX: X + x, destY: Y, cost: x, dir: RIGHT})
					break
				}
			}
		}
	}
	return mazeMap
}

func isNode(x, y int, maze [][]rune) bool {
	if maze[y][x] == '#' {
		return false
	} else if maze[y][x] == 'S' || maze[y][x] == 'E' {
		return true
	}
	up, down, left, right := false, false, false, false
	if y-1 >= 0 {
		if maze[y-1][x] == '.' {
			up = true
		}
	}
	if y+1 < len(maze) {
		if maze[y+1][x] == '.' {
			down = true
		}
	}
	if x-1 >= 0 {
		if maze[y][x-1] == '.' {
			right = true
		}
	}
	if x+1 < len(maze[0]) {
		if maze[y][x+1] == '.' {
			left = true
		}
	}

	return (up && (right || left)) || (down && (right || left))
}

func getStartEndPos(maze [][]rune) (int, int, int, int) {
	Sx, Sy, Ex, Ey := 0, 0, 0, 0
	for y, Y := range maze {
		for x, X := range Y {
			if X == 'S' {
				Sx, Sy = x, y
			}
			if X == 'E' {
				Ex, Ey = x, y
			}
		}
	}
	return Sx, Sy, Ex, Ey
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	maze := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, []rune(line))
	}
	return maze
}
