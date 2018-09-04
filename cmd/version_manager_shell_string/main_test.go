package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFindDirectoryNoDirectory(t *testing.T) {
	_, err := findDirectory("foo")

	if err == nil {
		t.Error("findDirectory(foo) is supposed to error out!")
	}
}

func setupProject() (base, dir string) {
	tmp := os.TempDir()
	base = strings.Join([]string{tmp, "go", "vmss", "test"}, "/")
	dir = fmt.Sprintf("%s/foo/bar", base)
	gitDir := fmt.Sprintf("%s/.git", base)

	os.MkdirAll(gitDir, 0755)
	os.MkdirAll(dir, 0755)

	return
}

func cleanUpProject(base string) {
	os.RemoveAll(base)
}

func TestFindDirectory(t *testing.T) {
	base, dir := setupProject()

	actual, err := findDirectory(dir)

	if err != nil {
		t.Errorf("findDirectory() is not supposed to err: %v", err)
	}

	expected := base

	if actual != base {
		t.Errorf("findDirectory() error: expected: %v, actual: %v", expected, actual)
	}

	cleanUpProject(base)
}
