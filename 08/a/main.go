package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Network struct {
	Instructions string
	NodeMap      map[string]Node
}

type Node struct {
	Left  string
	Right string
}

func getPartOne(network Network) int {
	pos := "AAA"
	for i := 0; true; i++ {
		r := rune(network.Instructions[i%len(network.Instructions)])
		//fmt.Printf("%c\n", r)
		if r == 'R' {
			pos = network.NodeMap[pos].Right
		} else {
			pos = network.NodeMap[pos].Left
		}

		if pos == "ZZZ" {
			return i + 1
		}
	}
	panic("not implemented")
}

func getPartTwo(network Network) int {
	positions := make([]string, 0, len(network.NodeMap))
	for key, _ := range network.NodeMap {
		if strings.HasSuffix(key, "A") {
			positions = append(positions, key)
		}
	}
	fmt.Println(positions)
	var result []int
	for j := 0; j < len(positions); j++ {
		for i := 0; true; i++ {
			r := rune(network.Instructions[i%len(network.Instructions)])
			if r == 'R' {
				positions[j] = network.NodeMap[positions[j]].Right
			} else {
				positions[j] = network.NodeMap[positions[j]].Left
			}
			if strings.HasSuffix(positions[j], "Z") {
				result = append(result, i+1)
				break
			}
		}
	}

	fmt.Println(result)
	sum := LeastCommonMultiplier(result[0], result[1], result[2:]...)
	return sum
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LeastCommonMultiplier) via GCD
func LeastCommonMultiplier(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiplier(result, integers[i])
	}

	return result
}

func getRows(filename string) Network {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var network Network
	var index int
	for scanner.Scan() {
		row := scanner.Text()
		if index == 0 {
			network.Instructions = row
		} else if index == 1 {
			network.NodeMap = make(map[string]Node)
		} else {
			var node Node
			var start string
			row = strings.Replace(row, ",", "", 1)
			row = strings.Replace(row, "(", "", 1)
			row = strings.Replace(row, ")", "", 1)
			fmt.Sscanf(row, "%s = %s %s", &start, &node.Left, &node.Right)
			network.NodeMap[start] = node
		}
		index++
	}
	fmt.Printf("%+v\n", network)
	return network
}
func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
