package langdef_test

import (
	"os/exec"
	"testing"

	"github.com/jcmuller/version_manager_shell_string/internal/langdef"
)

func TestStartCheck(t *testing.T) {
	l := &langdef.LangDef{
		Identifier: "T",
		Command:    exec.Command("echo", "foobar"),
	}

	l.StartCheck()
	l.GetVersion()
	l.IsDefined()

	actual := l.String()
	expected := "T:foobar"

	if actual != expected {
		t.Errorf("Incorrect version string.\n  Expected: %v\n    Actual: %s", expected, actual)
	}
}
