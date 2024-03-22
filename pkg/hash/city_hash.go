package hash

import (
	"github.com/zhenjl/cityhash"
)

type CityHash struct{}

func (c CityHash) Hash(input string) uint64 {
	data := []byte(input)
	hash := cityhash.CityHash64(data, uint32(len(data)))
	return hash
}
