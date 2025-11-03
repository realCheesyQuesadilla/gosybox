package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check for interactive mode flag or no arguments
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "-i") {
		runInteractiveMode()
		return
	}

	// Command mode: execute a single command
	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := getCommand(cmdName)
	if cmd == nil {
		fmt.Printf("Unknown command: %s\n", cmdName)
		fmt.Println("Use 'help' to see available commands.")
		fmt.Println("Run without arguments or with '-i' to enter interactive mode.")
		os.Exit(1)
		return
	}

	cmd.Handler(args)
}

// runInteractiveMode starts an interactive shell loop
func runInteractiveMode() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("gosybox interactive mode")
	fmt.Println("Type 'help' for available commands, 'exit' or 'quit' to exit")
	fmt.Print("gosybox> ")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			fmt.Print("gosybox> ")
			continue
		}

		// Parse command and arguments
		parts := strings.Fields(line)
		cmdName := parts[0]
		args := parts[1:]

		// Handle exit/quit commands
		if cmdName == "exit" || cmdName == "quit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
			return
		}

		// Get and execute command
		cmd := getCommand(cmdName)
		if cmd == nil {
			fmt.Printf("Unknown command: %s\n", cmdName)
			fmt.Println("Type 'help' to see available commands.")
		} else {
			cmd.Handler(args)
		}

		fmt.Print("gosybox> ")
	}

	// Handle EOF (Ctrl+D)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
		return
	}

	// Exit cleanly on EOF
	fmt.Println() // New line after EOF
	os.Exit(0)
}
