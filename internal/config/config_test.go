package config_test

import (
	"fmt"
	"testing"

	"github.com/jcmuller/version_manager_shell_string/internal/config"
)

func TestStartCheck(t *testing.T) {
	c := config.New()
	fmt.Printf("%+v\n", c)
}
