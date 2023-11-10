package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var err error

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stdout, "Use ./pipe '<command 1> | <command2>'")
		os.Exit(0)
	}

	command := strings.Split(os.Args[1], "|")

	command1 := strings.Fields(command[0])
	command2 := strings.Fields(command[1])

	cmd1 := exec.Command(command1[0], command1[1:]...)
	cmd2 := exec.Command(command2[0], command2[1:]...)

	reader, writer := io.Pipe()
	cmd1.Stdout = writer
	cmd2.Stdin = reader

	if cmd1.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if cmd2.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if cmd1.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	writer.Close()

	if cmd2.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

}
