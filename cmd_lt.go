//go:build !no_lt

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"syscall"
)

func init() {
	registerCommand(&Command{
		Name:        "lt",
		Description: "List directory contents and order by modification time",
		Handler:     handleLt,
	})
}

// handleLt implements the 'lt' command in-process (no forking).
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
