package checker

import (
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
	// Identifier is the string shown in the output
	Identifier string `json:"identifier"`
	// File is used to show whether the version is defined in the current directory
	File string `json:"file"`
}

func (c *Checker) Run(path string, out chan string) {
	c.Command = exec.Command(c.CommandName, c.Args...)
	c.BasePath = path
	c.setDefined()

	output, err := c.Command.Output()

	if err != nil {
		c.Version = "ERR"
	} else {
		c.Version = strings.Trim(string(output), "\n")
	}

	out <- c.String()
}

func (c *Checker) setDefined() {
	file := filepath.Join(c.BasePath, c.File)
	_, err := os.Stat(file)
	c.Defined = err == nil
}

// Output string
func (c *Checker) String() (str string) {
	str = fmt.Sprintf("%s:%s", c.Identifier, c.Version)

	if c.Defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return
}

func (c *Checker) IsDefined() bool {
	return c.Defined
}
