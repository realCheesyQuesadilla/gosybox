//go:build !no_exit

package main

import (
	"fmt"
	"os"
)

func init() {
	registerCommand(&Command{
		Name:        "pwd",
		Description: "Shows current path to working directory",
		Handler:     handlePwd,
	})
}

func handlePwd(args []string) {
	wd := os.Getenv("PWD")
	if wd == "" {
		fmt.Fprintf(os.Stderr, "pwd: cannot PWD environment")
		os.Exit(1)
	}

	fmt.Printf("%s\n", wd)
	os.Exit(0)
}
