package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Id      int
	Winning []int
	Yours   []int
}

func getPartOne(cards []Card) int {
	sum := 0
	for _, card := range cards {
		var localCount int
		for _, your := range card.Yours {
			if slices.Contains(card.Winning, your) {
				localCount++
			}
		}
		sum += int(math.Pow(2, float64(localCount)-1))
	}
	return sum
}

func getPartTwo(cards []Card) int {
	result := len(cards)
	queue := cards[:]
	cardMap := map[int]int{}
	for i := 0; i < len(queue); i++ {
		card := queue[i]
		var localCount int
		if _, exist := cardMap[card.Id]; exist {
			localCount = cardMap[card.Id]
		} else {
			for _, your := range card.Yours {
				if slices.Contains(card.Winning, your) {
					localCount++
				}
			}
			cardMap[card.Id] = localCount
		}
		for j := 1; j <= localCount; j++ {
			queue = append(queue, cards[card.Id+j])
		}
		result += localCount
	}

	fmt.Println(len(cardMap))
	return result
}

func getRows(filename string) []Card {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cards []Card
	for scanner.Scan() {
		row := scanner.Text()
		card := Card{}
		card.Id = len(cards)
		headerValue := strings.Split(row, ":")
		values := strings.Split(headerValue[1], "|")

		winValues := strings.Fields(values[0])
		for _, winVal := range winValues {
			val, err := strconv.Atoi(winVal)
			if err != nil {
				panic(err)
			}
			card.Winning = append(card.Winning, val)
		}
		yourValues := strings.Fields(values[1])
		for _, yourVal := range yourValues {
			val, err := strconv.Atoi(yourVal)
			if err != nil {
				panic(err)
			}
			card.Yours = append(card.Yours, val)
		}
		cards = append(cards, card)
	}
	return cards
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
