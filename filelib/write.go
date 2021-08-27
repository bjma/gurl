package filelib

import (
	"os"
	// "bufio"
)

// Wrties to file
// https://gobyexample.com/writing-files
func WriteFile(filePath string, data []byte) int {
	fd, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// Write byte data into file
	bytesWritten, err := fd.Write(data)
	if err != nil {
		panic(err)
	}
	fd.Sync()

	return bytesWritten
}
