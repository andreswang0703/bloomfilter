package bloomfilter

import "math"

const maxM = 81920 // maximum array size of 10KB to bits (10 * 1024 * 8)
const maxK = 3     // maximum hash function amount

type BloomFilterOption struct {
	n int     // # elements
	p float64 // desired false positive probability
	k int     // optimal # hash functions, capped at maxK
	m int     // optimal bit array size, capped at maxM
}

// SetN set the number of items to build the Bloom filter.
func (option *BloomFilterOption) SetN(n int) {
	option.n = n
	option.setOptimizedK()
}

// SetP Set the false positive probability of Bloom filter.
func (option *BloomFilterOption) SetP(p float64) {
	if p <= 0 || p >= 1 {
		return // Invalid p value
	}
	option.p = p
	option.setOptimizedM()
}

func (option *BloomFilterOption) GetM() int {
	return option.m
}

func (option *BloomFilterOption) GetN() int {
	return option.n
}

// set the optimal m based on p and n
func (option *BloomFilterOption) setOptimizedM() {
	m := int(-float64(option.n) * math.Log(option.p) / (math.Log(2) * math.Log(2)))
	option.m = int(math.Min(float64(m), float64(maxM)))
	option.setOptimizedK()
}

// set the optimal k based on m and n
func (option *BloomFilterOption) setOptimizedK() {
	n := option.n
	m := option.m
	k := float64(m) / float64(n) * math.Log(2)
	option.k = int(math.Min(math.Ceil(k), float64(maxK)))
}
