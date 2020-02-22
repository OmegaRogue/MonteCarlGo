package main

import (
	"MonteCarlGo/qrng"
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

//func Integration1D(f func(float64) float64, a, b float64, N int) float64 {
//
//
//	var t0 int64
//	var tEnd int64
//	var e float64
//	t0 = time.Now().UnixNano()
//	ba := b - a
//
//	R := arrayRandFloat(a, b, N)
//	fR := make([]float64, N)
//
//	for i, v := range R {
//		fR[i] = f(v)
//	}
//	sumFR := floatArraySum(fR, 0, N-1)
//
//	e = ba * (1.0 / float64(N)) * sumFR
//
//	tEnd = time.Now().UnixNano()
//
//	fmt.Println("The Final Estimate:", e, "Done after", tEnd-t0, "nanoseconds")
//
//	return e
//}

func Integration(f func(float64) float64, a, b float64, N int) float64 {

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var wg sync.WaitGroup

	ba := b - a

	var e float64
	t0 := time.Now()
	wg.Add(cores)

	var sum float64
	for i := 0; i < cores; i++ {
		go integrate(f, a, b, N/cores, &wg, &sum)
	}

	wg.Wait()

	e = ba * (1.0 / float64(N)) * sum

	fmt.Println("The Final Estimate:", e, "Done after", time.Since(t0).String())

	return e
}

func integrate(f func(float64) float64, a, b float64, n int, wg *sync.WaitGroup, sum *float64) {
	defer wg.Done()
	rng, _ := qrng.GetFloat64(n)
	for i := 0; i < n; i++ {
		*sum += f(mapTo(rng[i], math.SmallestNonzeroFloat64, math.MaxFloat64, a, b))
	}

}
