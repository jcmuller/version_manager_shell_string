package main

import (
	"testing"
	"os"
)

func TestFindDirectory(t *testing.T) {
	os.MkdirAll("/tmp/go/test/.git", 0755)
	os.MkdirAll("/tmp/go/test/foo/bar", 0755)

	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected string
		wantErr  bool
	}{
		{
			name:     "Empty directory",
			args:     args{"foo"},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Good directory",
			args:     args{"/tmp/go/test/foo/bar"},
			expected: "/tmp/go/test",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := findDirectory(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if actual != tt.expected {
				t.Errorf("findDirectory() = %v, expected %v", actual, tt.expected)
			}
		})
	}

	os.RemoveAll("/tmp/go/test")
}
