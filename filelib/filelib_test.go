package filelib

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/bjma/gurl/utils"
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
		{"~/bri-was-here.txt", filepath.Join(homedir, "bri-was-here.txt")},
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

func TestGetFileExtension(t *testing.T) {
	var tests = []struct {
		str  string
		want string
	}{
		{"foo.txt", "txt"},
		{"foo.bar.txt", "txt"},
		{"foo", "foo"},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s\n", test.str)
		t.Run(testname, func(t *testing.T) {
			ans := GetFileExtension(test.str)
			if ans != test.want {
				t.Errorf("ReverseString: Got %s, want %s\n", ans, test.want)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	var tests = []struct {
		file string
		want string
	}{
		{"../tmp/foo.json", "{\"data\":\"you got pwned :)\"}"},
		{"../tmp/bar.json", "[{\"data\":\"Thing One\"}, {\"data\":\"Thing 2\"}]"},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s\n", test.file)
		t.Run(testname, func(t *testing.T) {
			WriteFile(testname, []byte(test.want))
			ans := ReadFile(testname)
			if string(ans) != test.want {
				t.Errorf("ReadFile: Got %s, want %s\n", ans, test.want)
			}
		})
	}
}

func TestJsonIsArray(t *testing.T) {
    var tests = []struct {
		data string
		want bool
	}{
        {"{\"data\":\"i am an object\"}", false},
		{"[{\"data\":\"i am\"}, {\"data\":\"an array\"}]", true},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s\n", test.data)
		t.Run(testname, func(t *testing.T) {
            d := utils.StrToByteArray(test.data)
            ans := jsonIsArray(d)
            if ans != test.want {
                t.Errorf("jsonIsArray: Got %t, want %t\n", ans, test.want)
            }
		})
	}
}
