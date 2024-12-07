package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Execute() {
	var appName string

	// Get the current working directory (user's directory)
	userDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Check if project.json exists
	projectFilePath := filepath.Join(userDir, "project.json")
	if _, err := os.Stat(projectFilePath); os.IsNotExist(err) {
		fmt.Println("Error: No project initialized. Please run the 'init' command first.")
		return
	}

	reader := bufio.NewReader(os.Stdin) // Create a new reader

	for {
		fmt.Print("Enter application name: ")
		appName, err = reader.ReadString('\n') // Read the entire line
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		appName = strings.TrimSpace(appName) // Trim any extra whitespace
		if appName != "" {
			break
		}
		fmt.Println("Application name is required. Please enter a valid name.")
	}

	// Convert appName to lowercase and replace spaces with underscores
	appName = strings.ToLower(strings.ReplaceAll(appName, " ", "_"))

	// Create a directory with the modified application name in the user's directory
	appDir := filepath.Join(userDir, appName) // Set the directory path
	err = os.Mkdir(appDir, os.ModePerm)        // Create the directory
	if err != nil {
		fmt.Printf("Error creating directory '%s': %v\n", appDir, err)
		return
	}

	fmt.Printf("Directory '%s' created successfully.\n", appDir)
	// Continue with the rest of the create command logic...
}