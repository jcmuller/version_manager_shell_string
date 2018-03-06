package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jcmuller/utils/version_manager_shell_string/gochecker"
	"github.com/jcmuller/utils/version_manager_shell_string/nvmchecker"
	"github.com/jcmuller/utils/version_manager_shell_string/rubychecker"
	"github.com/jcmuller/utils/version_manager_shell_string/versions"
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

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	dir, err := os.Getwd()

	handle(err)

	path, err := findDirectory(dir)

	if err != nil {
		os.Exit(0)
	}

	v := versions.New()

	v.AddChecker(rubychecker.New(path))
	v.AddChecker(nvmchecker.New(path))
	v.AddChecker(gochecker.New(path))

	v.GetVersions()
	fmt.Println(v)
}
