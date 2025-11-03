//go:build !no_ls

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
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
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ls: cannot access '%s': %v\n", dir, err)
		os.Exit(1)
		return
	}
	fmt.Println("Name  Size  ModTime  Perms  Owner  Group")
	fmt.Println("----------------------------------------")
	for _, file := range files {
		name := file.Name()
		info, err := file.Info()
		if err != nil {
			fmt.Printf("%s: (error reading info)\n", file.Name())
			continue
		}
		size := info.Size()
		modTime := info.ModTime()
		perms := info.Mode()
		owner := info.Sys().(*syscall.Stat_t).Uid
		group := info.Sys().(*syscall.Stat_t).Gid
		// Mark directories with a trailing /
		if file.IsDir() {
			name = filepath.Join(name, "")
			fmt.Printf("%s/  %d  %s  %s  %d  %d\n", name, size, modTime.Format("2006-01-02 15:04:05"), perms.String(), owner, group)
		} else {
			fmt.Printf("%s  %d  %s  %s  %d  %d\n", name, size, modTime.Format("2006-01-02 15:04:05"), perms.String(), owner, group)

		}

	}
}
