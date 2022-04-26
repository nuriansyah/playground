package main

import (
	"fmt"
	"sync"
)

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {

	// TODO: answer here
	c := make(chan int) // buffer channel

	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//kirim 1 ke channel

			// TODO: answer here
			c <- 1
			wg.Done()

		}()
	}
	//mengubah nilai count menggunakan data dari channel

	// TODO: answer here
	go func() {
		for {
			count += <-c //mengambil data dari channel
		}
	}()
	wg.Wait() // menunggu seluruh goroutine selesai berjalan
	output <- count
}

func main() {
	output := make(chan int)
	go counter(output)
	fmt.Println(<-output)
}
