//go:build !no_ls

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	registerCommand(&Command{
		Name:        "ls",
		Description: "List directory contents",
		Handler:     handleLs,
	})
}

// handleLs implements the 'ls' command in-process (no forking).
func handleLs(args []string) {
	var dir string
	if len(args) > 0 {
		dir = args[0]
	} else {
		dir = "."
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ls: cannot access '%s': %v\n", dir, err)
		os.Exit(1)
		return
	}
	for _, file := range files {
		name := file.Name()
		// Mark directories with a trailing /
		if file.IsDir() {
			name = filepath.Join(name, "")
			fmt.Printf("%s/\n", name)
		} else {
			fmt.Println(name)
		}
	}
}
