package console

import (
	"fmt"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
)

var (
	BadRequestError   = errors.New("bad request")
	NotFoundError     = errors.New("not found")
	UnauthorizedError = errors.New("unauthorized")
	ForbiddenError    = errors.New("forbidden")
)

type ErrConsole interface {
	Error() string
}

type ErrorConsole struct {
	ErrMessage string
}

func (e ErrorConsole) Error() string {
	return fmt.Sprintf("error: %s", e.ErrMessage)
}

func NewConsoleError(err string) ErrConsole {
	return ErrorConsole{
		ErrMessage: err,
	}
}

func ParseError(err error) ErrConsole {
	if errors.Is(err, errs.ErrNotExist) {
		return NewConsoleError(NotFoundError.Error())
	}
	return NewConsoleError(err.Error())
}

func ErrorResponse(err error) {
	consoleErr := ParseError(err)
	fmt.Printf("%v\n", consoleErr)
}
