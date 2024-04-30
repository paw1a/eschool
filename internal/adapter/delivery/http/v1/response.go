package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrInternalServerError = "Internal Server Error"
	ErrRequestTimeout      = "Request Timeout"
	ErrUnmarshal           = "JSON Unmarshal Error"
)

var (
	BadRequestError     = errors.New("Bad request")
	NotFoundError       = errors.New("Not Found")
	UnauthorizedError   = errors.New("Unauthorized")
	ForbiddenError      = errors.New("Forbidden")
	InternalServerError = errors.New("Internal Server Error")
	UnmarshalError      = errors.New("JSON Unmarshal Error")
)

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

func ParseError(err error) RestErr {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeout)
	case errors.Is(err, UnauthorizedError):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized)
	case errors.Is(err, UnmarshalError):
		return NewRestError(http.StatusBadRequest, ErrUnmarshal)
	case errors.Is(err, errs.ErrInvalidToken):
		return NewRestError(http.StatusUnauthorized, errs.ErrInvalidToken.Error())
	case errors.Is(err, errs.ErrInvalidTokenSignMethod):
		return NewRestError(http.StatusUnauthorized, errs.ErrInvalidTokenSignMethod.Error())
	case errors.Is(err, errs.ErrInvalidTokenClaims):
		return NewRestError(http.StatusUnauthorized, errs.ErrInvalidTokenClaims.Error())
	case errors.Is(err, errs.ErrInvalidFingerprint):
		return NewRestError(http.StatusUnauthorized, errs.ErrInvalidFingerprint.Error())
	default:
		if restErr, ok := err.(*RestError); ok {
			return restErr
		}
		return NewRestError(http.StatusInternalServerError, ErrInternalServerError)
	}
}

func ErrorResponse(context *gin.Context, err error) {
	log.Error(err)
	restErr := ParseError(err)
	context.AbortWithStatusJSON(restErr.Status(), restErr)
}

func SuccessResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, data)
}

func CreatedResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusCreated, data)
}
