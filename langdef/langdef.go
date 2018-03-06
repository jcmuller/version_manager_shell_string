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
	basePath   string
	version    string
}

// Output string
func (l *LangDef) String() string {
	return fmt.Sprintf("%s:%s", l.identifier, l.version)
}

// GetVersion gets the Version
func (l *LangDef) GetVersion() {
	candidate := filepath.Join(l.basePath, l.filename)
	_, err := os.Stat(candidate)

	if err != nil {
		return
	}

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
func New(path string, filename string, identifier string) *LangDef {
	return &LangDef{
		basePath:   path,
		filename:   filename,
		identifier: identifier,
	}
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
