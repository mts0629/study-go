package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	// Simple logging, preconfigured to output to stderr
	log.Println("standard logger")

	// Set format flags: date, time and microseconds
	// LstdFlags = Ldate | Ltime
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Print the file name and line which the log function is called
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Create a custom logger with a prefix "my:"
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Set a new prefix
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Custom output target for io.Writer
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("from buflog:", buf.String())

	// Structured log in JSON format
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")    // Value of "msg"
	myslog.Info("hello again", // Value of "msg"
		"key", "val", // Key-value pairs
		"age", 25)
}
