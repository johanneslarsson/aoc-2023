package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Y int
	X int
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func getPartOne(digs []DigType) int {
	y, x := 0, 0
	yLowest, xLowest := 0, 0
	yHighest, xHighest := 0, 0
	visitMap := make(map[Position]bool)
	visitMap[Position{y, x}] = true
	for _, dig := range digs {
		yNew, xNew := y, x
		switch dig.Direction {
		case "U":
			yNew -= dig.Amount
		case "D":
			yNew += dig.Amount
		case "L":
			xNew -= dig.Amount
		case "R":
			xNew += dig.Amount
		default:
			panic("Unknown direction")
		}
		yMin, yMax := minMax(y, yNew)
		xMin, xMax := minMax(x, xNew)
		for i := yMin; i <= yMax; i++ {
			for j := xMin; j <= xMax; j++ {
				visitMap[Position{i, j}] = true
			}
		}
		y, x = yNew, xNew
		if y > yHighest {
			yHighest = y
		}
		if x > xHighest {
			xHighest = x
		}
		if y < yLowest {
			yLowest = y
		}
		if x < xLowest {
			xLowest = x
		}
	}

	fmt.Println("yLowest:", yLowest, "yHighest:", yHighest, "xLowest:", xLowest, "xHighest:", xHighest)

	var nonOccupied []Position
	for y := yLowest; y <= yHighest; y++ {
		for x := xLowest; x <= xHighest; x++ {
			if _, ok := visitMap[Position{y, x}]; !ok {
				nonOccupied = append(nonOccupied, Position{y, x})
			}
		}
	}

	insideCount := 0
	insideMap := make(map[Position]bool, len(digs)*9)
	for _, pos := range nonOccupied {
		wallCount := 0
		isWallUp, isWallDown := false, false
		for x := pos.X; x <= xHighest; x++ {
			if _, ok := visitMap[Position{pos.Y, x}]; ok {
				if _, ok := visitMap[Position{pos.Y - 1, x}]; ok {
					isWallUp = true
				}
				if _, ok := visitMap[Position{pos.Y + 1, x}]; ok {
					isWallDown = true
				}
				if isWallUp && isWallDown {
					wallCount++
				}
			} else {
				isWallUp, isWallDown = false, false
			}
		}

		if wallCount%2 != 0 {
			insideCount++
			insideMap[pos] = true
		}
	}

	for y := yLowest; y <= yHighest; y++ {
		for x := xLowest; x <= xHighest; x++ {
			if _, ok := visitMap[Position{y, x}]; ok {
				fmt.Print("X")
			} else if _, ok := insideMap[Position{y, x}]; ok {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	//	fmt.Println(visitMap)
	//	sum := 0
	return insideCount + len(visitMap)
}

func shoeLaceFormula(corners []Position) int {
	sum := 0
	for i := 0; i < len(corners)-1; i++ {
		sum += corners[i].X*corners[i+1].Y - corners[i].Y*corners[i+1].X
	}
	return sum / 2
}

func getPartTwo(digs []DigType) int {
	y, x := 0, 0
	corners := make([]Position, 0, len(digs)+1)
	length := 0
	for _, dig := range digs {
		switch dig.HexColorDirection {
		case "U":
			y -= dig.HexColorAmount
		case "D":
			y += dig.HexColorAmount
		case "L":
			x -= dig.HexColorAmount
		case "R":
			x += dig.HexColorAmount
		}
		length += dig.HexColorAmount
		corners = append(corners, Position{y, x})
	}
	shoeLaceValue := shoeLaceFormula(corners)
	return shoeLaceValue + length/2 + 1
}

func getDirection(hexValue string) string {
	switch hexValue {
	case "0":
		return "R"
	case "1":
		return "D"
	case "2":
		return "L"
	case "3":
		return "U"
	default:
		panic("Unknown direction")
	}
}

type DigType struct {
	Direction         string
	Amount            int
	HexColorAmount    int
	HexColorDirection string
}

func getRows(filename string) []DigType {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var digs []DigType
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)
		hexColor := strings.Replace(fields[2], "(", "", 1)
		hexColor = strings.Replace(hexColor, ")", "", 1)
		hexColor = strings.Replace(hexColor, "#", "", 1)
		hexColorAmountVal := hexColor[:len(hexColor)-1]
		hexColorAmount, err := strconv.ParseInt(hexColorAmountVal, 16, 64)
		if err != nil {
			panic(err)
		}
		amount, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		dig := DigType{
			Direction:         fields[0],
			Amount:            amount,
			HexColorAmount:    int(hexColorAmount),
			HexColorDirection: getDirection(hexColor[len(hexColor)-1:]),
		}
		digs = append(digs, dig)
	}
	return digs
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
