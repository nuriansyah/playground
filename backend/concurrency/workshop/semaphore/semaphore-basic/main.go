package main

import (
	"fmt"
	"sync"
	"time"
)

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

//kita bisa lihat hanya ada 10 goroutine yang berjalan dalam satu waktu
func main() {
	var wg sync.WaitGroup
	semaphore := newSemaphore(10) // kita ingin hanya ada 10 akses dalam satu waktu
	for i := 1; i <= 30; i++ {
		wg.Add(1)
		semaphore.Acquire()
		go func(i int) {
			defer semaphore.Release()
			longRunningProcess(i)
			wg.Done()
		}(i)
	}
	//kapan terjadi blocking pada program ini ?
	// TODO: answer here
	// time.Sleep(10 * time.Millisecond)
	wg.Wait()

}

func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(1 * time.Second) // melakukan sesuatu
}
