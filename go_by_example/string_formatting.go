package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	// Struct formatting
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)  // Value
	fmt.Printf("struct2: %+v\n", p) // With field name
	fmt.Printf("struct3: %#v\n", p) // Go syntax representation (snippet)

	// Type
	fmt.Printf("type: %T\n", p)
	// Boolean
	fmt.Printf("bool: %t\n", true)

	// Integer formatting
	fmt.Printf("int: %d\n", 123) // Decimal
	fmt.Printf("bin: %b\n", 14)  // Binary
	fmt.Printf("char: %c\n", 33) // ASCII character
	fmt.Printf("hex: %x\n", 456) // Hexadecimal

	// Float formatting
	fmt.Printf("float1: %f\n", 78.9)        // Decimal
	fmt.Printf("float2: %e\n", 123400000.0) // Scientific notation (e)
	fmt.Printf("float3: %E\n", 123400000.0) // Scientific notation (E)

	// String formatting
	fmt.Printf("str1: %s\n", "\"string\"") // Basic
	fmt.Printf("str2: %q\n", "\"string\"") // Double-quoted
	fmt.Printf("str3: %x\n", "hex this")   // Base-16 bytes

	// Pointer formatting
	fmt.Printf("pointer: %p\n", &p)

	// Number width and precision
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)         // Right-justified number
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)   // Right-justify and decimal precision in 2
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // Left-justified number
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")      // Right-justfied string
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")    // Left-justified string

	// Formatted string
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Format and print to stream
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
