package filelib

import (
	"os"
	// "bufio"
)

// Wrties to file
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
	// Might not need synchronization here either,
	// but in case we need to write from multiple goroutines
	// might as well
	lock <- bytesWritten
}
