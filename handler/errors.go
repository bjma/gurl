package handler

import (
	"errors"
	"fmt"
	"os"
)

// Same thing as `log.Fatalln`
func HandleError(err error) {
	fmt.Printf("gurl: %s\n", err.Error())
	os.Exit(1)
}

func NewError(s string) error {
	return errors.New(s)
}
