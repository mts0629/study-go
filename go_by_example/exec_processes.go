package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Get an absolute path to the `ls` command
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Command line strings
	args := []string{"ls", "-a", "-l", "-h"}

	// Environment variables
	env := os.Environ()

	// `exec` system call,
	// replace the current process with a process of `ls -a -l -h`
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
