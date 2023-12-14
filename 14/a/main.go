package main

import (
	"bufio"
	"fmt"
	"os"
)

type StateValue struct {
	Index int
	Load  int
}

func getRoundCubeRocks(rows []string) ([][]int, [][]int) {
	roundRocks := make([][]int, len(rows))
	cubeRocks := make([][]int, len(rows))
	for y, row := range rows {
		roundRocks[y] = make([]int, len(row))
		cubeRocks[y] = make([]int, len(row))
		for x, char := range row {
			if char == 'O' {
				roundRocks[y][x] = 1
			} else if char == '#' {
				cubeRocks[y][x] = 1
			}
		}
		//	fmt.Println(row)
	}
	return roundRocks, cubeRocks
}

func getPartOne(rows []string) int {
	roundRocks, cubeRocks := getRoundCubeRocks(rows)
	goNorth(roundRocks, cubeRocks)

	return calculateLoad(roundRocks)
}

func getPartTwo(rows []string) int {
	roundRocks, cubeRocks := getRoundCubeRocks(rows)
	stateMap := make(map[string]StateValue)

	start := 0
	length := 0
	for i := 0; true; i++ {
		for j := 0; j < 4; j++ {
			if j == 0 {
				goNorth(roundRocks, cubeRocks)
			} else if j == 1 {
				goWest(roundRocks, cubeRocks)
			} else if j == 2 {
				goSouth(roundRocks, cubeRocks)
			} else if j == 3 {
				goEast(roundRocks, cubeRocks)
			}
		}
		state := getState(roundRocks, cubeRocks)
		if val, ok := stateMap[state]; ok {
			start = val.Index
			length = i - val.Index
			break
		} else {
			stateMap[state] = StateValue{i, calculateLoad(roundRocks)}
		}
	}

	calculatedIndex := (1000000000-1-start)%length + start

	for _, val := range stateMap {
		if val.Index == calculatedIndex {
			return val.Load
		}
	}

	panic("NOT FOUND")
}

func getState(roundRocks [][]int, cubeRocks [][]int) string {
	state := ""
	for y, row := range roundRocks {
		for x, rock := range row {
			if cubeRocks[y][x] == 1 {
				state += "#"
			} else if rock == 1 {
				state += "O"
			} else {
				state += "."
			}
		}
	}
	return state

}

func calculateLoad(roundRocks [][]int) (sum int) {
	for i, row := range roundRocks {
		load := len(roundRocks) - i
		count := 0
		for _, rock := range row {
			if rock == 1 {
				count++
			}
		}
		sum += count * load
	}
	return sum
}

func goNorth(roundRocks [][]int, cubeRocks [][]int) {
	move(roundRocks, cubeRocks, -1, 0)
}

func goWest(roundRocks [][]int, cubeRocks [][]int) {
	move(roundRocks, cubeRocks, 0, -1)
}

func goSouth(roundRocks [][]int, cubeRocks [][]int) {
	move(roundRocks, cubeRocks, 1, 0)
}

func goEast(roundRocks [][]int, cubeRocks [][]int) {
	move(roundRocks, cubeRocks, 0, 1)
}

func move(roundRocks [][]int, cubeRocks [][]int, yMove int, xMove int) {
	moves := -1
	for moves > 0 || moves == -1 {
		moves = 0
		for y, row := range roundRocks {
			for x, rock := range row {
				if rock == 1 {
					newY := y
					newX := x
					for {
						if (yMove < 0 && newY == 0) ||
							(xMove < 0 && newX == 0) ||
							(yMove > 0 && newY == len(roundRocks)-1) ||
							(xMove > 0 && newX == len(row)-1) {
							break
						}

						newY += yMove
						newX += xMove
						if roundRocks[newY][newX] == 1 || cubeRocks[newY][newX] == 1 {
							break
						}
						roundRocks[newY][newX] = 1
						roundRocks[newY-yMove][newX-xMove] = 0
						moves++
					}
				}
			}
		}
	}
}

func getRows(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}
	return rows
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
