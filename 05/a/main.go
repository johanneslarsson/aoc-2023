package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Food struct {
	Seeds               []int
	SeedSoilMap         []MapType
	SoilFertilizerMap   []MapType
	FertilizerWaterMap  []MapType
	WaterLightMap       []MapType
	LightTemperature    []MapType
	TemperatureHumidity []MapType
	HumidityLocation    []MapType
}

type MapType struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func getDestinationNumber(maps *[]MapType, source int) int {
	//fmt.Printf("%+v\n", maps)
	for _, mapSourceDest := range *maps {
		if source >= mapSourceDest.SourceRangeStart && source < mapSourceDest.SourceRangeStart+mapSourceDest.RangeLength {
			val := mapSourceDest.DestinationRangeStart + source - mapSourceDest.SourceRangeStart
			return val
		}
	}
	return source
}

func getSourceNumber(maps []MapType, destination int) (result []int) {
	//fmt.Printf("%+v\n", maps)
	for _, mapSourceDest := range maps {
		if destination >= mapSourceDest.DestinationRangeStart && destination < mapSourceDest.DestinationRangeStart+mapSourceDest.RangeLength {
			val := mapSourceDest.SourceRangeStart + destination - mapSourceDest.DestinationRangeStart
			result = append(result, val)
		}
	}
	if len(result) == 0 {
		result = append(result, destination)
	}
	return result
}

func getPartOne(food Food) int {
	var result []int
	for _, seed := range food.Seeds {
		fmt.Println("Seed:", seed)
		soil := getDestinationNumber(&food.SeedSoilMap, seed)
		fmt.Println("Soil:", soil)
		fertilizer := getDestinationNumber(&food.SoilFertilizerMap, soil)
		fmt.Println("Fertilizer:", fertilizer)
		water := getDestinationNumber(&food.FertilizerWaterMap, fertilizer)
		fmt.Println("Water:", water)
		light := getDestinationNumber(&food.WaterLightMap, water)
		fmt.Println("Light:", light)
		temperature := getDestinationNumber(&food.LightTemperature, light)
		fmt.Println("Temperature:", temperature)
		humidity := getDestinationNumber(&food.TemperatureHumidity, temperature)
		fmt.Println("Humidity:", humidity)
		location := getDestinationNumber(&food.HumidityLocation, humidity)
		fmt.Println("Location:", location)
		result = append(result, location)
	}
	sort.Ints(result)
	return result[0]
}

func getPartTwo(food Food) int {
	fmt.Println("Calculating size")
	total := 0
	for i := 0; i < len(food.Seeds); i = i + 2 {
		if food.Seeds[i]+food.Seeds[i+1] > total {
			total = food.Seeds[i] + food.Seeds[i+1]
		}
	}
	fmt.Println("Size:", total)

	seeds := make([]int, total)

	fmt.Println("Generating seeds")
	for i := 0; i < len(food.Seeds); i = i + 2 {
		for j := 0; j < food.Seeds[i+1]; j++ {
			seeds[food.Seeds[i]+j] = 1
		}
	}
	fmt.Println("Calculating locations...")

	minVal := math.MaxInt
	for seed, _ := range seeds {
		if seeds[seed] == 0 {
			continue
		}
		//fmt.Println("Seed:", seed)
		soil := getDestinationNumber(&food.SeedSoilMap, seed)
		//fmt.Println("Soil:", soil)
		fertilizer := getDestinationNumber(&food.SoilFertilizerMap, soil)
		//fmt.Println("Fertilizer:", fertilizer)
		water := getDestinationNumber(&food.FertilizerWaterMap, fertilizer)
		//fmt.Println("Water:", water)
		light := getDestinationNumber(&food.WaterLightMap, water)
		//fmt.Println("Light:", light)
		temperature := getDestinationNumber(&food.LightTemperature, light)
		//fmt.Println("Temperature:", temperature)
		humidity := getDestinationNumber(&food.TemperatureHumidity, temperature)
		//fmt.Println("Humidity:", humidity)
		location := getDestinationNumber(&food.HumidityLocation, humidity)
		//fmt.Println("Location:", location)
		//result = append(result, location[0])
		if location < minVal {
			minVal = location
		}
	}
	//sort.Ints(result)
	return minVal
}

func getMapOrder(food *Food, index int) *[]MapType {
	switch index {
	case 1:
		return &food.SeedSoilMap
	case 2:
		return &food.SoilFertilizerMap
	case 3:
		return &food.FertilizerWaterMap
	case 4:
		return &food.WaterLightMap
	case 5:
		return &food.LightTemperature
	case 6:
		return &food.TemperatureHumidity
	case 7:
		return &food.HumidityLocation
	default:
		panic("Unknown index")
	}

}

func getRows(filename string) Food {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var food Food
	var index int
	var mapIndex int
	var expectMapValues bool
	for scanner.Scan() {
		row := scanner.Text()
		if index == 0 {
			values := strings.Fields(strings.Split(row, ": ")[1])
			for _, value := range values {
				val, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}
				food.Seeds = append(food.Seeds, val)
			}
		}
		if len(row) == 0 {
			expectMapValues = false
		} else if expectMapValues {
			mapType := MapType{}
			fmt.Sscanf(row, "%d %d %d",
				&mapType.DestinationRangeStart, &mapType.SourceRangeStart,
				&mapType.RangeLength)
			mapToAppend := getMapOrder(&food, mapIndex)
			*mapToAppend = append(*mapToAppend, mapType)
		} else if strings.HasSuffix(row, "map:") {
			mapIndex++
			expectMapValues = true
		}
		index++
	}
	fmt.Printf("%+v\n", food)
	return food
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
