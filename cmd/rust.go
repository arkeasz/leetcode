package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func RunRust(file string) {
	cmd := exec.Command("rustc", file)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error compiling Rust file: %v\n", err)
		return
	}
	defer os.Remove("main")

	cmd = exec.Command("./main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running Rust binary: %v\n", err)
	}
}
