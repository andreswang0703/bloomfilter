package bloomfilter

type BloomFilter interface {
	Feed(str ...string) bool
	Check(str string) bool
}

type BitArray struct {
}

type BloomFilterImpl struct {
	options  BloomFilterOption
	bitArray BitArray
}

func (b *BloomFilterImpl) Feed(str ...string) bool {
	return false
}

func (b *BloomFilterImpl) Check(str string) bool {
	return false
}

func (b *BloomFilterImpl) init() {
	bitArray := BitArray{}
	b.bitArray = bitArray
}
