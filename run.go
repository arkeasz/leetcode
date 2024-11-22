package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: run_problem.go <folder_name>")
		return
	}

	folder := os.Args[1]

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var mainFile string

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "main.") {
			mainFile = fmt.Sprintf("%s/%s", folder, file.Name())
			break
		}
	}

	if mainFile == "" {
		fmt.Println("Error: No main file found.")
		return
	}

	if strings.HasSuffix(mainFile, ".js") {
		cmd := exec.Command("node", mainFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else if strings.HasSuffix(mainFile, ".rs") {
		cmd := exec.Command("rustc", mainFile, "-o", "main_bin")
		cmd.Run()
		cmd = exec.Command("./main_bin")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		fmt.Println("Error: Unsupported file type.")
	}
}
