package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Construct a file path
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Path is normalized
	fmt.Println(filepath.Join("dir//", "filename"))        // "dir/filename"
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // "dir1/filename"

	// Get a directory
	fmt.Println("Dir(p):", filepath.Dir(p))
	// Get a basename (filename)
	fmt.Println("Base(p):", filepath.Base(p))

	// Check whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Get an extension
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Trim an extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Get a relative path from a base to a target
	rel, err := filepath.Rel("a/b", "a/b/t/file") // "t/file"
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file") // "../c/t/file"
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
