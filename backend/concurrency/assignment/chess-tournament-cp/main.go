package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}
func main() {
	startTournament()
}

func playMatch(i int) {
	fmt.Printf("match #%d started\n", i+1)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("match #%d finished\n", i+1)
}

func startTournament() {
	for i := 0; i < 10; i++ {
		go playMatch(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("match #%d started\n", i+1)
		go time.Sleep(150 * time.Millisecond)
		fmt.Printf("match #%d finished\n", i+1)
		go time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("all the match finished. Total time needed:", time.Since(start))
}
