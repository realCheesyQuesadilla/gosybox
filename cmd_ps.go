//go:build !no_ps

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func init() {
	registerCommand(&Command{
		Name:        "ps",
		Description: "Display process status",
		Handler:     handlePs,
	})
}

// handlePs implements the 'ps' command by reading from /proc
func handlePs(args []string) {
	procDir := "/proc"

	// Print header
	fmt.Printf("%-8s %-8s %s\n", "PID", "PPID", "CMD")
	fmt.Println(strings.Repeat("-", 60))

	// Read /proc directory
	entries, err := os.ReadDir(procDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ps: cannot access '%s': %v\n", procDir, err)
		os.Exit(1)
		return
	}

	// Iterate through entries
	for _, entry := range entries {
		// Check if entry is a directory and represents a PID (numeric)
		if !entry.IsDir() {
			continue
		}

		pidStr := entry.Name()
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			// Not a numeric directory, skip
			continue
		}

		// Read process information
		pidPath := filepath.Join(procDir, pidStr)

		// Read stat file for PID and PPID
		statPath := filepath.Join(pidPath, "stat")
		statData, err := os.ReadFile(statPath)
		if err != nil {
			// Process may have exited, skip
			continue
		}

		// Parse stat file (format: pid (comm) state ppid ...)
		statFields := strings.Fields(string(statData))
		if len(statFields) < 4 {
			continue
		}

		ppid, err := strconv.Atoi(statFields[3])
		if err != nil {
			continue
		}

		// Extract command name (between parentheses in stat)
		comm := ""
		statStr := string(statData)
		start := strings.Index(statStr, "(")
		end := strings.LastIndex(statStr, ")")
		if start != -1 && end != -1 && end > start {
			comm = statStr[start+1 : end]
		} else {
			// Fallback to comm file
			commPath := filepath.Join(pidPath, "comm")
			if commData, err := os.ReadFile(commPath); err == nil {
				comm = strings.TrimSpace(string(commData))
			} else {
				comm = "unknown"
			}
		}

		// Print process information
		fmt.Printf("%-8d %-8d %s\n", pid, ppid, comm)
	}
}
