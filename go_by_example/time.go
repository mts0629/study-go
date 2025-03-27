package main

import (
	"fmt"
	"time"
)

func main() {
	// Alias for print
	p := fmt.Println

	// Current time
	now := time.Now()
	p(now)

	// Build a time struct
	then := time.Date(
		2025, 3, 27, 23, 30, 45, 651387237, time.Local)
	p(then)

	// Time components
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday()) // Monday-Sunday

	// Compare two times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Duration
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add/subtract duration
	p(then.Add(diff))
	p(then.Add(-diff))
}
