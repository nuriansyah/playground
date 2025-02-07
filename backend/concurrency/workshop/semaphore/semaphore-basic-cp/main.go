package main

import (
	"fmt"
	"sync"
	"time"
)

type worker struct {
	work        func(string)
	currentWork int
	maxWork     int
	sem         Semaphore
	mu          *sync.Mutex
}

func (w *worker) doWork(input string) {
	for i := 0; i < 100; i++ {
		// TODO: answer here
		//w.work(input)
		w.sem.Acquire()
		go func() {
			// TODO: answer here
			defer w.sem.Release()
			//w.sem.Acquire()

			w.mu.Lock()
			w.currentWork++
			if w.maxWork < w.currentWork {
				w.maxWork = w.currentWork
			}
			w.mu.Unlock()
			time.Sleep(1 * time.Millisecond)

			w.work(input)

			w.mu.Lock()
			w.currentWork--
			w.mu.Unlock()
		}()

	}
}

func newWorker(work func(string), semToken int) *worker {
	return &worker{work: work, currentWork: 0, maxWork: 0, sem: newSemaphore(semToken), mu: &sync.Mutex{}}
}

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

func newSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}

func main() {
	sayHello := func(name string) {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}
	worker := newWorker(sayHello, 5)
	worker.doWork("friend")
	time.Sleep(10 * time.Millisecond)

}
