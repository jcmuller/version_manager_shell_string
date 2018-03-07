package gochecker

import (
	"errors"
	"os/exec"

	"github.com/jcmuller/version_manager_shell_string/langdef"
)

type GoChecker struct {
	*langdef.LangDef
}

func New(path string) (*GoChecker, error) {
	cmdPath, err := exec.LookPath("goenv")
	if err != nil {
		return nil, errors.New("Couldn't find goenv")
	}

	return &GoChecker{
		&langdef.LangDef{
			BasePath:   path,
			Identifier: "G",
			File:       ".go-version",
			Command:    exec.Command(cmdPath, "version-name"),
		},
	}, nil
}
