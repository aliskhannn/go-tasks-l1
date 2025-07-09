package main

import (
	"fmt"
	"math"
)

// Группировка температур
// Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить эти значения в группы с шагом 10 градусов.
//
// Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
//
// Пояснение: диапазон -20 включает значения от -20 до -29.9, диапазон 10 – от 10 до 19.9, и т.д.
// Порядок в подмножествах не важен.

func main() {
	tempGroup := make(map[float64][]float64)
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, temp := range temps {
		var key float64
		if temp < 0 {
			key = math.Ceil(temp/10) * 10 // for negative temperatures, round up to nearest 10 (e.g., -27 → -20)
			// Example: -25.4 / 10 = -2.54
			// math.Ceil(-2.54) = -2
			// -2 * 10 = -20
		} else {
			key = math.Floor(temp/10) * 10 // for positive temperatures, round down to nearest 10 (e.g., 15.5 → 10)
			// Example: 15.5 / 10 = 1.55
			// math.Floor(1.55) = 1
			// 1 * 10 = 10
		}

		// append the temperature to the corresponding group in the map
		tempGroup[key] = append(tempGroup[key], temp)
	}

	for key, group := range tempGroup {
		fmt.Printf("%.0f: %.1f\n", key, group) // format and print results
	}
}
