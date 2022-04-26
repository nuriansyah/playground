package main

import (
	"sync"
)

var mu sync.Mutex
var result int

//fungsi ini digunakan untuk menerima angka lalu mengembalikan hasil pangkat 2 angka tersebut
func squareWorker(workerInput <-chan int, workerOutput chan<- int) {
	for {
		num := <-workerInput      // menerima angka
		workerOutput <- num * num // mengirim hasil
	}
}

func createRequest(workerInput chan<- int, workerOutput <-chan int, wg *sync.WaitGroup) {
	for i := 1; i < 100; i++ {
		workerInput <- i
		go func(i int) {
			res := <-workerOutput
			mu.Lock()
			result += res
			mu.Unlock()

		}(i)
	}
	wg.Done()
}
