//go:build !no_exit

package main

import (
	"fmt"
	"os"
)

func init() {
	registerCommand(&Command{
		Name:        "exit",
		Description: "Exit interactive mode (alias: quit)",
		Handler:     handleExit,
	})
}

func handleExit(args []string) {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

