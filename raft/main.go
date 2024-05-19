package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var servers []*server = []*server{
	{id: 0, currentTerm: 0, state: Follower},
	{id: 1, currentTerm: 0, state: Follower},
	{id: 2, currentTerm: 0, state: Follower},
	{id: 3, currentTerm: 0, state: Follower},
	{id: 4, currentTerm: 0, state: Follower},
}

type server struct {
	id            int
	currentTerm   int
	votedFor      int
	votesReceived int

	state string
	mu    sync.Mutex
}

const (
	Follower  string = "follower"
	Leader    string = "leader"
	Candidate string = "candidate"
)

func (s *server) run() {
	for {
		switch s.state {
		case Leader:
			s.leader()
		case Follower:
			s.follower()
		case Candidate:
			s.candidate()
		}
	}
}

func (s *server) follower() {
	log.Printf("server%d is a follower\n", s.id)
	timeout := time.Duration(rand.Intn(300)+150) * time.Millisecond
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			s.state = Candidate
			s.mu.Unlock()
			return // IMPORTANT
		}
	}
}

func (s *server) leader() {
	log.Printf("server%d is the leader for term %d\n", s.id, s.currentTerm)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	s.mu.Lock()
	defer s.mu.Unlock()
	i := 0
	for i < 5 {
		select {
		case <-ticker.C:
			s.sendHeartBeats()
			i++
		}
	}
}

// send heart beats to all your followers
func (s *server) sendHeartBeats() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		if i != s.id {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				log.Printf("Leader server%d sending heartbeat to server%d\n", s.id, i)
			}(i)
		}
	}
	wg.Wait()
}

func (s *server) candidate() {
	log.Printf("server%d is now a candidate and attempting an election for term %d\n", s.id, s.currentTerm+1)

	s.mu.Lock()
	s.currentTerm++
	s.votedFor = s.id
	s.votesReceived = 1 // Reset vote to 1 for each term
	s.mu.Unlock()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		if i != s.id {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if s.requestVote(i) {
					s.mu.Lock()
					s.votesReceived++
					s.mu.Unlock()
				}
			}(i)
		}
	}

	wg.Wait()

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.votesReceived > 2 {
		log.Printf("server%d WON the election and is now the leader for term %d\n", s.id, s.currentTerm)
		s.state = Leader
	} else {
		s.state = Follower
	}

}

func (s *server) requestVote(i int) bool {
	voter := servers[i]
	voter.mu.Lock()
	voter.currentTerm++
	voter.mu.Unlock()

	return rand.Intn(2) == 0
}

func main() {

	for _, srv := range servers {
		go srv.run()
	}

	select {}
}
