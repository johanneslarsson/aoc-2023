package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	y int
	x int
}

func getMinMax(a, b int) (min int, max int) {
	if a > b {
		return b, a
	}
	return a, b
}
func calculateDistance(rows []string, spaceExpansion int) int {
	galaxyList := make([]Position, 0, len(rows))
	occupiedYLines := make([]bool, len(rows))
	occupiedXLines := make([]bool, len(rows[0]))
	for y, row := range rows {
		for x, char := range row {
			if char == '#' {
				occupiedYLines[y] = true
				occupiedXLines[x] = true
				galaxyList = append(galaxyList, Position{y, x})
			}
		}
	}

	distance := 0
	for i := 0; i < len(galaxyList); i++ {
		for j := i + 1; j < len(galaxyList); j++ {
			yU, xU := galaxyList[i].y, galaxyList[i].x
			yI, xI := galaxyList[j].y, galaxyList[j].x
			minY, maxY := getMinMax(yU, yI)
			extraYCount := 0
			for k := minY + 1; k < maxY; k++ {
				if !occupiedYLines[k] {
					extraYCount += spaceExpansion
				}
			}
			yDiff := maxY - minY + extraYCount

			minX, maxX := getMinMax(xU, xI)
			extraXCount := 0
			for k := minX + 1; k < maxX; k++ {
				if !occupiedXLines[k] {
					extraXCount += spaceExpansion
				}
			}
			xDiff := maxX - minX + extraXCount
			distance += xDiff + yDiff
			//fmt.Println(i+1, j+1, xDiff+yDiff)
		}
	}
	return distance
}

func getPartOne(rows []string) int {
	return calculateDistance(rows, 2-1)
}

func getPartTwo(rows []string) int {
	return calculateDistance(rows, 1000000-1)
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
