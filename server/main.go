package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: /\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: /hello\n")
	hasName := r.URL.Query().Has("name")
	name := r.URL.Query().Get("name")
	
	if hasName {
		io.WriteString(w, fmt.Sprintf("Hi %s! Welcome!\n", name))
	} else {
		io.WriteString(w, "Hello!\n")
	}
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("server: %s \n", r.Method)
	fmt.Fprintf(w, `{"message": "hello"}`)
}

func main() {
	// Define main mux / server and handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/message", getMessage)

	// Run server
	fmt.Println("Server starting on :3333")
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		log.Fatal(err)
	}
}
