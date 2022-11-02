package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHTTPError(t *testing.T) {
	t.Run("satisfy error interface", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "message", nil)
		assert.Error(t, err)
	})
}

func TestNewBadRequestError(t *testing.T) {
	badRequestError := NewBadRequestError("bad request error", errors.New("bad request error"))
	assert.Equal(t, "bad request error", badRequestError.Message)
	assert.Equal(t, http.StatusBadRequest, badRequestError.Code)
	assert.Error(t, badRequestError.Internal)
}

func TestNewInternalServerError(t *testing.T) {
	internalServerErr := NewInternalServerError(errors.New("cant connect to db"))
	assert.Equal(t, http.StatusInternalServerError, internalServerErr.Code)
	assert.Equal(t, ErrInternalServerError, internalServerErr.Message)
}

func TestNewNotFoundError(t *testing.T) {
	notFound := NewNotFoundError(nil)
	assert.Equal(t, http.StatusNotFound, notFound.Code)
	assert.Equal(t, ErrNotFoundError, notFound.Message)
}

func TestNewUnauthorizedError(t *testing.T) {
	unauthorizedError := NewUnauthorizedError(nil)
	assert.Equal(t, http.StatusUnauthorized, unauthorizedError.Code)
	assert.Equal(t, ErrUnauthorized, unauthorizedError.Message)
}

func TestParseErrorForResponse(t *testing.T) {
	t.Run("returns nil if error is not httpError type", func(t *testing.T) {
		anyError := errors.New("an non http error")
		err := ParseErrorForResponse(anyError)
		assert.Nil(t, err)
	})
	t.Run("returns an http error if so", func(t *testing.T) {
		var expected *httpError
		err := &httpError{
			Code:     http.StatusBadRequest,
			Message:  "message",
			Internal: nil,
		}
		got := ParseErrorForResponse(err)
		assert.ErrorAs(t, got, &expected)
	})
}
