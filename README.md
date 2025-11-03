<img width="250" height="300" alt="image" src="https://github.com/user-attachments/assets/7688dc9a-f46e-41a1-803a-681b37113c64" />

# gosybox

A lightweight, modular BusyBox replacement written in Go. gosybox provides essential Unix utilities in a single binary, with the ability to customize which commands are included at compile time using Go build tags.

## Features

- **Single Binary**: All commands are compiled into one executable
- **Modular Design**: Select which commands to include at build time
- **No Forking**: All commands run in-process for better performance
- **Go Build Tags**: Easy compile-time customization
- **Self-Documenting**: Built-in help command lists available commands

## Quick Start

### Building

Build with all commands (default):
```bash
go build -o gosybox
```

### Usage

Run any command:
```bash
./gosybox <command> [args...]
```

Examples:
```bash
./gosybox ls
./gosybox ls /home
./gosybox help
```

## Customizing Builds

### Excluding Commands

Use Go build tags to exclude specific commands:

```bash
# Exclude ls command
go build -tags no_ls -o gosybox

# Exclude help command
go build -tags no_help -o gosybox

# Exclude multiple commands
go build -tags "no_ls no_help" -o gosybox
```

This allows you to create minimal builds with only the commands you need, reducing binary size.

### How Build Tags Work

Each command is in its own file with a build tag:
- `cmd_ls.go` - includes `ls` unless `no_ls` tag is set
- `cmd_help.go` - includes `help` unless `no_help` tag is set

Commands register themselves in `init()` functions, which run automatically when the package loads.

## Available Commands

- `ls` - List directory contents (marks directories with `/`)
- `help` - Show available commands and help information

## Adding New Commands

To add a new command, create a file `cmd_<name>.go`:

```go
//go:build !no_<name>

package main

import (
    "fmt"
    "os"
)

func init() {
    registerCommand(&Command{
        Name:        "<name>",
        Description: "Brief description of what this command does",
        Handler:     handle<Name>,
    })
}

func handle<Name>(args []string) {
    // Your command implementation here
    // args contains all command-line arguments
}
```

The command will automatically appear in the help output and be available at runtime.

## Project Structure

```
gosybox/
├── gosybox.go      # Main entry point and command dispatcher
├── commands.go     # Command registry system
├── cmd_ls.go      # ls command implementation
├── cmd_help.go    # help command implementation
└── go.mod         # Go module file
```

## Future Plans

### Core Utilities
- [ ] `cat` - Concatenate and display files
- [ ] `echo` - Display text
- [ ] `mkdir` - Create directories
- [ ] `rm` - Remove files and directories
- [ ] `cp` - Copy files and directories
- [ ] `mv` - Move/rename files and directories
- [ ] `touch` - Create empty files or update timestamps
- [ ] `pwd` - Print working directory
- [ ] `cd` - Change directory (when running as a shell)
- [ ] `chmod` - Change file permissions
- [ ] `chown` - Change file ownership

### Text Processing
- [ ] `grep` - Search text patterns in files
- [ ] `sed` - Stream editor
- [ ] `awk` - Text processing
- [ ] `head` - Display first lines of files
- [ ] `tail` - Display last lines of files
- [ ] `sort` - Sort lines of text
- [ ] `uniq` - Remove duplicate lines
- [ ] `wc` - Word count

### System Information
- [ ] `ps` - List running processes
- [ ] `df` - Disk space usage
- [ ] `du` - Directory space usage
- [ ] `free` - Memory usage
- [ ] `uptime` - System uptime
- [ ] `uname` - System information
- [ ] `hostname` - Show/set hostname
- [ ] `id` - User and group IDs

### Networking
- [ ] `ping` - Network connectivity test
- [ ] `wget` / `curl` - Download files
- [ ] `nc` / `netcat` - Network utility
- [ ] `ifconfig` / `ip` - Network interface configuration

### Compression/Archives
- [ ] `tar` - Archive utility
- [ ] `gzip` / `gunzip` - Compression
- [ ] `zip` / `unzip` - ZIP archives

### Additional Features
- [ ] Command aliasing (e.g., `ll` → `ls -l`)
- [ ] Configuration file support
- [ ] Cross-platform compatibility improvements
- [ ] Performance optimizations
- [ ] Better error handling and exit codes
- [ ] Shell integration (symlink support)
- [ ] Unit tests for each command
- [ ] Integration tests
- [ ] CI/CD pipeline
- [ ] Release binaries for multiple platforms
- [ ] Documentation for each command

## License

Hosted under AGPLv3 (GNU Affero General Public License v3.0)

## Why Go?

- **Single Static Binary**: Easy to distribute, no dependencies
- **Fast Compilation**: Quick iteration and builds
- **Cross-Platform**: Build for Linux, macOS, Windows, etc.
- **Modern Tooling**: Great build system with build tags
- **No Forking Overhead**: In-process commands are faster than traditional busybox

## Acknowledgments

Inspired by BusyBox, which provides many Unix utilities in a single executable.

