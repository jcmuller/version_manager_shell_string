package versions

import (
	"strings"
)

type checker interface {
	GetVersion()
	String() string
	StartCheck()
}

// Versions Hold these guys
type Versions struct {
	checkers []checker
}

// New new version
func New() *Versions {
	return &Versions{}
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

	for _, element := range v.checkers {
		o = append(o, element.String())
	}

	return o
}

func (v *Versions) String() string {
	return strings.Join(v.presentVersions(), "|")
}
