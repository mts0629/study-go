package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create a new sub directory "subdir"
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Remove a whole directory tree (`rm -rf`)
	defer os.RemoveAll("subdir")

	// Create a new empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Create a hierarchy of directories (`mkdir -p`)
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// Get a list of directory by a slice of os.DirEntry
	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Change the current working directory (`cd`)
	err = os.Chdir("subdir/parent/child")
	check(err)

	// See the contents of a directory
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	// Visit a directory recursively
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// Callback function applied for each entry
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
