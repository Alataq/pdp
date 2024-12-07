package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
)

func Execute() {
	var appName, appDirName, appDir, appType, appTemplate, appTemplateDir string
	// Get the path of the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Get the directory of the executable
	exeDir := filepath.Dir(exePath)


}