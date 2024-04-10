package bloomfilter

type Factory struct{}

func (b Factory) Build(n int, p float64) (BloomFilter, error) {
	options := Option{}
	options.SetN(n)
	err := options.SetP(p)
	if err != nil {
		return nil, err
	}

	bloomFilter := &BloomFilterImpl{}
	bloomFilter.options = options

	bloomFilter.init()
	return bloomFilter, nil
}
