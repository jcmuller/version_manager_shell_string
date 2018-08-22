package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/jcmuller/version_manager_shell_string/internal/checker"
)

var (
	configFile = filepath.Join(
		os.Getenv("HOME"),
		".config",
		"version_manager_shell_string",
		"config.yml",
	)
)

func (c *Config) Checkers() (checkers []*checker.Checker) {
	checkers = c.checkers

	return
}

type Config struct {
	checkers []*checker.Checker
}

func New() (c *Config) {
	c = &Config{}

	configFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read configuration file (%s): %+v\n", configFile, err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(configFile), &c.checkers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing YAML: %+v\n", err)
		os.Exit(1)
	}

	return
}
