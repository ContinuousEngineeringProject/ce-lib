package error

import "fmt"

type errorCode string

func (e errorCode) ToString() string {
	return fmt.Sprintf("%s", e)
}

// ToDo: amend and add more error codes
const (
	NoCode       errorCode = "NO_CODE"
	ErrorPackage errorCode = "ERROR_PACKAGE"
	BashPackage  errorCode = "BASH_PACKAGE"
)
