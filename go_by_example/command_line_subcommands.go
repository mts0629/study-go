package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Declare a subcommand `foo`
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// Specific flags for `foo`: `-enable`, `-name`
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Declare a subcommand `bar`
	// and a specific flag: `-level`
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Exit if subcommands are missing
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	// Access flags and trailing positional arguments for each subcommand
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("	enable:", *fooEnable)
		fmt.Println("	name:", *fooName)
		fmt.Println("	tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("	level:", *barLevel)
		fmt.Println("	tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
