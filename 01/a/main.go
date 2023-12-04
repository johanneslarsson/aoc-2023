package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func calculateFirstLast(rows []string, replaceWords bool) int {
	numbers := make([][]string, len(rows))
	words := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
		"six": "6", "seven": "7", "eight": "8", "nine": "9"}
	index := 0
	sum := 0
	for _, row := range rows {
		for i := 0; i < len(row); i++ {
			r := rune(row[i])
			if unicode.IsNumber(r) {
				numbers[index] = append(numbers[index], fmt.Sprintf("%c", r))
			} else if replaceWords {
				for key, val := range words {
					if strings.HasPrefix(row[i:], key) {
						numbers[index] = append(numbers[index], val)
						break
					}
				}
			}
		}
		first := numbers[index][0]
		last := numbers[index][len(numbers[index])-1]
		val, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += val
		index++
	}
	return sum
}

func getPartOne(rows []string) int {
	return calculateFirstLast(rows, false)
}

func getPartTwo(rows []string) int {
	return calculateFirstLast(rows, true)
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
