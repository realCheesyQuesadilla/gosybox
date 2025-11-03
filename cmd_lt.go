//go:build !no_ls

package main

import (
	"fmt"
	"os"
	"sort"
)

func init() {
	registerCommand(&Command{
		Name:        "lt",
		Description: "List directory contents and order by modification time",
		Handler:     handleLt,
	})
}

// handleLs implements the 'ls' command in-process (no forking).
func handleLt(args []string) {
	var dir string
	if len(args) > 0 {
		dir = args[0]
	} else {
		dir = "."
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lt: cannot access '%s': %v\n", dir, err)
		os.Exit(1)
		return
	}
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.ModTime().Before(infoJ.ModTime())
	})
	for _, file := range files {
		name := file.Name()
		// Mark directories with a trailing /
		if file.IsDir() {
			info, err := file.Info()
			if err != nil {
				fmt.Printf("%s/ (error reading info)\n", file.Name())
			} else {
				fmt.Printf("%s/  %s\n", name, info.ModTime().Format("2006-01-02 15:04:05"))
			}
		} else {
			fmt.Println(name)
		}
	}
}
