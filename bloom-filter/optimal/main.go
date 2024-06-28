package main

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

type bloomfilter struct {
	bits      []byte
	hashFuncs []func([]byte) uint32
}

func newOptimalBloomFilter(n int, p float64) bloomfilter {
	m, k := calcOptimalParams(n, p)
	size := (m + 7) / 8
	hashFuncs := make([]func([]byte) uint32, k)
	for i := 0; i < k; i++ {
		hashFuncs[i] = makeHashFunc(uint32(i), uint32(size))
	}

	return bloomfilter{
		bits:      make([]byte, size),
		hashFuncs: hashFuncs,
	}
}

func calcOptimalParams(n int, p float64) (m, k int) {
	m = int(math.Ceil(float64(-n)*math.Log(p)) / (math.Ln2 * math.Ln2))
	k = int(math.Round((float64(m) * math.Ln2) / float64(n)))
	return m, k
}

func makeHashFunc(seed, size uint32) func([]byte) uint32 {
	return func(data []byte) uint32 {
		h1, h2 := murmur3.Sum128WithSeed(data, uint32(seed))
		return uint32((h1 + h2) % uint64(size))
	}
}

func (bf bloomfilter) Add(data []byte) {
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)
		bf.bits[idx/8] |= 1 << (idx % 8)
	}
}

func (bf bloomfilter) Contains(data []byte) bool {
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)
		if bf.bits[idx/8]&(1<<(idx%8)) == 0 {
			return false
		}
	}

	return true
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
