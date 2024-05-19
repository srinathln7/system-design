package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Msg struct {
	content string
}

type Follower struct {
	id int
	ch chan Msg
}

type Master struct {
	id            int
	followersChan []chan Msg
}

func (m *Master) Send(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Println("master starting to send message")

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("master%d - version:%d \n", m.id, i)
		msg := Msg{content: text}

		for _, followerCh := range m.followersChan {

			followerCh <- msg

			// Simulate network delay
			time.Sleep(1 * time.Second)
		}
	}

	// Close all follower channels -> Otherwise we have a deadlock
	for _, followerChan := range m.followersChan {
		close(followerChan)
	}

}

func (f *Follower) receive(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	// IMPORTANT: This reads the msg only once and exits immediately
	// msg := <-f.ch
	// log.Printf("Follower %d received: %s\n", f.id, msg.content)

	for message := range f.ch {
		log.Printf("Follower %d received: %s\n", f.id, message.content)
	}

	log.Printf("Follower %d: Channel closed\n", f.id)
}

func main() {

	master := &Master{id: 0}

	follower1 := &Follower{id: 1, ch: make(chan Msg)}
	follower2 := &Follower{id: 2, ch: make(chan Msg)}

	master.followersChan = append(master.followersChan, follower1.ch, follower2.ch)

	var wg sync.WaitGroup

	wg.Add(3)

	go follower1.receive(&wg)
	go follower2.receive(&wg)
	go master.Send(&wg)

	wg.Wait()
}
