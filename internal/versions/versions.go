package versions

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jcmuller/version_manager_shell_string/internal/checker"
)

type chkr interface {
	GetVersion()
	IsDefined() bool
	Prepare(string)
	StartCheck()
	String() string
}

type cfg interface {
	Checkers() []*checker.Checker
}

// Versions Hold these guys
type Versions struct {
	checkers    []chkr
	onlyDefined bool
	path        string
}

// New new version
func New(c cfg, path string) *Versions {
	checkers := make([]chkr, len(c.Checkers()))
	for i, v := range c.Checkers() {
		checkers[i] = chkr(v)
		v.Prepare(path)
	}

	return &Versions{
		checkers:    checkers,
		onlyDefined: false,
		path:        path,
	}
}

// GetVersions does that
func (v *Versions) GetVersions() {
	for _, element := range v.checkers {
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
