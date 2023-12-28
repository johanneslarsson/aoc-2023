package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spring struct {
	Pattern string
	Damaged []int
}

func getSpringCount(springs []Spring) int {
	sum := 0
	for _, spring := range springs {
		damagedPattern := "."
		for _, val := range spring.Damaged {
			for i := 0; i < val; i++ {
				damagedPattern += "#"
			}
			damagedPattern += "."
		}
		states := make(map[int]int, len(spring.Pattern))
		states[0] = 1
		newStates := make(map[int]int, len(spring.Pattern))
		for _, char := range spring.Pattern {
			for key, _ := range states {
				switch char {
				case '?':
					if key+1 < len(damagedPattern) {
						newStates[key+1] += states[key]
					}
					if damagedPattern[key] == '.' {
						newStates[key] += states[key]
					}
				case '.':
					if key+1 < len(damagedPattern) && damagedPattern[key+1] == '.' {
						newStates[key+1] += states[key]
					}
					if damagedPattern[key] == '.' {
						newStates[key] += states[key]
					}
				case '#':
					if key+1 < len(damagedPattern) && damagedPattern[key+1] == '#' {
						newStates[key+1] += states[key]
					}
				}
			}
			states = newStates
			//fmt.Println(states)
			newStates = make(map[int]int, len(spring.Pattern))
		}

		damagedPatternLength := len(damagedPattern)
		sum += states[damagedPatternLength-1] + states[damagedPatternLength-2]
	}
	return sum
}

func getPartOne(springs []Spring) int {
	return getSpringCount(springs)
}

func getPartTwo(springs []Spring) int {
	fiveSprings := make([]Spring, 0, len(springs))
	for _, spring := range springs {
		newSpringDamaged := make([]int, 0, len(spring.Damaged)*5)
		for i := 0; i < 5; i++ {
			for _, val := range spring.Damaged {
				newSpringDamaged = append(newSpringDamaged, val)
			}
		}
		spring.Pattern += "?"
		newPattern := strings.Repeat(spring.Pattern, 5)
		if strings.HasSuffix(newPattern, "?") {
			newPattern = newPattern[:len(newPattern)-1]
		}
		//fmt.Println(newPattern, newSpringDamaged)
		fiveSprings = append(fiveSprings, Spring{newPattern, newSpringDamaged})
	}
	return getSpringCount(fiveSprings)
}

func getRows(filename string) []Spring {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var springs []Spring
	for scanner.Scan() {
		row := scanner.Text()
		spring := Spring{}
		fields := strings.Fields(row)
		spring.Pattern = fields[0]
		values := strings.Split(fields[1], ",")
		for _, val := range values {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			spring.Damaged = append(spring.Damaged, intVal)
		}
		springs = append(springs, spring)
	}
	return springs
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
