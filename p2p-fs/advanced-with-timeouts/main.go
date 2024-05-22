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
	cond sync.Cond

	downloadedChunks int
	chunkSet         map[int]bool
	chunkCh          chan Chunk
}

func initDownloader() *Downloader {
	dl := &Downloader{
		chunkSet: make(map[int]bool),
		chunkCh:  make(chan Chunk),
	}

	dl.cond = *sync.NewCond(&dl.mu)
	return dl
}

func main() {
	timeOut, totalChunks := 5, len(chunks)
	dl := initDownloader()
	dl.downloadChunksFromPeer(peers, totalChunks, timeOut)
}

func (dl *Downloader) downloadChunksFromPeer(peers []Peer, totalChunks, timeout int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Initiate parallel download process
	for _, peer := range peers {
		dl.wg.Add(1)
		go dl.pollChunksFromPeer(ctx, peer)
	}

	// Wait until the entire peers complete their processes
	go func() {
		dl.wg.Wait()
		close(dl.chunkCh)
	}()

	// Chunk Handler to verify the integrity and download process
	go func() {
		dl.mu.Lock()
		for chunk := range dl.chunkCh {
			if !dl.chunkSet[chunk.Index] {
				if dl.VerifyChunk(chunk) {
					dl.downloadedChunks++
					dl.chunkSet[chunk.Index] = true
					log.Printf("Chunk%d verified and downloaded successfully \n", chunk.Index)
					if dl.downloadedChunks == totalChunks {
						dl.cond.Broadcast()
					}
				} else {
					log.Printf("Verification failed for chunk %d \n", chunk.Index)
				}
			} else {
				log.Printf("Chunk%d already verified and downloaded \n", chunk.Index)
			}
		}

		dl.mu.Unlock()
	}()

	// Wait until the full chunks are downloaded
	dl.mu.Lock()
	for dl.downloadedChunks != totalChunks {
		dl.cond.Wait()
	}
	dl.mu.Unlock()

	log.Printf("Download process completed. Total chunks downloaded: %d\n", dl.downloadedChunks)
}

func (dl *Downloader) pollChunksFromPeer(ctx context.Context, peer Peer) {
	defer dl.wg.Done()
	log.Printf("Requesting chunks from peer %d\n", peer.id)

	retires := 3
	for attempt := 0; attempt < retires; attempt++ {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(1)*time.Second)
		defer cancel()

		chunks, err := getChunksFromPeer(ctx, peer)
		if err == nil {
			log.Printf("Successfully retrieved chunks from peer %d on attempt %d\n", peer.id, attempt+1)
			for _, chunk := range chunks {
				dl.chunkCh <- chunk
			}
			return
		}

		// Retry after delay
		log.Printf("Attempt %d failed due to error %s. Retrying in 500ms \n", attempt, ctx.Err())
		time.Sleep(500 * time.Millisecond)
	}

}

func getChunksFromPeer(ctx context.Context, peer Peer) ([]Chunk, error) {

	for {

		// Simulate network delay
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

func (dl *Downloader) VerifyChunk(chunk Chunk) bool {
	return chunk.Hash == sha256.Sum256(chunk.Data)
}
