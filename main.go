package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jcmuller/version_manager_shell_string/gochecker"
	"github.com/jcmuller/version_manager_shell_string/nvmchecker"
	"github.com/jcmuller/version_manager_shell_string/rubychecker"
	"github.com/jcmuller/version_manager_shell_string/versions"
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

	v := versions.New(path)

	rubychecker, err := rubychecker.New(path)
	if err == nil {
		v.AddChecker(rubychecker)
	}

	nvmchecker, err := nvmchecker.New(path)
	if err == nil {
		v.AddChecker(nvmchecker)
	}

	gochecker, err := gochecker.New(path)
	if err == nil {
		v.AddChecker(gochecker)
	}

	v.GetVersions()
	fmt.Println(v)
}
