package gochecker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GoChecker is a thing
type GoChecker struct {
	basePath   string
	cmd        *exec.Cmd
	version    string
	defined    bool
	identifier string
}

var (
	file     = ".go-version"
	command  = "/home/jcmuller/.goenv/bin/goenv"
	argument = "version-name"
)

// New go checker
func New(path string) *GoChecker {
	return &GoChecker{
		basePath:   path,
		identifier: "G",
	}
}

func (l *GoChecker) setDefined() {
	file = filepath.Join(l.basePath, file)
	_, err := os.Stat(file)
	l.defined = err == nil
}

// StartCheck
func (l *GoChecker) StartCheck() {
	l.cmd = exec.Command(command, argument)

	reader, err := l.cmd.StdoutPipe()
	handle(err)

	scanner := bufio.NewScanner(reader)

	go func() {
		for scanner.Scan() {
			l.version = scanner.Text()
		}
	}()

	err = l.cmd.Start()
	handle(err)
}

// Wait
func (l *GoChecker) Wait() {
	err := l.cmd.Wait()
	handle(err)
}

// GetVersion does that
func (l *GoChecker) GetVersion() {
	l.setDefined()
	l.Wait()
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Output string
func (l *GoChecker) String() string {
	str := fmt.Sprintf("%s:%s", l.identifier, l.version)

	if l.defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return str
}

func (l *GoChecker) IsDefined() bool {
	return l.defined
}
