package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// URL includes:
	// scheme, authentication info, port, path, query params, query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // Scheme

	fmt.Println(u.User)            // All authentication info
	fmt.Println(u.User.Username()) // Username
	p, _ := u.User.Password()      // Password
	fmt.Println(p)

	fmt.Println(u.Host)                        // Hostname and port
	host, port, _ := net.SplitHostPort(u.Host) // Split
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)     // Path
	fmt.Println(u.Fragment) // Fragment

	fmt.Println(u.RawQuery)            // Query params
	m, _ := url.ParseQuery(u.RawQuery) // Parse to map from strings to slices of strings
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
