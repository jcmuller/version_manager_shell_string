package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jcmuller/version_manager_shell_string/internal/config"
	"github.com/jcmuller/version_manager_shell_string/internal/versions"
)

func findDirectory(input string) (string, error) {
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

	return findDirectory(dir)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting CWD: %+v\n", err)
		os.Exit(1)
	}

	path, err := findDirectory(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No directory found: %+v\n", err)
		os.Exit(1)
	}

	c := config.New()
	v := versions.New(c, path)
	v.GetVersions()
	fmt.Println(v)
}
