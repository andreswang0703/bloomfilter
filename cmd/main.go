package main

import (
	"bloomFilter/pkg/bloomfilter"
	"bloomFilter/pkg/parser"
	"flag"
	"fmt"
	"runtime"
)

func main() {
	n := flag.Int("n", 5000, "Number of elements expected to store in the Bloom Filter.")
	p := flag.Float64("p", 0.99, "False positive possibility")
	path := flag.String("path", "./data/testInput.txt", "Path for the input strings")

	flag.Parse()

	bl := bloomfilter.Factory{}.Build(*n, *p)

	hashSet := make(map[string]bool)
	strings, err := parser.Parser{}.Parse(*path)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to parse file at %s: %w", *path, err))
	}

	bl.Feed(strings...)

	// test positive
	testStr := []string{"asd", "dfsasas", "dssa", "asdfa", "asfas", "sdvcxz", "svewqe", "sa", "cxv", "ewrwfv", "vasaeew", "cva", "qwer", "vaeewq", "sdffcc", "wev", "hello", "bye", "dkse"}
	for _, s := range testStr {
		seen := bl.Check(s)
		fmt.Println("seen ", testStr, "? - ", seen)
	}

	fmt.Println("------------------------------------------------------------------------")

	// test negative
	testStr2 := []string{"aaa", "bbb", "ccccc", "ddd", "a", "jkl"}
	for _, s := range testStr2 {
		seen := bl.Check(s)
		fmt.Println("seen ", testStr, "? - ", seen)
	}
	var bloomFilterMem runtime.MemStats
	runtime.ReadMemStats(&bloomFilterMem)

	fmt.Printf("Allocated memory (HeapAlloc): %d bytes\n", bloomFilterMem.HeapAlloc)
	fmt.Printf("Total memory allocated (Alloc): %d bytes\n", bloomFilterMem.Alloc)
	fmt.Printf("Memory allocations (Mallocs): %d\n", bloomFilterMem.Mallocs)
	fmt.Printf("Memory frees (Frees): %d\n", bloomFilterMem.Frees)

	for _, s := range strings {
		hashSet[s] = true
	}

}
