package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
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

type RestErr interface {
	Status() int
	Error() string
}

type RestError struct {
	ErrStatus  int       `json:"status,omitempty"`
	ErrMessage string    `json:"error,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d, error: %s", e.ErrStatus, e.ErrMessage)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func NewRestError(status int, err string) RestErr {
	return RestError{
		ErrStatus:  status,
		ErrMessage: err,
		Timestamp:  time.Now().UTC(),
	}
}

func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("field '%s' must be not empty", strings.ToLower(err.Field()))
	case "email":
		return fmt.Sprintf("invalid email")
	case "url":
		return fmt.Sprintf("field '%s' must be URL", strings.ToLower(err.Field()))
	case "oneof":
		return fmt.Sprintf("field '%s' must be enum type", err.Field())
	default:
		return "json validation error"
	}
}

func ParseError(err error) RestErr {
	var validationErrors validator.ValidationErrors

	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeout)
	case errors.Is(err, UnauthorizedError):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized)
	case errors.Is(err, BadRequestError):
		return NewRestError(http.StatusBadRequest, ErrBadRequest)
	case errors.Is(err, ForbiddenError):
		return NewRestError(http.StatusForbidden, ErrForbidden)
	case errors.Is(err, errs.ErrNotExist):
		return NewRestError(http.StatusNotFound, ErrNotFound)
	case errors.Is(err, errs.ErrInvalidToken):
		return NewRestError(http.StatusUnauthorized, err.Error())
	case errors.As(err, &validationErrors):
		return NewRestError(http.StatusBadRequest, getValidationMessage(validationErrors[0]))
	default:
		if code, ok := errorStatusMap[err]; ok {
			return NewRestError(code, err.Error())
		}
		if restErr, ok := err.(*RestError); ok {
			return restErr
		}
		return NewRestError(http.StatusInternalServerError, ErrInternalServerError)
	}
}

func (h *Handler) errorResponse(context *gin.Context, err error) {
	h.logger.Error(err.Error())
	restErr := ParseError(err)
	context.AbortWithStatusJSON(restErr.Status(), restErr)
}

func (h *Handler) successResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, data)
}

func (h *Handler) createdResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusCreated, data)
}
