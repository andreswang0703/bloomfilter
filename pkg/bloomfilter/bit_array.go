package bloomfilter

import "fmt"

type BitArray struct {
	m    int
	data []uint64
}

func (b *BitArray) init(m int) {
	b.m = m
	size := (m + 63) / 64
	b.data = make([]uint64, size)
	fmt.Println("created an uint array of size", size, "with total bit size", 64*size)
}

// Flip the bits based on the indexes
func (b *BitArray) Flip(indexes []uint64) {
	for _, index := range indexes {
		bitIdx := index % uint64(b.m)
		arrIdx := bitIdx / 64      // which uint64 in the array
		bitPosition := bitIdx % 64 // bit position within that uint64

		b.data[arrIdx] |= 1 << bitPosition // left shift 1 by bitPosition
	}
}

// Check if all the given indexes are flipped
func (b *BitArray) Check(indexes []uint64) bool {
	for _, index := range indexes {
		bitIdx := index % uint64(b.m)
		arrIdx := bitIdx / 64
		bitPosition := bitIdx % 64
		bitFlipped := b.checkBit(arrIdx, bitPosition)
		if !bitFlipped {
			return false
		}
	}
	return true
}

// check the bit of number at index arrIdx and at its bitPosition is flipped or not
func (b *BitArray) checkBit(arrIdx uint64, bitPosition uint64) bool {
	mask := uint64(1) << bitPosition
	return b.data[arrIdx]&mask != 0
}
