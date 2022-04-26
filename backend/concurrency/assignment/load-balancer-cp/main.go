package main

import (
	"fmt"
	"sync"
)

const nWorkers = 3
const nRequesters = 5

type Work struct {
	data string
}

var mu = &sync.Mutex{}

var data = map[string]bool{}

func worker(in <-chan *Work, out chan<- *Work, number int) {
	for w := range in {
		w.data = fmt.Sprintf("worker %d accepted %s", number, w.data) // mengubah data
		out <- w                                                      // mengirimkan data ke channel out
	}
}

func createRequest(in chan<- *Work, number int) {
	for {
		data := fmt.Sprintf("request from client %d", number) // membuat data
		in <- &Work{data}                                     // memasukkan data ke dalam channel in

	}

}

func receiver(out <-chan *Work) {
	for {
		//pada bagian ini masukkan data dari channel out
		//ke dalam map `data`
		//gunakan mutex untuk mengamankan penulisan ke map secara concurrent
		w := <-out          // memansukan channel out ke w
		mu.Lock()           // mutex lock
		data[w.data] = true // menambahkan data ke dalam map bernilai true
		mu.Unlock()         // mutex Unlock
	}

}
