package main

import (
	"fmt"
	"os"
	"github.com/jcmuller/version_manager_shell_string/internal/config"
	"github.com/jcmuller/version_manager_shell_string/internal/versions"
	"github.com/jcmuller/version_manager_shell_string/internal/util"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting CWD: %+v\n", err)
		os.Exit(1)
	}

	path, _ := util.FindDirectory(dir)

	c := config.New()
	v := versions.New(c, path)
	v.GetVersions()
	fmt.Println(v)
}
