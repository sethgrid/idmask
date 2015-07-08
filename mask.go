package idmask

var TotalIDs uint = 150

func Mask(ids []uint) int {
	var mask int
	for _, id := range ids {
		mask = setBit(mask, id)
	}
	return mask
}

func Unmask(n int) []uint {
	var ids []uint

	// not using 0 because this use case is for ids, and id 0 is just silly.
	for i := uint(1); i <= TotalIDs; i++ {
		if hasBit(n, i) {
			ids = append(ids, i)
		}
	}
	return ids
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
