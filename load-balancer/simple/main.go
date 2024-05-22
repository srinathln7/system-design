package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	serverIDCounter int
)

type Server struct {
	id int
}

type LoadBalancer struct {
	done chan interface{}
	mu   sync.RWMutex
	wg   sync.WaitGroup

	reqs []int
}

func newLoadBalancer() *LoadBalancer {
	return &LoadBalancer{done: make(chan interface{}), reqs: make([]int, 0)}
}

func (lb *LoadBalancer) addRequest(reqID int) {
	defer lb.wg.Done()

	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.reqs = append(lb.reqs, reqID)
	log.Printf("added request id %d to the load balancer", reqID)
}

// newServer creates a new server instance with an auto-incremented ID
func newServer() *Server {
	serverIDCounter++
	log.Printf("created server%d", serverIDCounter)
	return &Server{id: serverIDCounter}
}

func (s *Server) doWork(reqID int, lb *LoadBalancer) <-chan string {
	log.Printf("request %d sent to server %d", reqID, s.id)

	// Simulate computing the response
	resCh := make(chan string)
	response := "response for request" + strconv.Itoa(reqID)
	go func() {
		defer close(resCh)
		select {
		case <-lb.done:
			log.Fatal("load balancer terminated abruptly")
		case resCh <- fmt.Sprintf("server %d computed %s", s.id, response):
		}
	}()

	return resCh
}

func main() {
	// Create a load balancer
	lb := newLoadBalancer()
	defer close(lb.done)

	// Create 5-virtual servers
	var num_of_servers int = 5
	servers := make([]*Server, num_of_servers)
	for i := 0; i < num_of_servers; i++ {
		servers[i] = newServer()
	}

	// Simulate adding a few requests to the load balancer
	var num_of_reqs int = 10
	for i := 1; i <= num_of_reqs; i++ {
		lb.wg.Add(1)
		go lb.addRequest(i)
	}

	lb.wg.Wait()
	log.Printf("added %d requests to the load balancer", num_of_reqs)

	// Implement a simple round-robin algorithm
	var res string
	for reqID := 1; reqID <= num_of_reqs; reqID++ {
		select {
		case res = <-servers[reqID%num_of_servers].doWork(reqID, lb):
			log.Println(res)
		case <-time.After(1 * time.Second):
			log.Println("request timed out")
		}
	}

	log.Printf("load balancer finished processing %d number of requests", num_of_reqs)
}
