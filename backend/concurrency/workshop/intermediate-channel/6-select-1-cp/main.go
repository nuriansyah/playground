package main

import (
	"fmt"
	"time"
)

type character struct {
	name            string
	activity        string
	defaultActivity string
}

//pada program ini dibuat character yang bernama john dengan defaulActivity mengawasi area dan activity kosong
//kita dapat mengubah activity dari john dengan cara mengirim ke channel
//untuk bergerak kita kirim ke movementInput, lalu untuk menyerang kita kirim ke attackInput

func (c *character) awake(movementInput, attackInput chan string) {
	for {
		//melakukan select
		//ketika menerima dari attackInput maka jalankan
		//fmt.Printf("%s melakukan serangan %s\n", c.name, c.activity)
		//ketika menerima dari movementInput maka jalankan
		//fmt.Printf("%s bergerak ke %s\n", c.name, c.activity)
		// TODO: answer here
		select {
		case movement := <-movementInput:
			c.activity = movement
			fmt.Printf("%s bergerak ke %s\n", c.name, c.activity)
		case attack := <-attackInput:
			c.activity = attack
			fmt.Printf("%s melakukan serangan %s\n", c.name, c.activity)
		}

	}
}

func main() {
	john := character{
		name:            "john",
		defaultActivity: "mengawasi area",
		activity:        "",
	}

	movementInput := make(chan string, 4)
	attackInput := make(chan string, 3)

	go john.awake(movementInput, attackInput)
	movementInput <- "atas"
	movementInput <- "bawah"
	attackInput <- "A"
	movementInput <- "kiri"
	attackInput <- "B"
	movementInput <- "kanan"
	attackInput <- "C"
	time.Sleep(100 * time.Millisecond)
}
