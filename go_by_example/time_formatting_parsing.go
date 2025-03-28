package main

import (
	"fmt"
	"time"
)

func main() {
	// Alias for print
	p := fmt.Println

	t := time.Now()
	p(t)
	// RFC3339: Date and Time on Internet: Timestamps
	p(t.Format(time.RFC3339))

	// Parse time
	t1, _ := time.Parse(
		time.RFC3339,
		"2025-03-28T23:30:15+09:00")
	p(t1)

	// Example-based layouts
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM") // Parse by example layout
	p(t2)

	// String formatting
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Error on malformed input
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
