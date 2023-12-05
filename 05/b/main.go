package main

import (
	"bufio"
	"fmt"
	"os"
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

func getSourceNumber(maps []MapType, destination int) int {
	for _, mapSourceDest := range maps {
		if destination >= mapSourceDest.DestinationRangeStart && destination < mapSourceDest.DestinationRangeStart+mapSourceDest.RangeLength {
			val := mapSourceDest.SourceRangeStart + destination - mapSourceDest.DestinationRangeStart
			return val
		}
	}
	return destination
}

func getPartOne(food Food) int {
	seedsMap := make(map[int]struct{}, len(food.Seeds))
	for _, seed := range food.Seeds {
		seedsMap[seed] = struct{}{}
	}
	for i := 1; true; i++ {
		//fmt.Println("Location:", i)
		humidity := getSourceNumber(food.HumidityLocation, i)
		//fmt.Println("Humidity:", humidity)
		temperature := getSourceNumber(food.TemperatureHumidity, humidity)
		//fmt.Println("Temperature:", temperature)
		light := getSourceNumber(food.LightTemperature, temperature)
		//fmt.Println("Light:", light)
		water := getSourceNumber(food.WaterLightMap, light)
		//fmt.Println("Water:", water)
		fertilizer := getSourceNumber(food.FertilizerWaterMap, water)
		//fmt.Println("Fertilizer:", fertilizer)
		soil := getSourceNumber(food.SoilFertilizerMap, fertilizer)
		//fmt.Println("Soil:", soil)
		seed := getSourceNumber(food.SeedSoilMap, soil)
		if _, exists := seedsMap[seed]; exists {
			return i
		}
	}
	return 0
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
	fmt.Println("Calculating location")
	for i := 1; true; i++ {
		//fmt.Println("Location:", i)
		humidity := getSourceNumber(food.HumidityLocation, i)
		//fmt.Println("Humidity:", humidity)
		temperature := getSourceNumber(food.TemperatureHumidity, humidity)
		//fmt.Println("Temperature:", temperature)
		light := getSourceNumber(food.LightTemperature, temperature)
		//fmt.Println("Light:", light)
		water := getSourceNumber(food.WaterLightMap, light)
		//fmt.Println("Water:", water)
		fertilizer := getSourceNumber(food.FertilizerWaterMap, water)
		//fmt.Println("Fertilizer:", fertilizer)
		soil := getSourceNumber(food.SoilFertilizerMap, fertilizer)
		//fmt.Println("Soil:", soil)
		seed := getSourceNumber(food.SeedSoilMap, soil)
		//fmt.Println("Seed:", seed)
		if seed < len(seeds) && seeds[seed] == 1 {
			return i
		}
	}
	return 0
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
