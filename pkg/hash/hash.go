package hash

import "fmt"

type HashFunction interface {
	Hash(input string) uint64
}

type AllHashes interface {
	GetIndexes(s string) []int
}

type AllHashesImpl struct {
	k             int
	hashFunctions []HashFunction
}

func (a *AllHashesImpl) Init() {
	// todo: create k hash functions
	var hashFunctions []HashFunction
	hashFunctions = append(hashFunctions, CityHash{})
	hashFunctions = append(hashFunctions, MurmurHash{})
	a.hashFunctions = hashFunctions
	fmt.Println("setting hash functions: ", a.hashFunctions[0], a.hashFunctions[1])
}

// GetIndexes returns the indexes of all the k hash functions' output with string as input
func (a *AllHashesImpl) GetIndexes(s string) []uint64 {
	var indexes []uint64
	for _, hashFunction := range a.hashFunctions {
		index := hashFunction.Hash(s)
		indexes = append(indexes, index)
	}
	if s == "hello" {
		fmt.Println("indexes for hello are", indexes)
	}
	return indexes
}
