package bloomfilter

import (
	"fmt"
	"math"
)

const maxM = 81920 // maximum array size of 100KB to bits (10 * 1024 * 8)
const minM = 8192  // minimum array size of 1KB to bits (1024 * 8)
const maxK = 3     // maximum hash function amount

type Option struct {
	n int     // # elements
	p float64 // desired false positive probability
	k int     // optimal # hash functions, capped at maxK
	m int     // optimal bit array size, range [minM, maxM]
}

// SetN set the number of items to build the Bloom filter.
func (option *Option) SetN(n int) {
	option.n = n
	option.setOptimizedK()
}

// SetP Set the false positive probability of Bloom filter.
func (option *Option) SetP(p float64) {
	if p <= 0 || p >= 1 {
		return // Invalid p value
	}
	option.p = p
	option.setOptimizedM()
}

func (option *Option) GetM() int {
	return option.m
}

func (option *Option) GetN() int {
	return option.n
}

// set the optimal m based on p and n
func (option *Option) setOptimizedM() {
	m := int(-float64(option.n) * math.Log(option.p) / (math.Log(2) * math.Log(2)))
	option.m = int(math.Max(math.Min(float64(m), float64(maxM)), minM))
	option.setOptimizedK()
	fmt.Println("optimized m to", m)
}

// set the optimal k based on m and n
func (option *Option) setOptimizedK() {
	n := option.n
	m := option.m
	k := float64(m) / float64(n) * math.Log(2)
	option.k = int(math.Min(math.Ceil(k), float64(maxK)))
	fmt.Println("optimized k to", k)
}
