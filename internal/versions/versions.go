package versions

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jcmuller/version_manager_shell_string/internal/config"
)

type checker interface {
	GetVersion()
	String() string
	StartCheck()
	IsDefined() bool
	Prepare(string)
}

// Versions Hold these guys
type Versions struct {
	path   string
	config *config.Config
	//	checkers []langdef.LangDef
	checkers    []checker
	onlyDefined bool
}

// New new version
func New(config *config.Config, path string) *Versions {
	checkers := make([]checker, len(config.Checkers()))
	for i, v := range config.Checkers() {
		checkers[i] = checker(v)
	}

	return &Versions{
		path:        path,
		onlyDefined: false,
		checkers:    checkers,
	}
}

// GetVersions does that
func (v *Versions) GetVersions() {
	for _, element := range v.checkers {
		element.Prepare(v.path)
		element.GetVersion()
	}
}

func (v *Versions) presentVersions() []string {
	o := make([]string, 0)

	v.setOnlyDefined()

	for _, element := range v.checkers {
		if v.onlyDefined {
			if element.IsDefined() {
				o = append(o, element.String())
			}
		} else {

			o = append(o, element.String())
		}
	}

	return o
}

func (v *Versions) String() string {
	return strings.Join(v.presentVersions(), "|")
}

func (v *Versions) setOnlyDefined() {
	file := filepath.Join(v.path, ".only_defined")
	_, err := os.Stat(file)

	v.onlyDefined = err == nil
}
