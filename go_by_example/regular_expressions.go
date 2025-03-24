package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// Pattern match of string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Compile regular expression
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Methods
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach"))

	fmt.Println("idx:", r.FindStringIndex("peace punch")) // Get indices of the first match

	fmt.Println(r.FindStringSubmatch("peach punch")) // Get whole-pattern matches and the submathces
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // Get all matches
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2)) // Get 2 patterns in maximum

	fmt.Println(r.Match([]byte("peach"))) // Byte match

	r = regexp.MustCompile("p([a-z]+)ch") // regexp.Compile variant, which panics instead of error
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // Replace matches

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) // Transform matches with a given function
	fmt.Println(string(out))
}
