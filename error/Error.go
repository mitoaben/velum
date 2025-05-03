package error

import (
	"fmt"
	"os"
)

type Error struct {
	Message string
	Line    uint
	Column  uint
	Modul   string
}

func (e Error) Error() string {
	return fmt.Sprintf("Error: (%s) %s at line %d, column %d\n", e.Modul, e.Message, e.Line, e.Column)
}

func NewError(message string, line uint, column uint) Error {
	return Error{Message: message, Line: line, Column: column}
}

func HandleError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func (e *Error) ReportError(message string, line uint, column uint) {
	e.Column = column
	e.Line = line
	e.Message = message
	HandleError(e)
}
