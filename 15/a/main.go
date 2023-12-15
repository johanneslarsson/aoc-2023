package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPartOne(instructions []Instruction) int {
	sum := 0
	for _, instruction := range instructions {
		sum += hash(getWholeWord(instruction))
	}
	return sum
}

func hash(s string) int {
	val := 0
	for _, c := range s {
		val += int(c)
		val *= 17
		val %= 256
	}
	return val
}

func getWholeWord(instruction Instruction) string {
	if instruction.Value == -1 {
		return fmt.Sprintf("%s%s", instruction.Prefix, instruction.Operation)
	}
	return fmt.Sprintf("%s%s%d", instruction.Prefix, instruction.Operation, instruction.Value)
}

func getPartTwo(instructions []Instruction) int {
	boxes := make(map[int][]Instruction, 256)
	for _, instruction := range instructions {
		prefixHashVal := hash(instruction.Prefix)
		if instruction.Operation == "=" {
			exist := false
			for i, lens := range boxes[prefixHashVal] {
				if lens.Prefix == instruction.Prefix {
					boxes[prefixHashVal][i] = instruction
					exist = true
					break
				}
			}
			if !exist {
				boxes[prefixHashVal] = append(boxes[prefixHashVal], instruction)
			}
		} else if instruction.Operation == "-" {
			for i, lens := range boxes[prefixHashVal] {
				if lens.Prefix == instruction.Prefix {
					boxes[prefixHashVal] = append(boxes[prefixHashVal][:i], boxes[prefixHashVal][i+1:]...)
					break
				}
			}
		}
	}

	sum := 0
	for key, box := range boxes {
		for i, val := range box {
			sum += (key + 1) * (i + 1) * val.Value
		}
	}
	return sum
}

func getPrefixOperation(word string) (prefix string, operation string, value int) {
	prefix, focalLength, found := strings.Cut(word, "=")
	operation = "="
	if !found {
		prefix, focalLength, found = strings.Cut(word, "-")
		operation = "-"
	}
	if focalLength == "" {
		return prefix, operation, -1
	}
	value, err := strconv.Atoi(focalLength)
	if err != nil {
		panic(err)
	}
	return prefix, operation, value
}

type Instruction struct {
	Prefix    string
	Value     int
	Operation string
}

func getRows(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var instructions []Instruction
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Split(row, ",")

		for _, field := range fields {
			var instruction Instruction
			prefix, operation, value := getPrefixOperation(field)
			instruction.Prefix = prefix
			instruction.Operation = operation
			instruction.Value = value

			instructions = append(instructions, instruction)
		}
	}
	return instructions
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
