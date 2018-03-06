package versions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type checker interface {
	GetVersion()
	String() string
	StartCheck()
	IsDefined() bool
}

// Versions Hold these guys
type Versions struct {
	path        string
	checkers    []checker
	onlyDefined bool
}

// New new version
func New(path string) *Versions {
	return &Versions{
		path:        path,
		onlyDefined: false,
	}
}

// Add a checker
func (v *Versions) AddChecker(checker checker) {
	v.checkers = append(v.checkers, checker)
	checker.StartCheck()
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
	stats, err := os.Stat(file)

	fmt.Println("Only defined")
	fmt.Println(spew.Sdump(stats))
	fmt.Println(spew.Sdump(err))
	v.onlyDefined = err == nil
}
