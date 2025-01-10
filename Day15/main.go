package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	warehouse, movements := readInput("input")
	warehouseCopy := make([][]rune, len(warehouse))
	for i := range warehouse {
		warehouseCopy[i] = make([]rune, len(warehouse[i]))
		copy(warehouseCopy[i], warehouse[i])
	}
	warehouseCopy = robotMakesMoves(warehouseCopy, movements)
	fmt.Println("sum of all boxes' GPS coordinates", calculateSumOfGPSCoords(warehouseCopy, 'O'))

	warehouse2 := getWarehouse2Map(warehouse)
	warehouse2 = robotMakesMovesInWarehouse2(warehouse2, movements)
	fmt.Println("sum of all boxes' GPS coordinates (part2)", calculateSumOfGPSCoords(warehouse2, '['))

}

func getRobotPos(warehouse [][]rune) (int, int) {
	for y, Y := range warehouse {
		for x, X := range Y {
			if X == '@' {
				return x, y
			}
		}
	}
	return -1, -1
}

func readInput(filename string) ([][]rune, []rune) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	warehouse := make([][]rune, 0)

	for scanner.Scan() {
		lineStr := scanner.Text()
		if lineStr == "" {
			break
		}
		line := make([]rune, 0)
		line = append(line, []rune(lineStr)...)
		warehouse = append(warehouse, line)
	}
	movements := make([]rune, 0)
	for scanner.Scan() {
		lineStr := scanner.Text()
		movements = append(movements, []rune(lineStr)...)
	}
	return warehouse, movements

}
func printWarehouse(warehouse [][]rune) {
	for _, Y := range warehouse {
		fmt.Println(string(Y))
	}
}

func calculateSumOfGPSCoords(warehouse [][]rune, box rune) int {
	sumOfGPCoords := 0
	for y, Y := range warehouse {
		for x, X := range Y {
			if X == box {
				sumOfGPCoords += 100*y + x
			}
		}
	}
	return sumOfGPCoords
}
func moveP2(x int, y int, dir rune, warehouse [][]rune) ([][]rune, int, int) {
	if dir == '>' {
		if warehouse[y][x+1] == '.' {
			warehouse[y][x] = '.'
			warehouse[y][x+1] = '@'
			return warehouse, x + 1, y
		} else if warehouse[y][x+1] == '#' {
			return warehouse, x, y
		} else if warehouse[y][x+1] == '[' {
			X := x + 1
			for warehouse[y][X] == '[' || warehouse[y][X] == ']' {
				X++
			}
			if warehouse[y][X] == '#' {
				return warehouse, x, y
			} else if warehouse[y][X] == '.' {
				for X != x {
					warehouse[y][X] = warehouse[y][X-1]
					X--
				}
				warehouse[y][x] = '.'
				// warehouse[y][x+1] = '@'
				return warehouse, x + 1, y
			}
		}
	} else if dir == '<' {
		if warehouse[y][x-1] == '.' {
			warehouse[y][x] = '.'
			warehouse[y][x-1] = '@'
			return warehouse, x - 1, y
		} else if warehouse[y][x-1] == '#' {
			return warehouse, x, y
		} else if warehouse[y][x-1] == ']' {
			X := x - 1
			for warehouse[y][X] == ']' || warehouse[y][X] == '[' {
				X--
			}
			if warehouse[y][X] == '#' {
				return warehouse, x, y
			} else if warehouse[y][X] == '.' {
				for X != x {
					warehouse[y][X] = warehouse[y][X+1]
					X++
				}
				warehouse[y][x] = '.'
				// warehouse[y][x-1] = '@'
				return warehouse, x - 1, y
			}
		}
	} else if dir == '^' {
		if warehouse[y-1][x] == '.' {
			warehouse[y-1][x] = '@'
			warehouse[y][x] = '.'
			return warehouse, x, y - 1
		} else if warehouse[y-1][x] == '#' {
			return warehouse, x, y
		} else if warehouse[y-1][x] == '[' {
			if canMoveUpOrDown('^', x, y-1, warehouse) {
				warehouse = moveBoxUp(x, y-1, warehouse)
				warehouse[y-1][x] = '@'
				warehouse[y][x] = '.'
				return warehouse, x, y - 1
			}
			return warehouse, x, y
		} else if warehouse[y-1][x] == ']' {
			if canMoveUpOrDown('^', x-1, y-1, warehouse) {
				warehouse = moveBoxUp(x-1, y-1, warehouse)
				warehouse[y][x] = '.'
				warehouse[y-1][x] = '@'
				return warehouse, x, y - 1
			}
			return warehouse, x, y
		}
	} else if dir == 'v' {
		if warehouse[y+1][x] == '.' {
			warehouse[y+1][x] = '@'
			warehouse[y][x] = '.'
			return warehouse, x, y + 1
		} else if warehouse[y+1][x] == '#' {
			return warehouse, x, y
		} else if warehouse[y+1][x] == '[' {
			if canMoveUpOrDown('v', x, y+1, warehouse) {
				warehouse = moveBoxDown(x, y+1, warehouse)
				warehouse[y+1][x] = '@'
				warehouse[y][x] = '.'
				return warehouse, x, y + 1
			}
			return warehouse, x, y
		} else if warehouse[y+1][x] == ']' {
			if canMoveUpOrDown('v', x-1, y+1, warehouse) {
				warehouse = moveBoxDown(x-1, y+1, warehouse)
				warehouse[y+1][x] = '@'
				warehouse[y][x] = '.'
				return warehouse, x, y + 1
			}
			return warehouse, x, y
		}
	}
	fmt.Println("Something is seriously wrong with this code ", dir)
	// fmt.Print("Or with me, if this line go executed")
	return warehouse, x, y
}

