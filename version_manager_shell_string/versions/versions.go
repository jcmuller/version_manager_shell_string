package versions

import (
	"strings"

	"github.com/jcmuller/utils/version_manager_shell_string/langdef"
)

// Versions Hold these guys
type Versions struct {
	path     string
	checkers []*langdef.LangDef
}

// New new version
func New(path string) *Versions {
	return &Versions{path: path}
}

// Add a checker
func (v *Versions) AddChecker(checker *langdef.LangDef) {
	v.checkers = append(v.checkers, checker)
}

// GetVersions does that
func (v *Versions) GetVersions() {
	for _, element := range v.checkers {
		element.GetVersion(v.path)
	}
}

func (v *Versions) presentVersions() []string {
	o := make([]string, 0)

	for _, element := range v.checkers {
		if element.Present() {
			o = append(o, element.String())
		}
	}

	return o
}

func (v *Versions) String() string {
	return strings.Join(v.presentVersions(), "|")
}
