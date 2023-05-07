package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Get modified time of a file
func getFileModTime(filePath string) (time.Time, error) {
	stat, err := os.Stat(filePath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return time.Time{}, err
	}
	if err != nil {
		log.Println(err)
		log.Println("Cannot get file stats")
		return time.Time{}, err
	}

	return stat.ModTime(), nil
}

// Open a file using an external text editor. Return a bool indicates if file has been changed
func OpenExternalEditor(filePath string) bool {
	const DefaultEditor = "nano"
	
	mtBefore, err := getFileModTime(filePath)
	if err != nil && !errors.Is(err, os.ErrNotExist){
		return false
	}
	editorExecutable := os.Getenv("EDITOR") // Get default text editor
	if editorExecutable == "" {
		editorExecutable = DefaultEditor
	}

	editor := exec.Command(editorExecutable, filePath)
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr

	err = editor.Run()
	if err != nil {
		return false
	}
	mtAfter, err := getFileModTime(filePath)
	if err != nil {
		return false
	}

	if mtBefore.Before(mtAfter) {
		return true
	}
	return false
}

// Create a Y/n cli prompt
func Prompt(message string, loop bool) bool {
	for ok := true; ok; ok = loop {
		var choice string
		fmt.Printf("%s [Y/n]: ", message)
		fmt.Scanf("%s", &choice)

		if strings.ToLower(choice) == "y"  {
			return true
		}
		if strings.ToLower(choice) == "n" {
			return false
		}
	}
	
	return false
}
