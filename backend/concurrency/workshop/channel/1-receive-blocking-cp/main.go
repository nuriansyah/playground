package main

import "fmt"

func receiveBlock(output chan int) {
	c := make(chan int)
	result := 0
	go func() {
		fmt.Println("send to channel c")
		//kirim 1 ke channel c
		c <- 1
	}()

	//result menerima data dari channel c
	result = <-c
	output <- result
	fmt.Println(c) //agar variabel c digunakan
}