func canMoveUpOrDown(dir rune, x int, y int, warehouse [][]rune) bool {
	// x, y -> '['
	if dir == 'v' {
		canMoveDown := true
		if warehouse[y+1][x] == '.' && warehouse[y+1][x+1] == '.' {
			return true
		} else if warehouse[y+1][x] == '#' || warehouse[y+1][x+1] == '#' {
			return false
		} else if warehouse[y+1][x] == '[' {
			canMoveDown = canMoveDown && canMoveUpOrDown('v', x, y+1, warehouse)
		} else if warehouse[y+1][x] == ']' {
			canMoveDown = canMoveDown && canMoveUpOrDown('v', x-1, y+1, warehouse)
			if warehouse[y+1][x+1] == '[' {
				canMoveDown = canMoveDown && canMoveUpOrDown('v', x+1, y+1, warehouse)
			}
		} else if warehouse[y+1][x+1] == '[' {
			canMoveDown = canMoveDown && canMoveUpOrDown('v', x+1, y+1, warehouse)
		}
		return canMoveDown
	} else if dir == '^' {
		canMoveUp := true
		if warehouse[y-1][x] == '.' && warehouse[y-1][x+1] == '.' {
			return true
		} else if warehouse[y-1][x] == '#' || warehouse[y-1][x+1] == '#' {
			return false
		} else if warehouse[y-1][x] == '[' {
			canMoveUp = canMoveUp && canMoveUpOrDown('^', x, y-1, warehouse)
		} else if warehouse[y-1][x] == ']' {
			canMoveUp = canMoveUp && canMoveUpOrDown('^', x-1, y-1, warehouse)
			if warehouse[y-1][x+1] == '[' {
				canMoveUp = canMoveUp && canMoveUpOrDown('^', x+1, y-1, warehouse)
			}
		} else if warehouse[y-1][x+1] == '[' {
			canMoveUp = canMoveUp && canMoveUpOrDown('^', x+1, y-1, warehouse)
		}
		return canMoveUp
	}
	return false
}

func getWarehouse2Map(warehouse [][]rune) [][]rune {
	warehouse2 := make([][]rune, 0)
	for _, Y := range warehouse {
		line := make([]rune, 0)
		for _, X := range Y {
			if X == 'O' {
				line = append(line, '[', ']')
			} else if X == '.' {
				line = append(line, '.', '.')
			} else if X == '@' {
				line = append(line, '@', '.')
			} else if X == '#' {
				line = append(line, '#', '#')
			}
		}
		warehouse2 = append(warehouse2, line)
	}
	return warehouse2
}

func robotMakesMovesInWarehouse2(warehouse [][]rune, movements []rune) [][]rune {
	x, y := getRobotPos(warehouse)
	for _, dir := range movements {
		warehouse, x, y = moveP2(x, y, dir, warehouse)
	}
	return warehouse
}

func robotMakesMoves(warehouse [][]rune, movements []rune) [][]rune {
	x, y := getRobotPos(warehouse)

	for _, dir := range movements {
		warehouse, x, y = move(x, y, dir, warehouse)
	}
	return warehouse
}

