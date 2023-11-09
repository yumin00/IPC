package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stdout, "Use './pipe <command 1> | <command2>")
		os.Exit(0)
	}
}
