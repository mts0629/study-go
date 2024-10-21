package main

import (
	"fmt"
	// "unicode/utf8"
)

func main() {
	s := "ABCあいうえお" // string: immutable array of UTF-8 code points in byte
	fmt.Println(s)

	fmt.Println("Byte expression") // byte: 8-bit integer
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Length (byte):", len(s))

	fmt.Println("Rune expression") // rune: 32-bit integer
	for _, runeVal := range s {
		fmt.Printf("%c (%#U) ", runeVal, runeVal)
	}
	fmt.Println()
	fmt.Println("Length (rune):", len([]rune(s)))
	// fmt.Println("Length (rune):", utf8.RuneCountInString(s)) // Count in runes
}
