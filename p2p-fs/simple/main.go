package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// Peer represents a peer in the BitTorrent network
type Peer struct {
	ID   int
	IP   string
	Port int
}

// Chunk represents a chunk of the file
type Chunk struct {
	Index   int
	Data    []byte
	Hash    [32]byte
	PeerIDs []int
}

var peers = []Peer{
	{ID: 0, IP: "127.0.0.1", Port: 80},
	{ID: 1, IP: "0.0.0.0", Port: 8080},
}

var chunks = []Chunk{
	{Index: 0, Data: []byte("chunk1"), Hash: sha256.Sum256([]byte("chunk1")), PeerIDs: []int{0}},
	{Index: 1, Data: []byte("chunk2"), Hash: sha256.Sum256([]byte("chunk2")), PeerIDs: []int{1}},
	{Index: 2, Data: []byte("chunk3"), Hash: sha256.Sum256([]byte("chunk3")), PeerIDs: []int{1}},
	{Index: 3, Data: []byte("chunk4"), Hash: sha256.Sum256([]byte("chunk4")), PeerIDs: []int{0}},
}

// PollChunkFromPeer polls a chunk from a peer
func PollChunkFromPeer(wg *sync.WaitGroup, peerID, chunkIdx int, chunkCh chan<- []byte) {
	defer wg.Done()
	fmt.Printf("Polling chunk:%d from peer%d\n", chunkIdx, peerID)

	chunk := chunks[chunkIdx]
	peer := peers[peerID]

	for _, id := range chunk.PeerIDs {
		if id == peer.ID {
			chunkCh <- chunk.Data
			return
		}
	}

	fmt.Printf("chunk%d is not available from peer%d\n", chunkIdx, peerID)
}

// VerifyIntegrity verifies the integrity of a chunk
func VerifyIntegrity(chunk Chunk) bool {
	return sha256.Sum256(chunk.Data) == chunk.Hash
}

func main() {
	n := len(chunks)
	chunkCh := make(chan []byte, n) // Buffered channel to avoid blocking
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		chunk := chunks[i]
		peerID := chunk.PeerIDs[0]
		go PollChunkFromPeer(&wg, peerID, chunk.Index, chunkCh)
	}

	go func() {
		wg.Wait()
		close(chunkCh)
	}()

	// wg.Wait()
	// close(chunkCh)

	for data := range chunkCh {
		fmt.Printf("Received data %s\n", string(data))
	}
}
