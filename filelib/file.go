/*
Package filelib implements a library for providing a
variety of functionality for files.

This includes resolving filepaths (relative and absolutes),
reading, and writing files.

Files with the prefix '@' are assumed to be files to be read from,
while all other files are written to.
*/
package filelib

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/bjma/gurl/utils"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Join(filepath.Dir(b), "../")
	cwd, _     = os.Getwd()
	homedir, _ = os.UserHomeDir()
)

// Parses filepath input and returns an absolute path to the file
func ParseFile(file string) string {
	var path string
	// Read file if file has prefix '@', else write file
	if strings.HasPrefix(file, "@") {
		path = ResolvePath(strings.TrimPrefix(file, "@"))
	} else {
		// User root directory, just return
		if strings.HasPrefix(file, "/") {
			return file
		}
		path = ResolvePath(file)
	}
	return path
}

// Resolves a file path by constructing an absolute path to the file,
// creating any directories that don't exist
func ResolvePath(path string) string {
	var absPath string
	if strings.HasPrefix(path, "~") {
		absPath = filepath.Join(homedir, "")
	} else {
		absPath = filepath.Join(cwd, "")
	}
	f, paths := tokenizeFilePath(path)
	for _, p := range paths {
		// Handle relative paths
		if match, _ := regexp.MatchString("^(.|~)+", p); match {
			if strings.HasPrefix(p, "~") {
				continue
			}
			p += "/"
		}
		absPath = filepath.Join(absPath, p)
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			os.Mkdir(absPath, 0755)
		}
	}
	absPath = filepath.Join(absPath, f)
	return absPath
}

// Returns the file extension of a filename
func GetFileExtension(file string) string {
	// Reverse string so that we can get LAST occurence of `.` (handles cases like `foo.bar.txt`)
	f := utils.ReverseString(file)
	i := strings.Index(f, ".")
	if i < 0 {
		return file
	} else if i == 0 {
		panic("gurl: Failed writing body to file with filename " + file)
	}
	return file[len(file)-i:]
}

// Tokenizes file path and returns (filename, [path, to, file])
func tokenizeFilePath(path string) (string, []string) {
	tokens := strings.Split(path, "/")
	n := len(tokens)
	return tokens[n-1], tokens[:n-1]
}
