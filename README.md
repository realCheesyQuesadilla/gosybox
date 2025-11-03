<img width="250" height="300" align="center" alt="image" src="https://github.com/user-attachments/assets/7688dc9a-f46e-41a1-803a-681b37113c64" />  

# gosybox

A lightweight, modular BusyBox replacement written in Go. gosybox provides essential Unix utilities in a single binary, with the ability to customize which commands are included at compile time using Go build tags.

## Features

- **Single Binary**: All commands are compiled into one executable
- **Interactive Shell Mode**: Run as a long-running process and execute commands interactively
- **Command Mode**: Traditional command-line usage (execute and exit)
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

gosybox supports two modes of operation:

#### Command Mode (Execute and Exit)

Run a single command:
```bash
./gosybox <command> [args...]
```

Examples:
```bash
./gosybox ls
./gosybox ls /home
./gosybox lt
./gosybox help
```

#### Interactive Shell Mode

Enter an interactive shell where you can run multiple commands:
```bash
./gosybox
# or explicitly:
./gosybox -i
```

Once in interactive mode, you'll see a prompt:
```
gosybox interactive mode
Type 'help' for available commands, 'exit' or 'quit' to exit
gosybox> 
```

Example interactive session:
```bash
$ ./gosybox
gosybox interactive mode
Type 'help' for available commands, 'exit' or 'quit' to exit
gosybox> ls
file1.txt
dir1/
gosybox> lt
Name  Size  ModTime  Perms  Owner  Group
----------------------------------------
dir1/  4096  2024-01-15 10:30:00  drwxr-xr-x  1000  1000
file1.txt  1024  2024-01-15 10:35:00  -rw-r--r--  1000  1000
gosybox> help
I am gosybox, a replacement for busybox written in Go.
Available commands:
  ls       List directory contents
  lt       List directory contents and order by modification time
  help     Show this help message
  exit     Exit interactive mode (alias: quit)
gosybox> exit
Goodbye!
```

Exit interactive mode with:
- `exit` or `quit` command
- `Ctrl+D` (EOF)

## Customizing Builds

### Excluding Commands

Use Go build tags to exclude specific commands:

```bash
# Exclude ls command
go build -tags no_ls -o gosybox

# Exclude lt command
go build -tags no_lt -o gosybox

# Exclude help command
go build -tags no_help -o gosybox

# Exclude exit command
go build -tags no_exit -o gosybox

# Exclude multiple commands
go build -tags "no_ls no_lt no_help no_exit" -o gosybox
```

This allows you to create minimal builds with only the commands you need, reducing binary size.

### How Build Tags Work

Each command is in its own file with a build tag:
- `cmd_ls.go` - includes `ls` unless `no_ls` tag is set
- `cmd_lt.go` - includes `lt` unless `no_lt` tag is set
- `cmd_help.go` - includes `help` unless `no_help` tag is set
- `cmd_exit.go` - includes `exit` unless `no_exit` tag is set

Commands register themselves in `init()` functions, which run automatically when the package loads.

## Available Commands

- `ls` - List directory contents (marks directories with `/`)
- `lt` - List directory contents sorted by modification time with detailed information (size, permissions, owner, group)
- `help` - Show available commands and help information
- `exit` - Exit interactive mode (alias: `quit`)

### Command Examples

**List current directory:**
```bash
./gosybox ls
./gosybox ls /home/user
```

**List with details sorted by modification time:**
```bash
./gosybox lt
./gosybox lt /var/log
```

The `lt` command displays:
- File/directory name (directories marked with `/`)
- File size in bytes
- Modification time (YYYY-MM-DD HH:MM:SS format)
- File permissions (e.g., `-rw-r--r--`, `drwxr-xr-x`)
- Owner UID
- Group GID

Files are sorted by modification time (oldest first).

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
├── gosybox.go      # Main entry point, command dispatcher, and interactive shell
├── commands.go     # Command registry system
├── cmd_ls.go      # ls command implementation
├── cmd_lt.go      # lt command implementation (sorted by modification time)
├── cmd_help.go    # help command implementation
├── cmd_exit.go    # exit command implementation
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
- [x] Interactive shell mode
- [ ] Command aliasing (e.g., `ll` → `ls -l`)
- [ ] Configuration file support
- [ ] Command history in interactive mode
- [ ] Tab completion in interactive mode
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
