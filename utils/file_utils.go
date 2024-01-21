package utils

import (
	"fmt"
	"os"

	"github.com/MhamzaAhmad/commence/constants"
)


func DeleteIfExists(filename string) {
	dir := constants.STORAGE_DIR

	CheckForDirectory(dir)

	if _, err := os.Stat(fmt.Sprintf("%s/%s", dir, filename)); err == nil {
		os.Remove(fmt.Sprintf("%s/%s", dir, filename))
	}
}

func CheckForDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

func CheckForFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CheckIfFileHasContent(path string) bool {
	if file, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		if file.Size() == 0 {
			return false
		}
	}

	return true
}
