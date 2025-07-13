package main

import (
	"fmt"
	"math"
)

func GroupTemperatures(temps []float64) map[float64][]float64 {
	tempGroup := make(map[float64][]float64)

	for _, temp := range temps {
		var key float64
		if temp < 0 {
			key = math.Ceil(temp/10) * 10
		} else {
			key = math.Floor(temp/10) * 10
		}
		tempGroup[key] = append(tempGroup[key], temp)
	}

	return tempGroup
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	grouped := GroupTemperatures(temps)

	for key, group := range grouped {
		fmt.Printf("%.0f: %v\n", key, group)
	}
}
