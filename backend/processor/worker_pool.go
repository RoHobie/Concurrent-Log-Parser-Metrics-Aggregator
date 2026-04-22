package processor

import (
	"sync"

	"Concurrent-Log-Parser-Metrics-Aggregator/backend/storage"
)

type WorkerPool struct {
	storage storage.Storage
	jobs    chan Job
	wg      sync.WaitGroup

	mu      sync.Mutex
	counter int
}

func NewWorkerPool(s storage.Storage, numWorkers int) *WorkerPool {
	wp := &WorkerPool{
		storage: s,
		jobs:    make(chan Job, 1000),
	}

	for i := 0; i < numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}

	return wp
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()

	for job := range wp.jobs {

		_ = wp.storage.Set(job.Key, job.Data)

		wp.mu.Lock()
		wp.counter++
		wp.mu.Unlock()
	}
}

func (wp *WorkerPool) Submit(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
}

func (wp *WorkerPool) Count() int {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	return wp.counter
}