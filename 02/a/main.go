package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Record struct {
	Blue  int
	Green int
	Red   int
}

type Game struct {
	Records []Record
}

func getPartOne(games []Game) int {
	maxRecord := Record{14, 13, 12}
	sum := 0
	for i, game := range games {
		withinLimit := true
		for _, record := range game.Records {
			if maxRecord.Red < record.Red || maxRecord.Blue < record.Blue || maxRecord.Green < record.Green {
				withinLimit = false
				break
			}
		}
		if withinLimit {
			sum += i + 1
		}
	}
	return sum
}

func getPartTwo(games []Game) int {
	sum := 0
	for _, game := range games {
		maxRecord := Record{0, 0, 0}
		for _, record := range game.Records {
			if maxRecord.Red < record.Red {
				maxRecord.Red = record.Red
			}
			if maxRecord.Blue < record.Blue {
				maxRecord.Blue = record.Blue
			}
			if maxRecord.Green < record.Green {
				maxRecord.Green = record.Green
			}
		}
		sum += maxRecord.Red * maxRecord.Blue * maxRecord.Green
	}
	return sum
}

func getRows(filename string) []Game {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var games []Game
	for scanner.Scan() {
		row := scanner.Text()
		rowVal := strings.Split(row, ":")[1]
		recordsFields := strings.Split(rowVal, ";")
		game := Game{}
		game.Records = make([]Record, 0)
		for _, recordsValue := range recordsFields {
			r := Record{}
			recordPair := strings.Split(recordsValue, ",")
			for _, pair := range recordPair {
				var color string
				var count int
				fmt.Sscanf(pair, " %d %s", &count, &color)

				if color == "red" {
					r.Red = count
				} else if color == "blue" {
					r.Blue = count
				} else if color == "green" {
					r.Green = count
				}
			}
			game.Records = append(game.Records, r)
		}
		games = append(games, game)
	}
	return games
}

/*

func getRows(filename string) []Game {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var games []Game
	for scanner.Scan() {
		row := scanner.Text()
		rowVal := strings.Split(row, ":")[1]
		recordsFields := strings.Split(rowVal, ";")
		game := Game{}
		game.Records = make([]Record, 0)
		for _, recordsValue := range recordsFields {
			r := Record{}
			recordPair := strings.Split(recordsValue, ",")
			for _, pair := range recordPair {
				var color string
				var count int
				fmt.Sscanf(pair, " %d %s", &count, &color)

				if color == "red" {
					r.Red = count
				} else if color == "blue" {
					r.Blue = count
				} else if color == "green" {
					r.Green = count
				}
			}
			game.Records = append(game.Records, r)
		}
		games = append(games, game)
	}
	return games
}
*/

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
