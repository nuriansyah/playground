package main

import (
	"fmt"
	"log"
	"net/http"
)

// Dari contoh yang telah diberikan, lengkapi handler dibawah ini.
// Lakukan pengecekan terhadap request method, jika request method tidak sama dengan GET maka return method not allowed.
// Selain itu gunakan variabel data dan err, jika err terdapat error, maka handler akan return internal server error.
// Jika panjang data dari data sama dengan 0, maka return not found.

var data []string
var err error

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	}
	fmt.Fprintln(w, data)

}

func main() {

	// Menambahkan handler pada objek server
	log.Println("Server run on port 8080")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
