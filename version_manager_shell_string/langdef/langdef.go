package langdef

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// LangDef Language definition
type LangDef struct {
	identifier string
	filename   string
	path       string
	version    string
}

// Present Does the language version exist?
func (l *LangDef) Present() bool {
	return l.version != ""
}

// Output string
func (l *LangDef) String() string {
	return fmt.Sprintf("%s:%s", l.identifier, l.version)
}

// GetVersion gets the Version
func (l *LangDef) GetVersion(path string) {
	candidate := filepath.Join(path, l.filename)
	_, err := os.Stat(candidate)

	if err != nil {
		return
	}

	l.path = candidate

	file, err := os.Open(candidate)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l.version = scanner.Text()
	}

	err = scanner.Err()
	handle(err)
}

// New Instance
func New(filename string, identifier string) *LangDef {
	return &LangDef{
		filename:   filename,
		identifier: identifier,
	}
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
