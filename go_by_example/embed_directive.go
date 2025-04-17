package main

import (
	"embed"
)

// Embed the contents of the file `work/single_file.txt` into variables
//
//go:embed work/single_file.txt
var fileString string

//go:embed work/single_file.txt
var fileByte []byte

// Embed multiple files/folders
//
//go:embed work/single_file.txt
//go:embed work/*.hash
var folder embed.FS

func main() {
	// Print contents
	print(fileString)
	print(string(fileByte))

	// Retrieve each file from the embedded folder
	content1, _ := folder.ReadFile("work/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("work/file2.hash")
	print(string(content2))

	content3, _ := folder.ReadFile("work/single_file.txt")
	print(string(content3))
}
