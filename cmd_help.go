//go:build !no_help

package main

import "fmt"

func init() {
	registerCommand(&Command{
		Name:        "help",
		Description: "Show this help message",
		Handler:     handleHelp,
	})
}

func handleHelp(args []string) {
	if len(args) == 0 {
		fmt.Println("I am gosybox, a replacement for busybox written in Go.")
		fmt.Println("Available commands:")
		for _, cmd := range listCommands() {
			fmt.Printf("  %-8s %s\n", cmd.Name, cmd.Description)
		}
	}
}
