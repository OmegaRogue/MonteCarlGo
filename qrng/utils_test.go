package qrng

import "testing"

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func TestSum(t *testing.T) {

	for u := uint(MinUint); u < MaxUint; u++ {
		s := unsignedToSigned(u)
		t.Logf("unsigned: %v signed:%v", u, s)
	}
}
