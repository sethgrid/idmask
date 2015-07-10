package idmask

import "fmt"

// set to this value to to have some cap. could set this to max uint64-1
var MaxID uint64 = 256

func Mask(ids []uint64) []uint64 {
	buckets := make([]uint64, uint64(MaxID)/64+1)
	for _, id := range ids {
		buckets[id/64] = setBit(buckets[id/64], id%64)
	}
	return buckets
}

func Unmask(buckets []uint64) []uint64 {
	var ids []uint64

	for i, idPool := range buckets {
		for j := uint64(0); j < 64; j++ {
			if hasBit(idPool, j) {
				ids = append(ids, j+uint64(i*64))
			}
		}
	}
	return ids
}

func setBit(n uint64, pos uint64) uint64 {
	n |= (1 << pos)
	return n
}

func clearBit(n uint64, pos uint64) uint64 {
	mask := uint64(^(1 << pos))
	n &= mask
	return n
}

func hasBit(n uint64, pos uint64) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func debug(format string, args ...interface{}) {
	debug := false
	if debug {
		fmt.Printf(format+"\n", args...)
	}
}
