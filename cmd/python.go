package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func RunPython(file string) {
	cmd := exec.Command("python3", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing Python file: %v\n", err)
	}
}
