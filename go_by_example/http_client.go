package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Issue an HTTP GET request to "Go by Example"
	// Create an http.Client and call its Get method
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print a response status
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 line of the response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
