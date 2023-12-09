package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func calculateNext(row []int) int {
	sum := 0
	var queue [][]int
	queue = append(queue, row)
	lastDigits := make([]int, 0, len(row))
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		differences := make([]int, 0, len(row))
		allZero := true
		for i := 1; i < len(first); i++ {
			diff := first[i] - first[i-1]
			if diff != 0 {
				allZero = false
			}
			differences = append(differences, diff)
		}
		lastDigits = append(lastDigits, first[len(first)-1])
		if allZero {
			for _, digit := range lastDigits {
				sum += digit
			}
			return sum
		}
		queue = append(queue, differences)
	}
	panic("no solution")
}

func getPartOne(values [][]int) int {
	sum := 0
	for _, row := range values {
		sum += calculateNext(row)
	}
	return sum
}

func getPartTwo(values [][]int) int {
	sum := 0
	for _, row := range values {
		slices.Reverse(row)
		sum += calculateNext(row)
	}
	return sum
}

func getRows(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var history [][]int
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)
		var rowInts []int
		for _, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			rowInts = append(rowInts, val)
		}
		history = append(history, rowInts)
	}
	fmt.Printf("%+v\n", history)
	return history
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
