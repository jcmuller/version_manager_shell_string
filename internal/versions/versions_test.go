package versions_test

import (
	"testing"

	"github.com/jcmuller/version_manager_shell_string/internal/checker"
	"github.com/jcmuller/version_manager_shell_string/internal/versions"
)

type mockConfig struct{}

var (
	checker_a = &checker.Checker{
		CommandName: "echo",
		Args:        []string{"1.23.4"},
		Identifier:  "A",
	}
	checker_b = &checker.Checker{
		CommandName: "echo",
		Args:        []string{"99Foo"},
		Identifier:  "Z",
	}
	checker_c = &checker.Checker{
		CommandName: "echo",
		Args:        []string{"3.1"},
		Identifier:  "R",
	}
)

func (c *mockConfig) Checkers() []*checker.Checker {
	return []*checker.Checker{checker_a, checker_b, checker_c}
}

var (
	c = &mockConfig{}
)

func TestString(t *testing.T) {
	v := versions.New(c, "")
	v.GetVersions()
	expected := "A:1.23.4|Z:99Foo|R:3.1"
	actual := v.String()

	if expected != actual {
		t.Errorf("Incorrect version string.\n  Expected: %v\n    Actual: %s", expected, actual)
	}
}
