package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Hand struct {
	Card string
	Bid  int
}

type FirstRank int64

const (
	HighCard FirstRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type ByRank []Hand

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	iFirstOrdering := FirstOrdering(a[i].Card)
	jFirstOrdering := FirstOrdering(a[j].Card)
	if iFirstOrdering == jFirstOrdering {
		for k := 0; k < 5; k++ {
			iSecondOrdering := SecondOrdering(rune(a[i].Card[k]), 11)
			jSecondOrdering := SecondOrdering(rune(a[j].Card[k]), 11)
			if iSecondOrdering == jSecondOrdering {
				continue
			}
			return iSecondOrdering < jSecondOrdering
		}
		panic(fmt.Sprintf("not implemented %s, %s\n", a[i].Card, a[j].Card))
	}
	return iFirstOrdering < jFirstOrdering
}

type ByRankAdvanced []Hand

func (a ByRankAdvanced) Len() int      { return len(a) }
func (a ByRankAdvanced) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRankAdvanced) Less(i, j int) bool {
	iFirstOrdering := FirstOrderingAdvanced(a[i].Card)
	jFirstOrdering := FirstOrderingAdvanced(a[j].Card)
	if iFirstOrdering == jFirstOrdering {
		for k := 0; k < 5; k++ {
			iSecondOrdering := SecondOrdering(rune(a[i].Card[k]), 1)
			jSecondOrdering := SecondOrdering(rune(a[j].Card[k]), 1)
			if iSecondOrdering == jSecondOrdering {
				continue
			}
			return iSecondOrdering < jSecondOrdering
		}
		panic(fmt.Sprintf("not implemented %s, %s\n", a[i].Card, a[j].Card))
	}
	return iFirstOrdering < jFirstOrdering
}

func FirstOrdering(s string) FirstRank {
	set := make(map[rune]int, 5)
	equalCount := 0
	for _, c := range s {
		set[c]++
		if set[c] > equalCount {
			equalCount = set[c]
		}
	}
	switch equalCount {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		for _, v := range set {
			if v == 2 {
				return FullHouse
			}
		}
		return ThreeOfAKind
	case 2:
		countOfPairs := 0
		for _, v := range set {
			if v == 2 {
				countOfPairs++
			}
		}
		if countOfPairs == 2 {
			return TwoPair
		}
		return OnePair
	}
	return HighCard
}

func FirstOrderingAdvanced(s string) FirstRank {
	set := make(map[rune]int, 5)
	var equalCount int
	for _, c := range s {
		set[c]++
		if set[c] > equalCount {
			equalCount = set[c]
		}
	}
	jCount := set['J']
	switch equalCount {
	case 5:
		return FiveOfAKind
	case 4:
		if jCount == 1 || jCount == 4 {
			return FiveOfAKind
		}
		return FourOfAKind
	case 3:
		if jCount == 3 {
			for _, v := range set {
				if v == 2 {
					return FiveOfAKind
				}
			}
			return FourOfAKind
		} else if jCount == 2 {
			return FiveOfAKind
		} else if jCount == 1 {
			return FourOfAKind
		}

		for _, v := range set {
			if v == 2 {
				return FullHouse
			}
		}
		return ThreeOfAKind
	case 2:
		countOfPairs := 0
		for _, v := range set {
			if v == 2 {
				countOfPairs++
			}
		}
		if jCount == 2 {
			if countOfPairs == 2 {
				return FourOfAKind
			}
			return ThreeOfAKind
		} else if jCount == 1 {
			if countOfPairs == 2 {
				return FullHouse
			}
			return ThreeOfAKind
		}
		if countOfPairs == 2 {
			return TwoPair
		}
		return OnePair
	}
	if jCount == 1 {
		return OnePair
	}
	return HighCard
}

func SecondOrdering(c rune, jVal int) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return jVal
	case 'T':
		return 10
	default:
		return int(c - '0')
	}
}

func getPartOne(hands []Hand) int {
	fmt.Println(hands)
	sort.Sort(ByRank(hands))
	fmt.Println(hands)

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += hands[i].Bid * (i + 1)
	}
	return sum
}

func getPartTwo(hands []Hand) int {
	fmt.Println(hands)
	sort.Sort(ByRankAdvanced(hands))
	fmt.Println(hands)

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += hands[i].Bid * (i + 1)
	}
	return sum
}

func getRows(filename string) []Hand {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		row := scanner.Text()
		hand := Hand{}
		fmt.Sscanf(row, "%s %d", &hand.Card, &hand.Bid)
		hands = append(hands, hand)
	}
	fmt.Printf("%+v\n", hands)
	return hands
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
