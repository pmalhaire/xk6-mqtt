package mqtt

import (
	"errors"
	"fmt"
)

var (
	ErrorNilState              = errors.New("State is nil")
	ErrorNilWriter             = errors.New("Writer is nil")
	ErrorNilReader             = errors.New("Reader is nil")
	ErrorWriterTimeout         = errors.New("Writer timeout")
	ErrorReaderTimeout         = errors.New("Reader timeout")
	ErrorMessageRecieveTimeout = errors.New("Reader message recieve timeout")
)

func ReportError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, ":", err)
	} else {
		fmt.Println(msg)
	}
}
