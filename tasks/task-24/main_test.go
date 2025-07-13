package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		p1       *Point
		p2       *Point
		expected float64
	}{
		{NewPoint(0, 0), NewPoint(3, 4), 5},                  // классический 3-4-5 треугольник
		{NewPoint(1, 2), NewPoint(1, 2), 0},                  // одинаковые точки
		{NewPoint(-1, -1), NewPoint(2, 3), 5},                // отрицательные координаты
		{NewPoint(0, 0), NewPoint(0, 0), 0},                  // нулевая дистанция
		{NewPoint(2.5, 3.5), NewPoint(4.5, 7.5), 4.47213595}, // с плавающей точкой
	}

	for _, tt := range tests {
		result := tt.p1.Distance(tt.p2)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("Distance(%v, %v) = %f; want %f", tt.p1, tt.p2, result, tt.expected)
		}
	}
}
