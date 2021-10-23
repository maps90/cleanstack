package mysql

import "fmt"

const (
	ErrorCredential uint32 = 100 << iota
	ErrorConnection
)

func NewSQLError(errorFlags uint32, errorString ...interface{}) *SQLError {
	return &SQLError{
		text:   fmt.Sprint(errorString...),
		Errors: errorFlags,
	}
}

type SQLError struct {
	Errors uint32
	text   string
}

func (e SQLError) Error() string {
	return fmt.Sprintf("SQLERROR %v : %v", e.Errors, e.text)
}
