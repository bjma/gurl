package filelib

import (
	"os"

	"github.com/bjma/gurl/handler"
)

// Writes byte array to file located at `path`
// and returns number of bytes written to file.
// If file does not exist, a new file with the filename
// defined by `path` will be created; else, the
// file is truncated.
//
// See: https://gobyexample.com/writing-files
func WriteFile(path string, data []byte) int {
	fd, err := os.Create(path)
	if err != nil {
		handler.HandleError(err)
	}
	defer fd.Close()

	// Write byte data into file
	bytesWritten, err := fd.Write(data)
	if err != nil {
		handler.HandleError(err)
	}
	fd.Sync()
	return bytesWritten
}
