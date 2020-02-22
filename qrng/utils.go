package qrng

import (
	"fmt"
	"math"
)

func checkBlockSize(blockArray []string, blockSize int) error {
	for block, i := range blockArray {
		if blockSize != len(string(block)) {
			return fmt.Errorf(BlockSizeMismatchError, i, len(string(block)), blockSize)
		}
	}
	return nil
}

func ConcatInt8(a, b int8) int16 {
	return (int16(a) << 8) + int16(b)
}
func ConcatUint8(a, b uint8) uint16 {
	return (uint16(a) << 8) + uint16(b)
}

func ConcatInt16(a, b int16) int32 {
	return (int32(a) << 16) + int32(b)
}
func ConcatUint16(a, b uint16) uint32 {
	return (uint32(a) << 16) + uint32(b)
}

func ConcatInt32(a, b int32) int64 {
	return (int64(a) << 32) + int64(b)
}
func ConcatUint32(a, b uint32) uint64 {
	return (uint64(a) << 32) + uint64(b)
}

func ConcatUintFloat32(a, b uint16) float32 {
	c := (uint32(a) << 16) + uint32(b)
	return math.Float32frombits(c)
}

func ConcatUintFloat64(a, b, c, d uint16) float64 {
	e := (uint64(a) << 48) + (uint64(b) << 32) + (uint64(c) << 16) + uint64(d)
	return math.Float64frombits(e)
}

func ConcatFloat32(a, b float32) float64 {
	ab := math.Float32bits(a)
	bb := math.Float32bits(b)
	c := (uint64(ab) << 32) + uint64(bb)
	return math.Float64frombits(c)
}
