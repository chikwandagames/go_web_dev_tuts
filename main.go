package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

func main() {
	// HandleFunc registers the path "/hello" to be handled by helloHandleFunc()
	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)

	// ListenAndServer listens on the TCP network address addr and then calls Serve
	// with handler to handle requests on incoming connections
	http.ListenAndServe("localhost:"+port, nil)

}

// Handler function that response to client http requests
func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Gophers, %s", r.URL.Path)
}
