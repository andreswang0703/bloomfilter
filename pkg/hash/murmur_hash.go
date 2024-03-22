package hash

import (
	"github.com/spaolacci/murmur3"
)

type MurmurHash struct{}

func (m MurmurHash) Hash(input string) uint64 {
	data := []byte(input)
	return murmur3.Sum64(data)
}
