package langdef

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type LangDef struct {
	BasePath   string
	Identifier string
	File       string
	Command    *exec.Cmd
	Version    string
	Defined    bool
}

// StartCheck
func (l *LangDef) StartCheck() {
	reader, err := l.Command.StdoutPipe()
	handle(err)

	scanner := bufio.NewScanner(reader)

	go func() {
		for scanner.Scan() {
			l.Version = scanner.Text()
		}
	}()

	err = l.Command.Start()
	handle(err)
}

// Wait
func (l *LangDef) Wait() {
	err := l.Command.Wait()
	handle(err)
}

func (l *LangDef) setDefined() {
	file := filepath.Join(l.BasePath, l.File)
	_, err := os.Stat(file)
	l.Defined = err == nil
}

// GetVersion does that
func (l *LangDef) GetVersion() {
	l.setDefined()
	l.Wait()
}

// Output string
func (l *LangDef) String() string {
	str := fmt.Sprintf("%s:%s", l.Identifier, l.Version)

	if l.Defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return str
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (l *LangDef) IsDefined() bool {
	return l.Defined
}
