package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Position struct {
	Y int
	X int
}

func getPartOne(rows []string) int {
	var numbers []int
	for posY, row := range rows {
		var numberVal string
		var adjacentToSymbol bool
		for posX, r := range row {
			if unicode.IsNumber(r) {
				numberVal += fmt.Sprintf("%c", r)
				if findAdjacentSymbol(rows, posY, posX) {
					adjacentToSymbol = true
				}
			} else if len(numberVal) > 0 {
				if adjacentToSymbol {
					val, err := strconv.Atoi(numberVal)
					if err != nil {
						panic(err)
					}
					numbers = append(numbers, val)
					adjacentToSymbol = false
				}
				numberVal = ""
			}
		}
		if len(numberVal) > 0 {
			//fmt.Println("border", numberVal)
			if adjacentToSymbol {
				val, err := strconv.Atoi(numberVal)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, val)
			}
		}
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func findAdjacentSymbol(rows []string, posY int, posX int) bool {
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			if x == 0 && y == 0 {
				continue
			}
			newY := posY + y
			newX := posX + x
			if newY < 0 || newY >= len(rows) {
				continue
			}
			if newX < 0 || newX >= len(rows[newY]) {
				continue
			}
			r := rune(rows[newY][newX])
			if !unicode.IsNumber(r) && r != '.' {
				return true
			}
		}
	}
	return false
}

func findAdjacentNumbers(rows []string, posY int, posX int) map[Position]int {
	result := map[Position]int{}
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			if x == 0 && y == 0 {
				continue
			}
			newY := posY + y
			newX := posX + x
			if newY < 0 || newY >= len(rows) {
				continue
			}
			if newX < 0 || newX >= len(rows[newY]) {
				continue
			}
			r := rune(rows[newY][newX])
			if unicode.IsNumber(r) {
				pos := Position{
					Y: newY,
					X: newX,
				}
				result[pos]++
			}
		}
	}
	return result
}

func getPartTwo(rows []string) int {
	sum := 0
	for posY, row := range rows {
		for posX, r := range row {
			if r == '*' {
				numbersPos := findAdjacentNumbers(rows, posY, posX)
				unique := findUniqueNumbersBasedOnPosition(rows, numbersPos)
				if len(unique) == 2 {
					sum += unique[0] * unique[1]
				}
			}
		}
	}
	return sum
}

func findUniqueNumbersBasedOnPosition(rows []string, positions map[Position]int) []int {
	visited := map[Position]int{}
	var result []int
	for position, _ := range positions {
		start := findStartOfNumber(rows, position)
		startPos := Position{Y: position.Y, X: position.X + start}
		if _, exist := visited[startPos]; exist {
			continue
		} else {
			visited[startPos]++
		}
		var numbers string
		for i := position.X + start; i < len(rows[position.Y]); i++ {
			r := rune(rows[position.Y][i])
			if unicode.IsNumber(r) {
				numbers += fmt.Sprintf("%c", r)
			} else {
				break
			}
		}
		//fmt.Println("Number val", numbers)
		val, err := strconv.Atoi(numbers)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}
	return result
}

func findStartOfNumber(rows []string, position Position) int {
	if position.X == 0 {
		return 0
	}
	for i := -1; ; i-- {
		newX := position.X + i
		if !unicode.IsNumber(rune(rows[position.Y][newX])) {
			return i + 1
		} else if newX == 0 {
			return i
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
