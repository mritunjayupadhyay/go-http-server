package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHelloFunc(w http.ResponseWriter, r *http.Request) {
	// Check if path is correct
	if r.URL.Path != "/hello" {
		http.Error(w, "path not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Incorrect Method", 401)
		return
	}
	fmt.Fprintf(w, "hello from hello handler 2")
}

func main() {
	fmt.Printf("Hello world")
	http.HandleFunc("/hello", handleHelloFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
