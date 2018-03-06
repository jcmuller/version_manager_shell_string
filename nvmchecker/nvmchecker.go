package nvmchecker

import (
	"log"
	"os"
	"os/exec"

	"github.com/jcmuller/version_manager_shell_string/langdef"
)

// NvmChecker
type NvmChecker struct {
	*langdef.LangDef
}

var (
	argument = "/tmp/nvm_checker.sh"
)

// New instance
func New(path string) *NvmChecker {
	return &NvmChecker{
		&langdef.LangDef{
			BasePath:   path,
			Identifier: "N",
			File:       ".nvmrc",
			Command:    exec.Command("bash", argument),
		},
	}
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
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

	l.LangDef.StartCheck()
}
