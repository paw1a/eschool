package console

import (
	"fmt"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
	"net/http"
)

const (
	ErrBadRequest          = "bad request"
	ErrNotFound            = "not found"
	ErrUnauthorized        = "unauthorized"
	ErrForbidden           = "forbidden"
	ErrInternalServerError = "internal server error"
	ErrRequestTimeout      = "request timeout"
)

var (
	BadRequestError          = errors.New("bad request")
	NotFoundError            = errors.New("not Found")
	UnauthorizedError        = errors.New("unauthorized")
	ForbiddenError           = errors.New("forbidden")
	InternalServerError      = errors.New("internal server error")
	PathIdParamIsEmptyError  = errors.New("empty id query parameter")
	PathIdParamIsInvalidUUID = errors.New("id query parameter is not uuid")
)

var errorStatusMap = map[error]int{
	errs.ErrCourseNotEnoughLessons:               http.StatusBadRequest,
	errs.ErrCourseLessonInvalidScore:             http.StatusBadRequest,
	errs.ErrCoursePracticeLessonEmptyTests:       http.StatusBadRequest,
	errs.ErrCoursePracticeLessonEmptyTestTaskUrl: http.StatusBadRequest,
	errs.ErrCoursePracticeLessonEmptyTestOptions: http.StatusBadRequest,
	errs.ErrCoursePracticeLessonInvalidTestScore: http.StatusBadRequest,
	errs.ErrCoursePracticeLessonInvalidTestLevel: http.StatusBadRequest,
	errs.ErrCourseTheoryLessonEmptyUrl:           http.StatusBadRequest,
	errs.ErrCourseVideoLessonEmptyUrl:            http.StatusBadRequest,
	errs.ErrCourseReadyState:                     http.StatusBadRequest,
	errs.ErrCoursePublishedState:                 http.StatusBadRequest,
	errs.ErrCourseInvalidLevel:                   http.StatusBadRequest,
	errs.ErrCourseInvalidPrice:                   http.StatusBadRequest,
	errs.ErrCertificateCourseNotPassed:           http.StatusBadRequest,
	errs.ErrFilenameEmpty:                        http.StatusBadRequest,
	errs.ErrFilepathEmpty:                        http.StatusBadRequest,
	errs.ErrFileReaderEmpty:                      http.StatusBadRequest,
	errs.ErrSaveFileError:                        http.StatusBadRequest,
	errs.ErrUserIsNotSchoolTeacher:               http.StatusBadRequest,
	errs.ErrUserIsAlreadyCourseStudent:           http.StatusConflict,
	errs.ErrInvalidPaymentSum:                    http.StatusBadRequest,
	errs.ErrDecodePaymentKeyFailed:               http.StatusBadRequest,

	errs.ErrDuplicate:         http.StatusBadRequest,
	errs.ErrNotExist:          http.StatusNotFound,
	errs.ErrUpdateFailed:      http.StatusInternalServerError,
	errs.ErrDeleteFailed:      http.StatusInternalServerError,
	errs.ErrPersistenceFailed: http.StatusInternalServerError,
	errs.ErrEnumValueError:    http.StatusInternalServerError,
	errs.ErrTransactionError:  http.StatusInternalServerError,

	errs.ErrNotUniqueEmail:          http.StatusConflict,
	errs.ErrInvalidCredentials:      http.StatusUnauthorized,
	errs.ErrAuthSessionIsNotPresent: http.StatusUnauthorized,
	errs.ErrInvalidTokenSignMethod:  http.StatusUnauthorized,
	errs.ErrInvalidTokenClaims:      http.StatusUnauthorized,
	errs.ErrInvalidFingerprint:      http.StatusUnauthorized,

	PathIdParamIsEmptyError:  http.StatusBadRequest,
	PathIdParamIsInvalidUUID: http.StatusBadRequest,
}

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
	switch {
	//case errors.Is(err, context.DeadlineExceeded):
	//	return NewConsoleError(http.StatusRequestTimeout, ErrRequestTimeout)
	//case errors.Is(err, UnauthorizedError):
	//	return NewConsoleError(http.StatusUnauthorized, ErrUnauthorized)
	//case errors.Is(err, BadRequestError):
	//	return NewConsoleError(http.StatusBadRequest, ErrBadRequest)
	//case errors.Is(err, ForbiddenError):
	//	return NewConsoleError(http.StatusForbidden, ErrForbidden)
	//case errors.Is(err, errs.ErrNotExist):
	//	return NewConsoleError(http.StatusNotFound, ErrNotFound)
	//case errors.Is(err, errs.ErrInvalidToken):
	//	return NewConsoleError(http.StatusUnauthorized, err.Error())
	//case errors.As(err, &validationErrors):
	//	return NewConsoleError(http.StatusBadRequest, getValidationMessage(validationErrors[0]))
	//default:
	//	if code, ok := errorStatusMap[err]; ok {
	//		return NewConsoleError(code, err.Error())
	//	}
	//	if restErr, ok := err.(*ErrorConsole); ok {
	//		return restErr
	//	}
	//	return NewConsoleError(http.StatusInternalServerError, ErrInternalServerError)
	}
	return NewConsoleError(err.Error())
}

func ErrorResponse(err error) {
	consoleErr := ParseError(err)
	fmt.Printf("%v\n", consoleErr)
}
