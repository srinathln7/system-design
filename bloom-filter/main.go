package main

import (
	"crypto/sha256"
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	bitset    []bool
	size      int
	hashFuncs []func(data []byte) int
}

func NewBloomFilter(n int, p float64) *BloomFilter {
	size, numHashFuncs := OptimalParams(n, p)
	bf := &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
	}

	// Add hash functions
	for i := 0; i < numHashFuncs; i++ {
		bf.hashFuncs = append(bf.hashFuncs, bf.makeHashFunc(uint32(i)))
	}

	return bf
}

// Calculate optimal number of hash functions and size for the filter
func OptimalParams(n int, p float64) (int, int) {
	m := int(math.Ceil(float64(-n) * math.Log(p) / (math.Ln2 * math.Ln2)))
	k := int(math.Round(math.Ln2 * float64(m) / float64(n)))
	return m, k
}

// makeHashFunc generates a hash function based on the seed
func (bf *BloomFilter) makeHashFunc(seed uint32) func(data []byte) int {
	return func(data []byte) int {
		// Combine SHA-256 hash with MurmurHash3
		h := sha256.New()
		h.Write(data)
		sum := h.Sum(nil)

		hash := murmur3.Sum32WithSeed(sum, seed)
		return int(hash % uint32(bf.size))
	}
}

// Add inserts an element into the Bloom filter
func (bf *BloomFilter) Add(data []byte) {
	for _, hashFunc := range bf.hashFuncs {
		index := hashFunc(data)
		bf.bitset[index] = true
	}
}

// Contains checks if an element is in the Bloom filter
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hashFunc := range bf.hashFuncs {
		index := hashFunc(data)
		if !bf.bitset[index] {
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
