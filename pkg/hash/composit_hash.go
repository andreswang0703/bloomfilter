package hash

import (
	"github.com/spaolacci/murmur3"
	"github.com/zhenjl/cityhash"
)

type CompositeHash struct {
	index int
}

func (m *CompositeHash) Hash(input string) uint64 {
	data := []byte(input)
	mHash := murmur3.Sum64(data)
	cHash := cityhash.CityHash64(data, uint32(len(data)))
	return mHash + (1<<m.index)*cHash
}

// BuildHashFunction builds a composite hash function out of murmur hash and city hash.
// H(x, i) = murmurHash(x) + 2^i * cityHash(x)
func BuildHashFunction(index int) HashFunction {
	return &CompositeHash{index: index}
}
