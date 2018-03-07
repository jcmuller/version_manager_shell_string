package rubychecker

import (
	"errors"
	"os/exec"

	"github.com/jcmuller/version_manager_shell_string/langdef"
)

// RubyChecker is a thing
type RubyChecker struct {
	*langdef.LangDef
}

// New ruby checker
func New(path string) (*RubyChecker, error) {
	cmdPath, err := exec.LookPath("rbenv")
	if err != nil {
		return nil, errors.New("Couldn't find rbenv")
	}

	return &RubyChecker{
		&langdef.LangDef{
			BasePath:   path,
			Identifier: "R",
			File:       ".ruby-version",
			Command:    exec.Command(cmdPath, "version-name"),
		},
	}, nil
}
