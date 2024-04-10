package main

import (
	"bloomFilter/pkg/bloomfilter"
	"bloomFilter/pkg/parser"
	"flag"
	"fmt"
	"runtime"
)

func main() {

	//n := flag.Int("n", 20, "Number of elements expected to store in the Bloom Filter.")
	p := flag.Float64("p", 0.01, "False positive possibility")
	path := flag.String("path", "./data/urlList.txt", "Path for the input strings")

	flag.Parse()

	strings, err := parser.Parser{}.Parse(*path, "\\n")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to parse file at %s: %w", *path, err))
	}

	bl, bfBuildError := bloomfilter.Factory{}.Build(len(strings), *p)
	if bfBuildError != nil {
		fmt.Println(fmt.Errorf("failed to build bloom filter %s", bfBuildError))
		return
	}

	bl.Feed(strings...)

	// test positive
	testStr := []string{"https://www.picture.sample.info/", "http://www.smell.sample.info/", "http://authority.sample.info/flower/cattle.php", "http://www.sample.com/#team"}
	for _, s := range testStr {
		seen := bl.Check(s)
		fmt.Println("seen ", s, "? - ", seen)
	}

	fmt.Println("------------------------------------------------------------------------")

	// test negative
	testStr2 := []string{"aaa", "bbb", "ccccc", "ddd", "a", "jkl"}
	for _, s := range testStr2 {
		seen := bl.Check(s)
		fmt.Println("seen ", s, "? - ", seen)
	}
	var bloomFilterMem runtime.MemStats
	runtime.ReadMemStats(&bloomFilterMem)

	fmt.Printf("Allocated memory (HeapAlloc): %d bytes\n", bloomFilterMem.HeapAlloc)
	fmt.Printf("Total memory allocated (Alloc): %d bytes\n", bloomFilterMem.Alloc)
	fmt.Printf("Memory allocations (Mallocs): %d\n", bloomFilterMem.Mallocs)
	fmt.Printf("Memory frees (Frees): %d\n", bloomFilterMem.Frees)
}
