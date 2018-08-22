package checker_test

import (
	"testing"

	"github.com/jcmuller/version_manager_shell_string/internal/checker"
)

func TestStartCheck(t *testing.T) {
	l := &checker.Checker{
		CommandName: "echo",
		Args:        []string{"1.23.4"},
		Identifier:  "A",
	}

	l.Prepare("some path")
	l.GetVersion()

	actual := l.String()
	expected := "A:1.23.4"

	if actual != expected {
		t.Errorf("Incorrect version string.\n  Expected: %v\n    Actual: %s", expected, actual)
	}
}
