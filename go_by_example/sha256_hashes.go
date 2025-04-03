package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// New SHA256 hash
	h := sha256.New()

	// Write bytes and get its hash
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// SHA512
	s2 := "sha512 this string"
	h2 := sha512.New()
	h2.Write([]byte(s2))
	fmt.Println(s2)
	fmt.Printf("%x\n", h2.Sum(nil))
}
