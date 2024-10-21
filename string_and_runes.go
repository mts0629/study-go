package main

import "fmt"

func main() {
	s := "あいうえお"

	fmt.Println(s)

	fmt.Println("Byte expression")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Length:", len(s))

	fmt.Println("Unicode expression")
	for _, runeVal := range s {
		fmt.Printf("%#U ", runeVal)
	}
	fmt.Println()
}
