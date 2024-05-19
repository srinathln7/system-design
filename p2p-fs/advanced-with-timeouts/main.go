package main

import (
	"context"
	"crypto/sha256"
	"log"
	"sync"
	"time"
)

type Peer struct {
	id     int
	chunks []Chunk
}

// Chunk represents a chunk of the file
type Chunk struct {
	Index int
	Data  []byte
	Hash  [32]byte
}

var (
	chunks = []Chunk{
		{Index: 0, Data: []byte("chunk1"), Hash: sha256.Sum256([]byte("chunk1"))},
		{Index: 1, Data: []byte("chunk2"), Hash: sha256.Sum256([]byte("chunk2"))},
		{Index: 2, Data: []byte("chunk3"), Hash: sha256.Sum256([]byte("chunk3"))},
		{Index: 3, Data: []byte("chunk4"), Hash: sha256.Sum256([]byte("chunk4"))},
		{Index: 4, Data: []byte("chunk5"), Hash: sha256.Sum256([]byte("chunk5"))},
		{Index: 5, Data: []byte("chunk6"), Hash: sha256.Sum256([]byte("chunk6"))},
		{Index: 6, Data: []byte("chunk7"), Hash: sha256.Sum256([]byte("chunk7"))},
	}

	peers = []Peer{
		{id: 0, chunks: []Chunk{chunks[0], chunks[1], chunks[2]}},
		{id: 1, chunks: []Chunk{chunks[2], chunks[3], chunks[0]}},
		{id: 2, chunks: []Chunk{chunks[4], chunks[5], chunks[6]}},
	}
)

type Downloader struct {
	wg   sync.WaitGroup
	mu   sync.Mutex
	cond *sync.Cond

	chunksDownloaded int
	chunkSet         map[int]bool
	chunkCh          chan Chunk
}

func downloadChunksFromPeers(peers []Peer, totalChunks int) {
	log.Println("Starting download process...")

	dl := &Downloader{
		chunkSet: make(map[int]bool),
		chunkCh:  make(chan Chunk),
	}
	dl.cond = sync.NewCond(&dl.mu)

	timeout := 1
	for _, peer := range peers {
		dl.wg.Add(1)
		go dl.reqChunkFromPeer(peer, timeout)
	}

	go func() {
		dl.wg.Wait()
		close(dl.chunkCh)
	}()

	go func() {
		for chunk := range dl.chunkCh {
			dl.mu.Lock()

			// If the chunk has not yet been downloaded and verified
			if !dl.chunkSet[chunk.Index] {
				if dl.VerifyIntegrity(chunk) {
					log.Printf("Chunk %d integrity verified and added.\n", chunk.Index)
					dl.chunksDownloaded++
					dl.chunkSet[chunk.Index] = true
					if dl.chunksDownloaded == totalChunks {
						dl.cond.Broadcast() // Notify all waiting goroutines
					}
				} else {
					log.Printf("Chunk %d integrity verification failed.\n", chunk.Index)
				}
			}

			dl.mu.Unlock()
		}
	}()

	dl.mu.Lock()
	for dl.chunksDownloaded != totalChunks {
		dl.cond.Wait()
	}
	dl.mu.Unlock()

	log.Printf("Download process completed. Total chunks downloaded: %d\n", dl.chunksDownloaded)

}

func (dl *Downloader) VerifyIntegrity(chunk Chunk) bool {
	return chunk.Hash == sha256.Sum256(chunk.Data)
}

func (dl *Downloader) reqChunkFromPeer(peer Peer, timeout int) {
	defer dl.wg.Done()

	log.Printf("Requesting chunks from peer %d\n", peer.id)

	retries := 3
	for attempt := 0; attempt < retries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()

		chunks, err := pollChunksFromPeer(ctx, peer)
		if err == nil {
			log.Printf("Successfully retrieved chunks from peer %d on attempt %d\n", peer.id, attempt+1)
			for _, chunk := range chunks {
				dl.chunkCh <- chunk
			}
			return
		}

		log.Printf("Attempt %d to retrieve chunks from peer %d failed: %v\n", attempt+1, peer.id, err)
		time.Sleep(500 * time.Millisecond)
	}

	log.Printf("Failed to retrieve chunks from peer %d after %d attempts\n", peer.id, retries)
}

func pollChunksFromPeer(ctx context.Context, peer Peer) (chunks []Chunk, err error) {
	log.Printf("Polling chunks from peer %d\n", peer.id)

	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Printf("Polling from peer %d timed out\n", peer.id)
			return nil, ctx.Err()
		default:
			log.Printf("Successfully polled chunks from peer %d\n", peer.id)
			return peer.chunks, nil
		}
	}
}

func main() {
	n := len(chunks)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Println("Download process timeout")
			return
		default:
			downloadChunksFromPeers(peers, n)
			return
		}
	}
}
