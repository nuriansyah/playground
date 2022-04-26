package main

import "fmt"

func squareWorker(input <-chan int, output chan<- int) {
	//lakukan for range loop
	// TODO: answer here
	for i := range input {
		output <- i
	}
}

func receiver(output chan<- int) {
	input := make(chan int) // mengirim 0-10 ke squareWorker
	go squareWorker(input, output)
	for i := 0; i < 10; i++ {
		//kirim nilai i ke channel
		// TODO: answer here
		input <- i * i
	}
}

func main() {
	output := make(chan int)
	go receiver(output)
	for i := 0; i < 10; i++ {
		fmt.Println(<-output)
	}
}
