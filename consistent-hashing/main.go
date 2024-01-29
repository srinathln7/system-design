package main

import (
	"hash/fnv"
	"log"
	"sort"
	"sync"
)

const HASH_LENGTH = 2 ^ 32

// HashRing : Contains a set of nodes, replicas for each node, and slots.
type HashRing struct {
	mu      sync.RWMutex
	nodes   []int
	hashMap map[int]string
}

func newHashRing() *HashRing {
	return &HashRing{
		nodes:   make([]int, 0, HASH_LENGTH),
		hashMap: make(map[int]string),
	}
}

// hashKey hashes a key using the FNV-1a algorithm.
func hashKey(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32())
}

// @@Param : IP address of the node
func (hr *HashRing) AddNodeToRing(node string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	// Get the hash key for the nodeID
	nodeID := hashKey(node)
	if _, exists := hr.hashMap[nodeID]; exists {
		log.Fatal("specified node already exists")
		return
	}

	hr.nodes = append(hr.nodes, nodeID)
	hr.hashMap[nodeID] = node

	// Each nodeID represents a slot in the ring.
	// Sort the slots in an ascending order for efficient look up of data ownership.
	// Remember binary serach only runs in a sorted array
	sort.Ints(hr.nodes)
	log.Printf("succesfully added node %s with id %d to the ring", node, nodeID)
}

func (hr *HashRing) DeleteNodeFromRing(node string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	nodeID := hashKey(node)
	index := sort.SearchInts(hr.nodes, nodeID)
	if index == len(hr.nodes) {
		log.Fatal("specified node does not exist in the ring")
		return
	}

	var nodes []int
	if index != 0 {
		nodes = append(nodes, hr.nodes[0:index-1]...)
		nodes = append(nodes, hr.nodes[index+1:]...)
	} else {
		nodes = append(nodes, hr.nodes[1:]...)
	}

	hr.nodes = nodes
	delete(hr.hashMap, nodeID)

	log.Printf("succesfully deleted node %s with id %d from the ring", node, nodeID)
}

func (hr *HashRing) GetNodeForKey(key string) int {
	hr.mu.RLock()
	defer hr.mu.RUnlock()

	keyID := hashKey(key)

	// Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true,
	// assuming that on the range [0, n), f(i) == true implies f(i+1) == true. That is, Search requires that
	// f is false for some (possibly empty) prefix of the input range [0, n) and then true for the (possibly empty) remainder;
	// Search returns the first true index. If there is no such index, Search returns n. (Note that the "not found" return value is not -1 as in,
	//for instance, strings.Index.) Search calls f(i) only for i in the range [0, n).
	index := sort.Search(len(hr.nodes), func(i int) bool {
		return keyID <= hr.nodes[i]
	})

	// Circle back
	if index == len(hr.nodes) {
		index = 0
	}

	return hr.nodes[index]
}

func main() {
	// Create a new hash ring
	hashRing := newHashRing()

	// Add nodes to the hash ring
	hashRing.AddNodeToRing("nodeA")
	// hashRing.AddNodeToRing("nodeA")
	hashRing.AddNodeToRing("nodeB")
	hashRing.AddNodeToRing("nodeC")

	// Get the node responsible for a key
	key := "test_consistent_hasing"
	nodeID := hashRing.GetNodeForKey(key)
	log.Printf("node responsible for key '%s': %s\n", key, hashRing.hashMap[nodeID])

	// Delete a node from the hash ring
	// deletedNode := "noded"
	deletedNode := "nodeA"
	hashRing.DeleteNodeFromRing(deletedNode)

	// Get the node responsible for the key after deletion
	nodeAfterDeletion := hashRing.GetNodeForKey(key)
	log.Printf("node responsible for key '%s' after deleting '%s' is now node with %s\n", key, deletedNode, hashRing.hashMap[nodeAfterDeletion])
}

// Illustrate the use of sort.Search
// func guessGame() {
// 	var s string
// 	fmt.Printf("Pick an integer from 0 to 100.\n")
// 	answer := sort.Search(100, func(i int) bool {
// 		fmt.Printf("Is your number <= %d? ", i)
// 		fmt.Scanf("%s", &s)
// 		return s != "" && s[0] == 'y'
// 	})
// 	fmt.Printf("Your number is %d.\n", answer)
// }
