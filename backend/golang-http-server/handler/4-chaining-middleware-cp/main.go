package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan chaining middleware.
// Lengkapi middleware AdminMiddleware dan RequestMethodGetMiddleware.
// Pada AdminMiddleware, lakukan pengecekan terhadap header "role". Jika nilai header "role" tidak sama dengan "ADMIN", maka return error HTTP Unauthorized.
// Pada RequestMethodGetMiddleware, lakukan pengecekan terhadap request method. Jika request method selain GET, maka return error HTTP Status Not Allowed.

func TimeHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		output := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		fmt.Fprint(writer, output)
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("role") != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}) // TODO: replace this
}

func main() {
	handler := TimeHandler()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
