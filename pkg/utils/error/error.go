package error

import (
	"errors"
	"fmt"
	"strings"

	// "github.com/Zomato/go/logger"
	errors2 "github.com/pkg/errors"
)

type FactoryError struct {
	error         error
	errorCode     errorCode
	errorMsg      string
	loggingParams map[string]interface{}
	exists        bool
}

func NewFactoryError(errCode errorCode, error string) FactoryError {
	c := FactoryError{errorCode: errCode, errorMsg: error, exists: true}
	e := errors.New(fmt.Sprintf("Code: %s | %s", c.errorCode, c.errorMsg))
	c.error = errors2.WithStack(e)
	c.loggingParams = make(map[string]interface{}, 0)
	return c
}

func (c FactoryError) Exists() bool {
	return c.exists
}

/* ToDo: add error logging
func (c FactoryError) Log() {
	logger.Errorln(c.ToString())
}
*/

func (c FactoryError) LoggingParams() map[string]interface{} {
	return c.loggingParams
}

func (c FactoryError) ErrorCode() errorCode {
	return c.errorCode
}

func (c FactoryError) ToError() error {
	return c.error
}

func (c FactoryError) Error() string {
	return c.error.Error()
}

func (c FactoryError) ErrorMessage() string {
	return c.errorMsg
}

func (c FactoryError) ToString() string {
	logMsg := fmt.Sprintf("Code: %s, Msg: %s", c.errorCode, c.errorMsg)

	paramStrings := make([]string, 0)
	for key, val := range c.loggingParams {
		paramStrings = append(paramStrings, fmt.Sprintf("%s: {%+v}", strings.ToUpper(key), val))
	}
	return fmt.Sprintf("%s, Params: [%+v]", logMsg, strings.Join(paramStrings, " | "))
}

// WithParam value should not be a pointer
func (c FactoryError) WithParam(key string, val interface{}) FactoryError {
	if c.loggingParams == nil {
		c.loggingParams = make(map[string]interface{}, 0)
	}
	c.loggingParams[key] = val
	return c
}

func (c FactoryError) ErrorString() string {
	return c.errorMsg
}

func (c FactoryError) UserMessage() string {
	if val, found := errorCodeToUserMessages[c.errorCode]; found {
		return val
	}
	return ""
}
