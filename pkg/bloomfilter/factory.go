package bloomfilter

type Factory struct{}

func (b Factory) Build(n int, p float64) BloomFilter {
	options := Option{}
	options.SetN(n)
	options.SetP(p)

	bloomFilter := &BloomFilterImpl{}
	bloomFilter.options = options

	bloomFilter.init()
	return bloomFilter
}
