package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	id int
}

type Worker struct {
	id     int
	taskCh chan Task
}

type LoadBalancer struct {
	wg sync.WaitGroup
	mu sync.Mutex

	workers    []*Worker
	nextWorker int
}

func newWorker(id int) *Worker {
	return &Worker{
		id:     id,
		taskCh: make(chan Task),
	}
}

func newLoadBalancer(numWorkers int) *LoadBalancer {
	lb := LoadBalancer{
		workers: make([]*Worker, numWorkers),
	}

	// Worker id starts from 1 to n
	for i := 0; i < numWorkers; i++ {
		lb.workers[i] = newWorker(i + 1)
	}

	return &lb
}

func (lb *LoadBalancer) start() {

	for _, worker := range lb.workers {
		lb.wg.Add(1)
		go worker.start(&lb.wg)
	}

	go func() {
		lb.wg.Wait()
	}()
}

func (w *Worker) start(wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range w.taskCh {
		fmt.Printf("Worker %d processing task %d\n", w.id, task.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate processing time
		fmt.Printf("Worker %d finished task %d\n", w.id, task.id)
	}
}

func (lb *LoadBalancer) DistributeTasks(task Task) {
	lb.mu.Lock()
	worker := lb.workers[lb.nextWorker]

	// Classic Round Robin Algorithm
	lb.nextWorker = (lb.nextWorker + 1) % len(lb.workers)
	lb.mu.Unlock()

	worker.taskCh <- task
}

func (lb *LoadBalancer) Stop() {
	for _, worker := range lb.workers {
		close(worker.taskCh)
	}
}

func main() {
	numWorkers, numTasks := 3, 10

	lb := newLoadBalancer(numWorkers)

	// Start the load balancer
	lb.start()

	for i := 0; i < numTasks; i++ {
		task := Task{id: i + 1}
		lb.DistributeTasks(task)
		time.Sleep(100 * time.Millisecond)
	}

	lb.Stop()

	time.Sleep(2 * time.Second)
	fmt.Println("All tasks processed. Exiting...")
}
