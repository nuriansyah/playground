package main

import (
	"fmt"
	"net/http"
)

var (
	username = "aditira"
	password = "aditira123"
)

func main() {
	fmt.Println("Starting Server at port :8080")
	http.ListenAndServe(":8080", Routes())
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error parsing basic auth"))
			return
		}
		if u != username {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid username"}`))
			return
		}
		if p != password {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid password"}`))
			return
		}
		fmt.Printf("Username: %s\n", u)
		fmt.Printf("Password: %s\n", p)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "welcome to CAMP 2022!"}`))

	})

	return mux
}

// Encode auth aditira:aditira123 disini -> https://www.base64encode.org/

// $ curl -v -X POST http://localhost:8080/login -H "Authorization: Basic YWRpdGlyYTphZGl0aXJhMTIz"
