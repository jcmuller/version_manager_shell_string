package main

import (
	"testing"
	"github.com/jcmuller/version_manager_shell_string/internal/versions"
	"github.com/jcmuller/version_manager_shell_string/internal/config"

	versionsOld "github.com/jcmuller/version_manager_shell_string_old/versions"
	"github.com/jcmuller/version_manager_shell_string_old/rubychecker"
	"github.com/jcmuller/version_manager_shell_string_old/nvmchecker"
	"github.com/jcmuller/version_manager_shell_string_old/gochecker"
)

var (
	path = ""
	output string
)

func newMain() {
	v := versions.New(config.New(), path)
	v.Check()
	//fmt.Println(v)
	output = v.String()
}

func oldMain() {
	v := versionsOld.New(path)

	rc, err := rubychecker.New(path)
	if err == nil {
		v.AddChecker(rc)
	}

	nc, err := nvmchecker.New(path)
	if err == nil {
		v.AddChecker(nc)
	}

	gc, err := gochecker.New(path)
	if err == nil {
		v.AddChecker(gc)
	}

	v.GetVersions()
	//fmt.Println(v)
	output = v.String()
}

func BenchmarkOld(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		oldMain()
	}
}

func BenchmarkNew(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		newMain()
	}
}
