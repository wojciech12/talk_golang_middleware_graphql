package main

import (
	"io"
	"log"
	"net/http"
)

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Middleware before")
		next.ServeHTTP(w, r)
		log.Print("Middleware after")
	})
}

func main() {
	hello := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!")
	})
	err := http.ListenAndServe(":8080", exampleMiddleware(hello))
	log.Fatal(err)
}
