package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := getCommand(cmdName)
	if cmd == nil {
		fmt.Printf("Unknown command: %s\n", cmdName)
		fmt.Println("Use 'help' to see available commands.")
		os.Exit(1)
		return
	}

	cmd.Handler(args)
}
