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
	return &RubyChecker{
		&langdef.LangDef{
			BasePath:   path,
			Identifier: "R",
			File:       ".ruby-version",
			Command:    exec.Command("/home/jcmuller/.rbenv/bin/rbenv", "version-name"),
		},
	}
}
