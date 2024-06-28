package main

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

type bloomfilter struct {
	bits      []byte
	hashFuncs []func([]byte) int
}

// newOptimalBloomFilter creates a new Bloom filter with optimal parameters.
func newOptimalBloomFilter(n int, p float64) bloomfilter {
	m, k := calcOptimalParams(n, p) // Calculate bit size (m) and number of hash functions (k)
	size := (m + 7) / 8             // Calculate byte size
	hashFuncs := make([]func([]byte) int, k)
	for i := 0; i < k; i++ {
		hashFuncs[i] = makeHashFunc(i, m) // Pass bit size to makeHashFunc
	}

	return bloomfilter{
		bits:      make([]byte, size), // Initialize bit array
		hashFuncs: hashFuncs,          // Assign hash functions
	}
}

// calcOptimalParams calculates the optimal size of the bit array (m) and number of hash functions (k).
func calcOptimalParams(n int, p float64) (m, k int) {
	m = int(math.Ceil(-float64(n) * math.Log(p) / (math.Ln2 * math.Ln2))) // Calculate bit size
	k = int(math.Round(float64(m) * math.Ln2 / float64(n)))               // Calculate number of hash functions
	return m, k
}

// makeHashFunc creates a hash function with a given seed and size.
func makeHashFunc(seed, size int) func([]byte) int {
	return func(data []byte) int {
		h1, h2 := murmur3.Sum128WithSeed(data, uint32(seed))
		return int((h1 + h2) % uint64(size)) // Use bit size for modulo operation
	}
}

// Add adds an element to the Bloom filter.
func (bf bloomfilter) Add(data []byte) {
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)            // Get the bit index from hash function
		bf.bits[idx/8] |= 1 << (idx % 8) // Set the corresponding bit
	}
}

// Contains checks if an element is in the Bloom filter.
func (bf bloomfilter) Contains(data []byte) bool {
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)                 // Get the bit index from hash function
		if bf.bits[idx/8]&(1<<(idx%8)) == 0 { // Check if the corresponding bit is set
			return false
		}
	}
	return true // All bits are set, element might be in the set
}

func main() {
	n := 1000 // Expected number of elements
	p := 0.01 // Desired false positive probability

	obf := newOptimalBloomFilter(n, p)

	// Add elements to the Bloom filter
	elements := [][]byte{
		[]byte("apple"),
		[]byte("banana"),
		[]byte("cherry"),
	}

	for _, elem := range elements {
		obf.Add(elem)
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
		if obf.Contains(elem) {
			fmt.Printf("%s might be in the set\n", elem)
		} else {
			fmt.Printf("%s is definitely not in the set\n", elem)
		}
	}
}
