// util is a collection of utilities
package util

import (
	"path/filepath"
	"os"
	"path"
	"strings"
	"errors"
)

// FindDirectory finds a directory that contains the given file
func FindDirectory(input string) (string, error) {
	candidate := filepath.Join(input, ".git")
	info, err := os.Stat(candidate)

	if err == nil && info.IsDir() {
		return input, nil
	}

	// Go up one level
	dir, _ := path.Split(input)
	dir = strings.TrimRight(dir, "/")

	if dir == "" {
		err := errors.New("No .git directory found")
		return "", err
	}

	return FindDirectory(dir)
}

