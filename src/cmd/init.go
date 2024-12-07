package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
)

func Execute() {
	var appName, author, gitRepo, description, license string

	// Prompt for project name (required)
	for {
		fmt.Print("Enter project name: ")
		fmt.Scanln(&appName)
		if appName != "" {
			break
		}
		fmt.Println("Project name is required. Please enter a valid name.")
	}

	// Prompt for author name (required)
	for {
		fmt.Print("Enter author name: ")
		fmt.Scanln(&author)
		if author != "" {
			break
		}
		fmt.Println("Author name is required. Please enter a valid name.")
	}

	// Set Git repository URL to none by default
	gitRepo = ""

	// Set default values
	defaultDescription := "A project that encompasses various applications and services."
	defaultLicense := "MIT"

	// Prompt for description with default value
	fmt.Printf("Enter description (%s): ", defaultDescription)
	var descriptionInput string
	_, err := fmt.Scanln(&descriptionInput)
	if err == nil && descriptionInput != "" {
		description = descriptionInput
	} else {
		description = defaultDescription
	}

	// Prompt for license with default value
	fmt.Printf("Enter license (%s): ", defaultLicense)
	var licenseInput string
	_, err = fmt.Scanln(&licenseInput)
	if err == nil && licenseInput != "" {
		license = licenseInput
	} else {
		license = defaultLicense
	}

	// Get the path of the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Get the directory of the executable (output directory)
	exeDir := filepath.Dir(exePath)

	// Construct the path to the init directory
	initDir := filepath.Join(exeDir, "init")

	// Check if the init directory exists
	if _, err := os.Stat(initDir); os.IsNotExist(err) {
		fmt.Printf("Error: The init directory does not exist at path: %s\n", initDir)
		return
	}

	// Copy files and directories from the init directory to the current working directory
	err = copyDir(initDir, ".")
	if err != nil {
		fmt.Printf("Error copying files: %v\n", err)
		return
	}

	// Create a full project.json file path in the current directory
	fullProjectJsonFilePath := filepath.Join(".", "project.json")
	projectContent := map[string]interface{}{
		"projectName":  appName,
		"author":       author,
		"gitRepo":      gitRepo,
		"description":  description,
		"license":      license,
	}

	// Marshal the project content to JSON
	jsonData, err := json.MarshalIndent(projectContent, "", "  ")
	if err != nil {
		fmt.Printf("Error creating project.json content: %v\n", err)
		return
	}

	// Write the JSON data to the project.json file
	err = os.WriteFile(fullProjectJsonFilePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error creating project.json file: %v\n", err)
		return
	}

	fmt.Println("Project initialized successfully. You can now manage applications and services within this project.")
}

// copyDir copies the contents of the source directory to the destination directory.
func copyDir(src string, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Create the destination directory
			err := os.MkdirAll(dstPath, os.ModePerm)
			if err != nil {
				return err
			}
			// Recursively copy the directory
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			// Copy the file
			input, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			err = os.WriteFile(dstPath, input, entry.Mode())
			if err != nil {
				return err
			}
		}
	}
	return nil // Return nil to indicate no error occurred
}
