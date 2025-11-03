package main

// Command represents a single command handler
type Command struct {
	Name        string
	Description string
	Handler     func(args []string)
}

var commands = make(map[string]*Command)

// registerCommand registers a command in the global registry
func registerCommand(cmd *Command) {
	commands[cmd.Name] = cmd
}

// getCommand retrieves a command by name
func getCommand(name string) *Command {
	return commands[name]
}

// listCommands returns all registered commands
func listCommands() []*Command {
	var result []*Command
	for _, cmd := range commands {
		result = append(result, cmd)
	}
	return result
}
