package filelib

import (
	"os"
	// "bufio"
)

// Writes to file
// https://gobyexample.com/writing-files
func WriteFile(filePath string, data []byte, lock chan int) {
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
	// Might not need synchronization here
	if lock != nil {
		lock <- bytesWritten
	}
}
