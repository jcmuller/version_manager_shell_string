package langdef_test

import (
	"testing"

	"github.com/jcmuller/version_manager_shell_string/internal/langdef"
)

func TestStartCheck(t *testing.T) {
	l := &langdef.LangDef{
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
