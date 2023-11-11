package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./pipe '<command1> | <command2>'")
		os.Exit(1)
	}

	commands := strings.Split(os.Args[1], "|")
	if len(commands) != 2 {
		fmt.Fprintln(os.Stderr, "Error: You must provide exactly two commands separated by '|'")
		os.Exit(1)
	}

	command1 := strings.Fields(commands[0])
	command2 := strings.Fields(commands[1])

	cmd1 := exec.Command(command1[0], command1[1:]...)
	cmd2 := exec.Command(command2[0], command2[1:]...)

	pipe, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating stdout pipe for command 1:", err)
		os.Exit(1)
	}

	cmd2.Stdin = pipe
	cmd2.Stdout = os.Stdout

	if err := cmd1.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting command 1:", err)
		os.Exit(1)
	}

	if err := cmd2.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting command 2:", err)
		os.Exit(1)
	}

	if err := cmd1.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for command 1 to finish:", err)
		os.Exit(1)
	}

	if err := cmd2.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for command 2 to finish:", err)
		os.Exit(1)
	}
}
