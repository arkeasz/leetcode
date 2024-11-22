package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func RunJavaScript(file string) {
	cmd := exec.Command("node", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing JavaScript file: %v\n", err)
	}
}
