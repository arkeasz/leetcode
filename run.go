package main

import (
	"fmt"
	"os"
	"strings"
	"run/cmd"
	"run/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: run.go <folder_name>")
		return
	}

	folder := "solutions/"
	folder += os.Args[1]

	files, err := utils.ReadDir(folder)
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
		cmd.RunJavaScript(mainFile)
	} else if strings.HasSuffix(mainFile, ".rs") {
		cmd.RunRust(mainFile)
	} else if strings.HasSuffix(mainFile, ".py") {
		cmd.RunPython(mainFile)
	} else {
		fmt.Println("Error: Unsupported file type.")
	}
}
