package main

import (
	"fmt"
	"net/http"
)

// Handler interface: response "hello\n"
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// Read HTTP request headers and echo them into the response body
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Regiser handlers on server routes
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Launch an http server (in localhost:8090)
	http.ListenAndServe(":8090", nil)
}
