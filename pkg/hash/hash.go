package hash

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

func (a *AllHashesImpl) Init(k int) {
	var hashFunctions []HashFunction
	for i := 0; i < k; i++ {
		hashFunctions = append(hashFunctions, BuildHashFunction(i))
	}
	a.hashFunctions = hashFunctions
}

// GetIndexes returns the indexes of all the k hash functions' output with string as input
func (a *AllHashesImpl) GetIndexes(s string) []uint64 {
	var indexes []uint64
	for _, hashFunction := range a.hashFunctions {
		index := hashFunction.Hash(s)
		indexes = append(indexes, index)
	}
	return indexes
}
