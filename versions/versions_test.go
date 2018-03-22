package versions_test

import (
	"testing"

	"github.com/jcmuller/version_manager_shell_string/versions"
)

type checkerMock struct {
	version string
}

func (c *checkerMock) GetVersion() {}

func (c *checkerMock) String() string {
	if c.version != "" {
		return c.version
	}

	return "foo"
}

func (c *checkerMock) StartCheck() {}

func (c *checkerMock) IsDefined() bool { return true }

func TestAddChecker(t *testing.T) {
	v := versions.New("foo")
	c := &checkerMock{}

	v.AddChecker(c)

	if len(v.Checkers()) != 1 {
		t.Error("Incorrect number of checkers")
	}
}

func TestGetVersions(t *testing.T) {
	v := versions.New("foo")
	v.GetVersions()
}

func TestString(t *testing.T) {
	v := versions.New("foo")

	v.AddChecker(&checkerMock{"V:1"})
	v.AddChecker(&checkerMock{"G:10"})
	v.AddChecker(&checkerMock{"H:21"})

	expected := "V:1|G:10|H:21"
	actual := v.String()

	if expected != actual {
		t.Errorf("Incorrect version string.\n  Expected: %v\n    Actual: %s", expected, actual)
	}
}
