package versions

import (
	"strings"

	"github.com/jcmuller/version_manager_shell_string/internal/checker"
	"sort"
)

type cfg interface {
	Checkers() []*checker.Checker
}

// Versions Hold these guys
type Versions struct {
	config  cfg
	path    string
	results []string
}

// New new version
func New(config cfg, path string) *Versions {
	checkerCount := len(config.Checkers())
	results := make([]string, checkerCount, checkerCount)
	return &Versions{config, path, results}
}

func (v *Versions) checkers() (checkers []*checker.Checker) {
	return v.config.Checkers()
}

func (v *Versions) Check() {
	versionChannel := make(chan string)

	for _, c := range v.checkers() {
		go c.Run(v.path, versionChannel)
	}

	for i := 0; i < len(v.checkers()); i++ {
		msg := <-versionChannel
		v.results[i] = msg
	}

	sort.Strings(v.results)
}

func (v *Versions) String() string {
	return strings.Join(v.results, "|")
}
