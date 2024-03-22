package bloomfilter

type BloomFilterFactory interface {
	build(n int, p float64) BloomFilter
}

type BloomFilterFactoryImpl struct{}

func (b BloomFilterFactoryImpl) build(n int, p float64) BloomFilter {
	options := BloomFilterOption{}
	options.SetN(n)
	options.SetP(p)

	bloomFilter := &BloomFilterImpl{}
	bloomFilter.options = options

	bloomFilter.init()
	return bloomFilter
}
