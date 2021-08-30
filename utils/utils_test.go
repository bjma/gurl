package utils

import (
	"fmt"
	"path/filepath"
	"testing"
)

// Tests if file parser can create a file in current directory
func TestParseFileBasic(t *testing.T) {
	var tests = []struct {
		path string
		want string
	}{
		{"../file.txt", filepath.Join(basepath, "file.txt")},
		{"../tmp/file.txt", filepath.Join(basepath, "tmp/file.txt")},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s\n", test.path)
		t.Run(testname, func(t *testing.T) {
			ans := ParseFile(test.path)
			if ans != test.want {
				t.Errorf("ParseFile: Got %s, want %s\n", ans, test.want)
			}
		})
	}
}

// Tests ParseFile on nested relative paths and home path
func TestParseFileDriven(t *testing.T) {
	var tests = []struct {
		path string
		want string
	}{
		{"../../file.txt", filepath.Join(filepath.Join(basepath, "../"), "file.txt")},
		{"~/bee-was-here.txt", filepath.Join("/Users/bjma", "bee-was-here.txt")},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s\n", test.path)
		t.Run(testname, func(t *testing.T) {
			ans := ParseFile(test.path)
			if ans != test.want {
				t.Errorf("ParseFile: Got %s, want %s\n", ans, test.want)
			}
		})
	}
}