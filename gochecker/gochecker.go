package gochecker

import (
	"os/exec"

	"github.com/jcmuller/version_manager_shell_string/langdef"
)

type GoChecker struct {
	*langdef.LangDef
}

func New(path string) *GoChecker {
	cmdPath, err := exec.LookPath("goenv")
	if err != nil {
		panic("Unable to find 'goenv' executable")
	}

	return &GoChecker{
		&langdef.LangDef{
			BasePath:   path,
			Identifier: "G",
			File:       ".go-version",
			Command:    exec.Command(cmdPath, "version-name"),
		},
	}
}
