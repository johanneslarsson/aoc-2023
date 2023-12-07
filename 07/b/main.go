package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

type ByRankPartTwo []Hand

func (a ByRankPartTwo) Len() int      { return len(a) }
func (a ByRankPartTwo) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRankPartTwo) Less(i, j int) bool {
	iFirstOrdering := FirstOrderingWithJoker(a[i].Card)
	jFirstOrdering := FirstOrderingWithJoker(a[j].Card)
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
	for _, c := range s {
		set[c]++
	}
	equalCounts := make([]int, 0, len(set))
	for _, v := range set {
		equalCounts = append(equalCounts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(equalCounts)))
	if reflect.DeepEqual(equalCounts, []int{5}) {
		return FiveOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{4, 1}) {
		return FourOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{3, 2}) {
		return FullHouse
	} else if reflect.DeepEqual(equalCounts, []int{3, 1, 1}) {
		return ThreeOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{2, 2, 1}) {
		return TwoPair
	} else if reflect.DeepEqual(equalCounts, []int{2, 1, 1, 1}) {
		return OnePair
	} else if reflect.DeepEqual(equalCounts, []int{1, 1, 1, 1, 1}) {
		return HighCard
	}
	panic(fmt.Sprintf("not implemented %s\n", s))
}

func FirstOrderingWithJoker(s string) FirstRank {
	set := make(map[rune]int, 5)
	jCount := 0
	for _, c := range s {
		if c == 'J' {
			jCount++
		} else {
			set[c]++
		}
	}
	equalCounts := make([]int, 0, len(set))
	for _, v := range set {
		equalCounts = append(equalCounts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(equalCounts)))
	if len(equalCounts) == 0 {
		return FiveOfAKind
	}
	equalCounts[0] += jCount
	if reflect.DeepEqual(equalCounts, []int{5}) {
		return FiveOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{4, 1}) {
		return FourOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{3, 2}) {
		return FullHouse
	} else if reflect.DeepEqual(equalCounts, []int{3, 1, 1}) {
		return ThreeOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{2, 2, 1}) {
		return TwoPair
	} else if reflect.DeepEqual(equalCounts, []int{2, 1, 1, 1}) {
		return OnePair
	} else if reflect.DeepEqual(equalCounts, []int{1, 1, 1, 1, 1}) {
		return HighCard
	}
	panic(fmt.Sprintf("not implemented %s\n", s))
}

/*
func FirstOrderingWithJoker(s string) FirstRank {
	set := make(map[rune]int, 5)
	for _, c := range s {
		set[c]++
	}
	jCount := set['J']
	equalCounts := make([]int, 0, len(set))
	for _, v := range set {
		equalCounts = append(equalCounts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(equalCounts)))
	if reflect.DeepEqual(equalCounts, []int{5}) {
		return FiveOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{4, 1}) {
		if jCount == 1 || jCount == 4 {
			return FiveOfAKind
		}
		return FourOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{3, 2}) {
		if jCount == 3 || jCount == 2 {
			return FiveOfAKind
		}
		return FullHouse
	} else if reflect.DeepEqual(equalCounts, []int{3, 1, 1}) {
		if jCount == 3 || jCount == 1 {
			return FourOfAKind
		}
		return ThreeOfAKind
	} else if reflect.DeepEqual(equalCounts, []int{2, 2, 1}) {
		if jCount == 2 {
			return FourOfAKind
		} else if jCount == 1 {
			return FullHouse
		}
		return TwoPair
	} else if reflect.DeepEqual(equalCounts, []int{2, 1, 1, 1}) {
		if jCount == 2 || jCount == 1 {
			return ThreeOfAKind
		}
		return OnePair
	} else if reflect.DeepEqual(equalCounts, []int{1, 1, 1, 1, 1}) {
		if jCount == 1 {
			return OnePair
		}
		return HighCard
	}
	panic(fmt.Sprintf("not implemented %s\n", s))
}
*/

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
	sort.Sort(ByRankPartTwo(hands))
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
