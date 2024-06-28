package main

import (
	"crypto/sha256"
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

type bloomfilter struct {
	bitSize   []bool
	hashFuncs []func([]byte) int
}

func NewBloomFilter(numofElements int, successRate float64) bloomfilter {

	size, numOfHashFuncs := calcOptimalParams(numofElements, successRate)
	bitSize := make([]bool, size)

	hashFuncs := make([]func([]byte) int, numOfHashFuncs)
	for i := 0; i < numOfHashFuncs; i++ {
		hashFuncs[i] = makeHashFunc(i, size)
	}

	return bloomfilter{
		bitSize:   bitSize,
		hashFuncs: hashFuncs,
	}
}

func calcOptimalParams(n int, p float64) (m int, k int) {

	// Calc. `m` and `k` based on the derived formula
	// m = - n * ln(p) /(ln(2))^2  => bit size
	// p = m/n * ln(2) => num of hash functions

	m = int(math.Ceil(-float64(n)*math.Log(p)) / (math.Ln2 * math.Ln2))
	k = int(math.Round(float64(m)*math.Ln2) / float64(n))

	return m, k
}

func makeHashFunc(seed, size int) func([]byte) int {
	return func(data []byte) int {
		h1 := sha256.New()
		h1.Write(data)
		sum := h1.Sum(nil)

		h2 := murmur3.Sum32WithSeed(sum, uint32(seed))
		return int(h2) % size
	}
}

func (bf bloomfilter) Add(data []byte) {
	// Spread across multiple hash functions to get uniform distribution of index
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)
		bf.bitSize[idx] = true
	}
}

func (bf bloomfilter) Contains(data []byte) bool {

	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)

		if !bf.bitSize[idx] {
			return false
		}
	}

	return true
}

func main() {
	n := 1000 // Expected number of elements
	p := 0.01 // Desired false positive probability

	bf := NewBloomFilter(n, p)

	// Add elements to the Bloom filter
	elements := [][]byte{
		[]byte("apple"),
		[]byte("banana"),
		[]byte("cherry"),
	}

	for _, elem := range elements {
		bf.Add(elem)

	}

	// Check if elements are in the Bloom filter
	testElements := [][]byte{
		[]byte("apple"),
		[]byte("banana"),
		[]byte("cherry"),
		[]byte("date"),
		[]byte("honey"),
		[]byte("sugar"),
		[]byte("sirup"),
	}

	for _, elem := range testElements {
		if bf.Contains(elem) {
			fmt.Printf("%s might be in the set\n", elem)
		} else {
			fmt.Printf("%s is definitely not in the set\n", elem)
		}
	}
}
