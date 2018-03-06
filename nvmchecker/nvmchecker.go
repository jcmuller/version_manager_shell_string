package nvmchecker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// NvmChecker
type NvmChecker struct {
	basePath   string
	cmd        *exec.Cmd
	version    string
	defined    bool
	identifier string
}

var (
	file     = ".nvmrc"
	command  = "bash"
	argument = "/tmp/nvm_checker.sh"
)

// New instance
func New(path string) *NvmChecker {
	return &NvmChecker{
		basePath:   path,
		identifier: "N",
	}
}

// StartCheck
func (l *NvmChecker) StartCheck() {
	_, err := os.Stat(argument)

	if err != nil {
		file, err := os.Create(argument)
		handle(err)
		_, err = file.WriteString(script)
		handle(err)
		err = file.Sync()
		handle(err)
	}

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
func (l *NvmChecker) Wait() {
	err := l.cmd.Wait()
	handle(err)
}

func (l *NvmChecker) setDefined() {
	file = filepath.Join(l.basePath, file)
	_, err := os.Stat(file)
	l.defined = err == nil
}

// GetVersion does that
func (l *NvmChecker) GetVersion() {
	l.setDefined()
	l.Wait()
}

// Output string
func (l *NvmChecker) String() string {
	str := fmt.Sprintf("%s:%s", l.identifier, l.version)

	if l.defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return str
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (l *NvmChecker) IsDefined() bool {
	return l.defined
}
