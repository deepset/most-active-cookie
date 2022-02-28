package errors

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrEmptyList          = errors.New("empty data list")
	ErrEmptyStack         = errors.New("empty stack")
	ErrFullStack          = errors.New("full stack")
	ErrFileNotFound       = errors.New("file not found")
	ErrDateNotFound       = errors.New("date not found in list")
	ErrInvalidFile        = errors.New("invalid csv file")
	ErrDataFormatMismatch = errors.New("data record and header format mismatch")
)

func ErrorCheck(e error) {
	if e != nil {
		ExitGracefully(e)
	}
}

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
