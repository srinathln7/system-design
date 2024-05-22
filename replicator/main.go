package main

import (
	"log"
	"sync"
)

var counter int

// Read only
type Follower struct {
	id int
	ch chan int
}

// Write/Read
type Master struct {
	mu sync.Mutex

	followersChan []chan int
}

func (m *Master) WriteUpdate(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Master writing update")

	for i := 0; i < 100; i++ {
		m.mu.Lock()
		counter++
		m.mu.Unlock()

		for i, fChan := range m.followersChan {
			log.Printf("Sending update to follower%d => Update counter value to %d\n", i, counter)
			fChan <- counter
		}

	}

	// Close the channel once communication is done
	for _, fChan := range m.followersChan {
		close(fChan)
	}

}

func (f *Follower) recvUpdate(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Follower receiving update")

	for value := range f.ch {
		log.Printf("Follower%d received update from master. Counter value updated to %d \n", f.id, value)
	}

}

func main() {

	f1 := &Follower{ch: make(chan int)}
	f2 := &Follower{id: 1, ch: make(chan int)}
	f3 := &Follower{id: 2, ch: make(chan int)}

	m := &Master{followersChan: make([]chan int, 0)}
	m.followersChan = append(m.followersChan, f1.ch, f2.ch, f3.ch)

	var wg sync.WaitGroup

	wg.Add(4)
	go m.WriteUpdate(&wg)
	go f1.recvUpdate(&wg)
	go f2.recvUpdate(&wg)
	go f3.recvUpdate(&wg)

	wg.Wait()
}