func moveBoxUp(x int, y int, warehouse [][]rune) [][]rune {

	if warehouse[y-1][x] == '.' && warehouse[y-1][x+1] == '.' {
		warehouse[y-1][x] = '['
		warehouse[y-1][x+1] = ']'
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
	} else if warehouse[y-1][x] == '[' {
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
		warehouse = moveBoxUp(x, y-1, warehouse)
		warehouse[y-1][x] = '['
		warehouse[y-1][x+1] = ']'
	} else if warehouse[y-1][x] == ']' {
		warehouse = moveBoxUp(x-1, y-1, warehouse)
		warehouse[y-1][x-1] = '.'
		warehouse[y-1][x] = '.'
		if warehouse[y-1][x+1] == '[' {
			warehouse = moveBoxUp(x+1, y-1, warehouse)
			warehouse[y-1][x+2] = '.'
		}
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
		warehouse[y-1][x] = '['
		warehouse[y-1][x+1] = ']'
	} else if warehouse[y-1][x+1] == '[' {
		warehouse = moveBoxUp(x+1, y-1, warehouse)
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
		warehouse[y-1][x] = '['
		warehouse[y-1][x+1] = ']'
		warehouse[y-1][x+2] = '.'
	}
	return warehouse
}
func moveBoxDown(x int, y int, warehouse [][]rune) [][]rune {
	if warehouse[y+1][x] == '.' && warehouse[y+1][x+1] == '.' {
		warehouse[y+1][x] = '['
		warehouse[y+1][x+1] = ']'
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
	} else if warehouse[y+1][x] == '[' {
		warehouse = moveBoxDown(x, y+1, warehouse)
		warehouse[y+1][x] = '['
		warehouse[y+1][x+1] = ']'
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
	} else if warehouse[y+1][x] == ']' {
		warehouse = moveBoxDown(x-1, y+1, warehouse)
		warehouse[y+1][x-1] = '.'
		if warehouse[y+1][x+1] == '[' {
			warehouse = moveBoxDown(x+1, y+1, warehouse)
			warehouse[y+1][x+2] = '.'
		}
		warehouse[y+1][x] = '['
		warehouse[y+1][x+1] = ']'
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
	} else if warehouse[y+1][x+1] == '[' {
		warehouse = moveBoxDown(x+1, y+1, warehouse)
		warehouse[y][x] = '.'
		warehouse[y][x+1] = '.'
		warehouse[y+1][x] = '['
		warehouse[y+1][x+1] = ']'
	}
	return warehouse
}
func move(x int, y int, dir rune, warehouse [][]rune) ([][]rune, int, int) {
	if dir == '>' {
		if warehouse[y][x+1] == '.' {
			warehouse[y][x] = '.'
			warehouse[y][x+1] = '@'
			return warehouse, x + 1, y
		} else if warehouse[y][x+1] == '#' {
			return warehouse, x, y
		} else if warehouse[y][x+1] == 'O' {
			X := x + 1
			for warehouse[y][X] == 'O' {
				X++
			}
			if warehouse[y][X] == '#' {
				return warehouse, x, y
			} else if warehouse[y][X] == '.' {
				for X != x {
					warehouse[y][X] = warehouse[y][X-1]
					X--
				}
				warehouse[y][x] = '.'
				// warehouse[y][x+1] = '@'
				return warehouse, x + 1, y
			}
		}
	} else if dir == '<' {
		if warehouse[y][x-1] == '.' {
			warehouse[y][x] = '.'
			warehouse[y][x-1] = '@'
			return warehouse, x - 1, y
		} else if warehouse[y][x-1] == '#' {
			return warehouse, x, y
		} else if warehouse[y][x-1] == 'O' {
			X := x - 1
			for warehouse[y][X] == 'O' {
				X--
			}
			if warehouse[y][X] == '#' {
				return warehouse, x, y
			} else if warehouse[y][X] == '.' {
				for X != x {
					warehouse[y][X] = warehouse[y][X+1]
					X++
				}
				warehouse[y][x] = '.'
				// warehouse[y][x-1] = '@'
				return warehouse, x - 1, y
			}
		}
	} else if dir == '^' {
		if warehouse[y-1][x] == '.' {
			warehouse[y][x] = '.'
			warehouse[y-1][x] = '@'
			return warehouse, x, y - 1
		} else if warehouse[y-1][x] == '#' {
			return warehouse, x, y
		} else if warehouse[y-1][x] == 'O' {
			Y := y - 1
			for warehouse[Y][x] == 'O' {
				Y--
			}
			if warehouse[Y][x] == '#' {
				return warehouse, x, y
			} else if warehouse[Y][x] == '.' {
				for Y != y {
					warehouse[Y][x] = warehouse[Y+1][x]
					Y++
				}
				warehouse[y][x] = '.'
				// warehouse[y+1][x] = '@'
				return warehouse, x, y - 1
			}
		}
	} else if dir == 'v' {
		if warehouse[y+1][x] == '.' {
			warehouse[y][x] = '.'
			warehouse[y+1][x] = '@'
			return warehouse, x, y + 1
		} else if warehouse[y+1][x] == '#' {
			return warehouse, x, y
		} else if warehouse[y+1][x] == 'O' {
			Y := y + 1
			for warehouse[Y][x] == 'O' {
				Y++
			}
			if warehouse[Y][x] == '#' {
				return warehouse, x, y
			} else if warehouse[Y][x] == '.' {
				for Y != y {
					warehouse[Y][x] = warehouse[Y-1][x]
					Y--
				}
				warehouse[y][x] = '.'
				// warehouse[y-1][x] = '@'

				return warehouse, x, y + 1
			}
		}
	}
	return warehouse, x, y
}
