package main

import (
	"math"
	"math/rand"
)

func distanceTo0(x, y float64) float64 {
	x2 := x * x
	y2 := y * y
	d := math.Sqrt(x2 + y2)
	return d
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func arrayRandFloat(min, max float64, size int) []float64 {
	a := make([]float64, size)
	for i := range a {
		a[i] = randFloat(min, max)
	}
	return a
}

func floatArraySum(a []float64, i0, iMax int) float64 {
	var sum float64
	aSum := a[i0:iMax]
	for _, v := range aSum {
		sum += v
	}
	return sum
}

func mapTo(x, in_min, in_max, out_min, out_max float64) float64 {
	return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min
}
