package main

import (
	"bloomFilter/pkg/bloomfilter"
	"bloomFilter/pkg/hash"
	"fmt"

	"flag"
)

func main() {
	options := bloomfilter.BloomFilterOption{}
	n := flag.Int("n", 1000, "Number of elements expected to store in the Bloom Filter.")
	p := flag.Float64("p", 0.99, "False positive possibility")

	flag.Parse()

	options.SetN(*n)
	options.SetP(*p)

	fmt.Println("Bloom Filter Options:", *n, *p)
}

func performHash(hash hash.Hash_Interface, input string) uint64 {
	return hash.Hash(input)
}

func getInt(hash uint64, size uint64) uint64 {
	return hash % size
}
