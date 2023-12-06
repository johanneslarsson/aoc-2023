package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TimeDistance struct {
	Time     int
	Distance int
}

func getPartOne(tds []TimeDistance) int {
	var results []int
	for _, td := range tds {
		var localCount int
		for i := 0; i < td.Time; i++ {
			speed := i
			travelTime := td.Time - i
			distance := speed * travelTime
			if distance > td.Distance {
				localCount++
			}
			//fmt.Println("speed:", speed, "travelTime:", travelTime, "distance:", distance)
		}
		results = append(results, localCount)
	}
	//fmt.Println(results)
	sum := 1
	for _, result := range results {
		sum *= result
	}
	return sum
}

func getPartTwo(tds []TimeDistance) int {
	timeVal := ""
	distanceVal := ""
	for _, td := range tds {
		timeVal += strconv.Itoa(td.Time)
		distanceVal += strconv.Itoa(td.Distance)
	}
	time, err := strconv.Atoi(timeVal)
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(distanceVal)
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := 0; i < time; i++ {
		speed := i
		travelTime := time - i
		localDistance := speed * travelTime
		if localDistance > distance {
			sum++
		}
		//fmt.Println("speed:", speed, "travelTime:", travelTime, "distance:", distance)
	}
	//fmt.Println(results)
	return sum
}

func getRows(filename string) []TimeDistance {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var timeDistances []TimeDistance
	var index int
	for scanner.Scan() {
		row := scanner.Text()
		if strings.HasPrefix(row, "Time: ") {
			fields := strings.Fields(strings.Split(row, "Time: ")[1])
			for _, field := range fields {
				val, err := strconv.Atoi(field)
				if err != nil {
					panic(err)
				}
				timeDistances = append(timeDistances, TimeDistance{Time: val})
			}
		}
		if strings.HasPrefix(row, "Distance: ") {
			fields := strings.Fields(strings.Split(row, "Distance: ")[1])
			for _, field := range fields {
				val, err := strconv.Atoi(field)
				if err != nil {
					panic(err)
				}
				timeDistances[index].Distance = val
				index++
			}
		}
	}
	fmt.Println(timeDistances)
	return timeDistances
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
