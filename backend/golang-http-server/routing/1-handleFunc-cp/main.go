package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan http.HandleFunc

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	t := time.Now()
	writer.Write([]byte(fmt.Sprintf("%s, %d %s %d", t.Weekday(), t.Day(), t.Month(), t.Year())))
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	name := r.URL.Query().Get("name")
	if name == "" {
		w.Write([]byte("Hello there"))
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func main() {
	http.HandleFunc("/time", TimeHandler)
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
