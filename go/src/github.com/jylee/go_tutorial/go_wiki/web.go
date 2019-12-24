package main

import "net/http"

import "fmt"

import "log"

// http.ResponseWriter -> assembles the HTTP server response -> We send data to the HTTP client
// http.Request -> data structure that represent the client HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h3>Hi there, I love %s!</h3></html>", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler) // HTTP package to handle ALL REQUEST to the web root ("/") with handler

	// http.HandleFunc CALLS http.ListenAndServe -> Listen only on port 8080
	// log.Fatal function calls os.Exit(1) after writing log message
	log.Fatal(http.ListenAndServe(":8080", nil))
}
