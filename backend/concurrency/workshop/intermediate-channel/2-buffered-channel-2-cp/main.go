package main

import (
	"fmt"
	"time"
)

var person = "andi"
var names = []string{"budi", "toni", "adi", "ado", "alif", "yudi"}

//mengembalikan string, dimana `name` menyapa semua `names`
func greetAll(person string, names []string, output chan<- string) {
	for _, name := range names {
		output <- fmt.Sprintf("%s say hello to %s", person, name)
	}
	fmt.Println("selesai mengirim")

}

// buat size buffered channel sesuai jumlah names
func testBufferedChannel(result chan<- string) {
	output := make(chan string, len(names))

	go greetAll(person, names, output)
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 6; i++ {
		greeting := <-output
		result <- greeting
	}
}

//goroutine greetAll dapat mengirim ke goroutine testBufferedChannel walaupun channel belum siap menerima

func main() {
	result := make(chan string, 6)
	go testBufferedChannel(result)
	for i := 0; i < 6; i++ {
		fmt.Println(<-result)
	}
}
