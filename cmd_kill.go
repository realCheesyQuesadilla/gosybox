//go:build !no_kill

package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func init() {
	registerCommand(&Command{
		Name:        "kill",
		Description: "Send a signal to a process",
		Handler:     handleKill,
	})
}

func handleKill(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: kill [-s signal] pid")
		fmt.Println("  signal:  Signal to send (default is SIGTERM)")
		fmt.Println("  pid:      Process ID to send signal to")
		fmt.Println("\nAvailable signals:")
		printAvailableSignals()
		os.Exit(1)
		return
	}

	// SIGTERM will be the default
	var signal syscall.Signal = syscall.SIGTERM 
	var pid int

	if len(args) == 1 {
		// Just a PID
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("kill: '%s': Invalid PID\n", pid)
			os.Exit(1)
			return
		}
	} else if len(args) >= 2 {
		// Signal and PID
		var err error
		signal, err = parseSignal(args[0])
		if err != nil {
			fmt.Printf("kill: invalid signal: %s\n", err)
			os.Exit(1)
			return
		}
		pid, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("kill: '%s': Invalid PID\n", pid)
			os.Exit(1)
			return
		}
	} else {
		fmt.Println("Usage: kill [-s signal] pid")
		os.Exit(1)
		return
	}

	// Send the signal
	if err := sendSignal(pid, signal); err != nil {
		fmt.Printf("kill: %v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Signal %d sent to process %d\n", signal, pid)
}

func parseSignal(signalName string) (syscall.Signal, error) {
	// Try parsing as number first
	if num, err := strconv.Atoi(signalName); err == nil {
		return syscall.Signal(num), nil
	}

	// Try parsing as signal name
	switch signalName {
	case "TERM", "KILL", "INT", "HUP", "QUIT", "USR1", "USR2", "CHLD", "CONT", "STOP", "TSTP", "TTIN", "TTOUT", "URG", "ALRM", "VTALRM", "PROF", "WINCH", "IO", "POWER":
		return syscall.Signal(syscall.Signal(syscall.SIGTERM)), nil
	default:
		return 0, fmt.Errorf("invalid signal name")
	}
}

// sendSignal sends a signal to a process
func sendSignal(pid int, signal syscall.Signal) error {
	err := syscall.Kill(pid, signal)
	if err != nil {
		if err.Error() == "no such process" || err.Error() == "ESRCH" {
			return fmt.Errorf("process %d not found", pid)
		}
		return err
	}

	return nil
}

// printAvailableSignals prints available signals
func printAvailableSignals() {
	fmt.Println("  TERM (15)  - Terminate process")
	fmt.Println("  KILL (9)   - Force kill process")
	fmt.Println("  INT (2)    - Interrupt (SIGINT)")
	fmt.Println("  HUP (1)    - Hang up")
	fmt.Println("  QUIT (3)   - Quit")
	fmt.Println("  CHLD (20)  - Child stopped or gone")
	fmt.Println("  CONT (19)  - Continue")
	fmt.Println("  STOP (18)  - Stop")
	fmt.Println("  TSTP (21)  - Stop job control")
	fmt.Println("  TTIN (24)  - Read from tty")
	fmt.Println("  TTOUT (25) - Write to tty")
	fmt.Println("  URG (26)   - Urgent SO band")
	fmt.Println("  ALRM (27)  - Alarm clock")
	fmt.Println("  VTALRM (28)- Virtual alarm clock")
	fmt.Println("  PROF (29)  - Profiling alarm clock")
	fmt.Println("  WINCH (30) - Window resize")
	fmt.Println("  IO (23)    - I/O error")
	fmt.Println("  POWER (6)  - Power event")
}
