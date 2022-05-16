package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah http.ServeMux yang memiliki dua route
// Route pertama yaitu "/time" yang menghandle TimeHandler
// Route kedua yaitu "/hello" yang menghandle SayHelloHandler

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Fprintf(w, "%s, %d %s %d", t.Weekday(), t.Day(), t.Month(), t.Year())
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello there")
		fmt.Fprint(w, "Hello, Roger!")
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())

	return mux

}
