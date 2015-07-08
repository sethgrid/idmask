package idmask

import "fmt"

var MaxID uint64 = 150

func Mask(ids []uint64) []uint64 {
	mask := make([]uint64, uint64(MaxID)/64+1)
	for _, id := range ids {
		mask[id/64] = setBit(mask[id/64], id%64)
	}
	return mask
}

func Unmask(n []uint64) []uint64 {
	var ids []uint64

	// not using 0 because this use case is for ids, and id 0 is just silly.
	for i, idPool := range n {
		for j := uint64(0); j < MaxID; j++ {
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
