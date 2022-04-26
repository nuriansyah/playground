package main

import (
	"fmt"
	"time"
)

//numberWorker digunakan untuk mengirim data ke channel output
func numberWorker(output chan int) {
	for i := 0; i < 100; i++ {
		//kirim ke channel
		// TODO: answer here
		output <- i //kirim data ke channel
		fmt.Println("kirim data ke ", i)

	}
	close(output)
	//pastikan tulisan ini berhasil ditampilkan
	fmt.Println("finished sending data. channel closed")
}

func receiver(result chan int) {
	//buat channel dengan ukuran yang sesuai dengan data-
	//yang akan dikirim numWorker
	output := make(chan int, len(result))
	fmt.Println(len(result))
	// TODO: replace this
	sum := 0

	go numberWorker(output)
	time.Sleep(100 * time.Millisecond)
	for number := range output {
		square := number * number
		sum += square
		fmt.Println("Number: ", number)
	}
	result <- sum
}

func main() {
	result := make(chan int)
	go receiver(result)
	fmt.Println("result:", <-result)
}
