package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
)

// main is the entry point of the Project Development Platform (PDP) application.
// It checks for a command provided by the user and executes the corresponding
// command plugin. The available commands include initializing a new project
// and displaying help information.
func main() {
	// Check if a command is provided
	if len(os.Args) < 2 {
		// Display an error message for missing command
		fmt.Println("Please provide a command (e.g., help or init).")
		return
	}

	command := os.Args[1] // Get the command from the command-line arguments

	// Get the path of the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Get the directory of the executable
	exeDir := filepath.Dir(exePath)

	// Construct the path to the cmd directory where command plugins are located
	cmdDir := filepath.Join(exeDir, "cmd")

	// Read all files in the cmd directory
	files, err := ioutil.ReadDir(cmdDir)
	if err != nil {
		fmt.Println("Error reading cmd directory:", err)
		return
	}

	// Load and execute the specified command plugin
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".so" {
			pluginPath := filepath.Join(cmdDir, file.Name())

			// Load the plugin
			p, err := plugin.Open(pluginPath)
			if err != nil {
				// Display an error message for loading plugin
				fmt.Printf("Error loading plugin %s: %v\n", file.Name(), err)
				continue
			}

			// Lookup the Execute function in the plugin
			execFunc, err := p.Lookup("Execute")
			if err != nil {
				// Display an error message for looking up Execute function
				fmt.Printf("Error looking up Execute function in %s: %v\n", file.Name(), err)
				continue
			}

			// Assert the function type (no return value)
			execute, ok := execFunc.(func())
			if !ok {
				// Display an error message for wrong function type
				fmt.Printf("Plugin function in %s has wrong type\n", file.Name())
				continue
			}

			// Check if the command matches the plugin name (without the .so extension)
			if file.Name() == command+".so" {
				// Call the function from the plugin
				execute() // No need to capture a return value
				return // Exit after executing the command
			}
		}
	}

	// Display an error message for command not found
	fmt.Printf("Command '%s' not found.\n", command)
}