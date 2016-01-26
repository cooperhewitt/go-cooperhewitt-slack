package main

import (
	"io"
	"net/http"
	"os"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello slack!")
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", hello)
	http.ListenAndServe(":" + port, nil)
}
