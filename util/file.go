package util

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	root       = filepath.Join(basepath, "../")
	homedir, _ = os.UserHomeDir()
)

// Parses filepath input and returns an absolute path to the file
func ParseFile(file string) string {
	var path string
	// Read file if file has prefix '@', else write file
	if strings.HasPrefix(file, "@") {
		path = ResolvePath(strings.TrimPrefix(file, "@"))
	} else {
		// Root directory, just return
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
		absPath = filepath.Join(root, "")
	}
	f, paths := tokenizeFilePath(path)
	for _, p := range paths {
		if match, _ := regexp.MatchString("^(.|~)+", p); match {
			// Already handled case, so skip iteration
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

// Tokenizes file path and returns (filename, [path, to, file])
func tokenizeFilePath(path string) (string, []string) {
	tokens := strings.Split(path, "/")
	n := len(tokens)
	return tokens[n-1], tokens[:n-1]
}
