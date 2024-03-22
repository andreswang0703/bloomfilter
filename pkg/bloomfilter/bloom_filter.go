package bloomfilter

import "bloomFilter/pkg/hash"

type BloomFilter interface {
	Feed(str ...string)
	Check(str string) bool
}

type BloomFilterImpl struct {
	options   Option
	bitArray  BitArray
	allHashes hash.AllHashesImpl
}

// Feed the bloom filter with strings
func (b *BloomFilterImpl) Feed(str ...string) {

	// todo: potentially use goroutine here?
	for _, s := range str {
		indexes := b.allHashes.GetIndexes(s)
		b.bitArray.Flip(indexes)
	}
}

// Check if the given string is already seen by bloom filter
func (b *BloomFilterImpl) Check(str string) bool {
	indexes := b.allHashes.GetIndexes(str)
	return b.bitArray.Check(indexes)
}

// init initialize bit array based on n
func (b *BloomFilterImpl) init() {
	m := b.options.m
	bitArray := BitArray{}
	bitArray.init(m)
	b.bitArray = bitArray
	b.allHashes.Init()
}
