package util

import (
	"os"
	"path/filepath"
	"strings"
)

// Parses file input by tokenizing identifier prefix
// and prepends current working directory,
// returning the aboslute file path.
func ParseFile(fileName string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Create /tmp if directory doesn't already exist
	outPath := filepath.Join(path, "tmp")
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		os.Mkdir(outPath, 0755)
	}
	return outPath + "/" + strings.TrimPrefix(fileName, "@")
}

// Parses an absolute filepath as string
// and returns the resolved filename;
// tbh might be better named as GetFileFromPath or TokenizePath idk
func ParsePath(filePath string) string {
	token := strings.Split(filePath, "/")
	return token[len(token)-1]
}
