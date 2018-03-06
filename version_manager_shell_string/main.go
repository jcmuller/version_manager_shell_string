package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jcmuller/utils/version_manager_shell_string/langdef"
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

func output(checkers []*langdef.LangDef) {
	for _, element := range checkers {
		fmt.Printf(element.String())
	}
}

func main() {
	dir, err := os.Getwd()

	handle(err)

	path, err := findDirectory(dir)

	if err != nil {
		os.Exit(0)
	}

	v := versions.New(path)

	v.AddChecker(langdef.New(".ruby-version", "R"))
	v.AddChecker(langdef.New(".go-version", "G"))
	v.AddChecker(langdef.New(".nvmrc", "N"))

	v.GetVersions()
	//getVersions(path, checkers)
	//output(checkers)
	fmt.Println(v)
}
