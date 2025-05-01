package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// Create an external process of `date` command
	dateCmd := exec.Command("date")

	// Run the process
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	// Print an output
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Error handling
	_, err = exec.Command("date", "-x").Output() // Invalid option
	if err != nil {
		switch e := err.(type) {
		case *exec.Error: // Error while execution
			fmt.Println("failed executing:", err)
		case *exec.ExitError: // Erro by non-zero return code
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	// Input/output pipe
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	// Input via the pipe
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	// Output via the pipe
	grepBytes, _ := io.ReadAll(grepOut)
	// Wait finish of the process
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Spawn a full command with a string by `bash -c`
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
