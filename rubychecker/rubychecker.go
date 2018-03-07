package rubychecker

import (
	"os/exec"

	"github.com/jcmuller/version_manager_shell_string/langdef"
)

// RubyChecker is a thing
type RubyChecker struct {
	*langdef.LangDef
}

// New ruby checker
func New(path string) *RubyChecker {
	cmdPath, err := exec.LookPath("rbenv")
	if err != nil {
		panic("Unable to find 'rbenv' executable")
	}

	return &RubyChecker{
		&langdef.LangDef{
			BasePath:   path,
			Identifier: "R",
			File:       ".ruby-version",
			Command:    exec.Command(cmdPath, "version-name"),
		},
	}
}
