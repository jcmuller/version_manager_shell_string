package rubychecker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RubyChecker is a thing
type RubyChecker struct {
	basePath   string
	cmd        *exec.Cmd
	version    string
	defined    bool
	identifier string
}

var (
	file     = ".ruby-version"
	command  = "/home/jcmuller/.rbenv/bin/rbenv"
	argument = "version-name"
)

// New ruby checker
func New(path string) *RubyChecker {
	return &RubyChecker{
		basePath:   path,
		identifier: "R",
	}
}

func (l *RubyChecker) setDefined() {
	file = filepath.Join(l.basePath, file)
	_, err := os.Stat(file)
	l.defined = err == nil
}

// GetVersion does that
func (l *RubyChecker) GetVersion() {
	l.setDefined()
	l.Wait()
}

// StartCheck
func (l *RubyChecker) StartCheck() {
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
func (l *RubyChecker) Wait() {
	err := l.cmd.Wait()
	handle(err)
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Output string
func (l *RubyChecker) String() string {
	str := fmt.Sprintf("%s:%s", l.identifier, l.version)

	if l.defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return str
}

func (l *RubyChecker) IsDefined() bool {
	return l.defined
}
