package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

const (
	MaxUint   = ^uint(0)
	MaxUint16 = ^uint16(0)
	MinUint   = 0
	MaxInt    = int(MaxUint >> 1)
	MinInt    = -MaxInt - 1
)

func main() {
	//Pi(10000000000)

	e := Integration(f, 0, math.Pi, 100000)
	fmt.Println(e, "Diff:", 2.0-e)

}

func f(x float64) (y float64) {
	y = math.Sin(x)
	return
}
