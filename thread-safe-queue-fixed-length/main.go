package main

import (
	"log"
	"sync"
)

type Queue struct {
	vals           []int
	fixed_length   int
	total_capacity int

	mu   sync.Mutex
	cond sync.Cond
}

const (
	QUEUE_CAPACITY   int = 50
	QUEUE_FIXED_SIZE int = 5
)

func initQueue(fixed_size, capacity int) *Queue {
	queue := Queue{
		vals:           make([]int, 0, capacity),
		fixed_length:   fixed_size,
		total_capacity: capacity,

		mu: sync.Mutex{},
	}

	queue.cond = *sync.NewCond(&queue.mu)
	return &queue
}

func (q *Queue) enqueue(wg *sync.WaitGroup, thread_id, start, end int) {

	defer wg.Done()

	log.Printf("THREAD%d enqueuing values between %d and %d to the queue \n", thread_id, start, end)
	log.Printf("Current QUEUE %d \n", q.vals)
	for i := start; i <= end; i++ {
		q.mu.Lock()

		for len(q.vals) == q.fixed_length {
			q.cond.Wait()
		}

		log.Printf("Enqueuing value %d to the queue \n", i)
		q.vals = append(q.vals, i)

		// Schedule dequeue in the background
		if i != end {
			go q.dequeue()
		}

		q.mu.Unlock()
	}

}

func (q *Queue) dequeue() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.vals) > 0 {
		log.Printf("Dequeuing element %d from the queue \n", q.vals[0])
		q.vals = q.vals[1:]
		q.cond.Signal()
	}
}

func main() {

	queue := initQueue(QUEUE_FIXED_SIZE, QUEUE_CAPACITY)

	//  At any point in time , queue's size should not exceed queue's fixed size

	var wg sync.WaitGroup

	n := 5
	wg.Add(n)

	for i := 0; i < n; i++ {
		go queue.enqueue(&wg, i, (i*10)+1, (i+1)*10)
	}

	wg.Wait()

	log.Println("Queue now contains: ", queue.vals)
	log.Println("Does queue contain fixed size length?", len(queue.vals) == queue.fixed_length)
}
