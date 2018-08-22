package checker

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Checker struct {
	BasePath    string
	Command     *exec.Cmd
	Version     string
	Defined     bool
	CommandName string   `json:"command"`
	Args        []string `json:"args"`
	Identifier  string   `json:"identifier"`
	File        string   `json:"file"`
}

// StartCheck
func (l *Checker) StartCheck() {
	reader, err := l.Command.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting stdout pipe: %+v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(reader)

	go func() {
		for scanner.Scan() {
			l.Version = scanner.Text()
		}
	}()

	err = l.Command.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting command: %+v\n", err)
		os.Exit(1)
	}
}

// Wait
func (l *Checker) Wait() {
	err := l.Command.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for command: %+v\n", err)
		os.Exit(1)
	}
}

func (l *Checker) Prepare(path string) {
	l.Command = exec.Command(l.CommandName, l.Args...)
	l.BasePath = path
	l.StartCheck()
}

func (l *Checker) setDefined() {
	file := filepath.Join(l.BasePath, l.File)
	_, err := os.Stat(file)
	l.Defined = err == nil
}

// GetVersion does that
func (l *Checker) GetVersion() {
	l.setDefined()
	l.Wait()
}

// Output string
func (l *Checker) String() (str string) {
	str = fmt.Sprintf("%s:%s", l.Identifier, l.Version)

	if l.Defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return
}

func (l *Checker) IsDefined() bool {
	return l.Defined
}
