package utils

import (
	"os"
)

func ReadDir(folder string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	return files, nil
}
