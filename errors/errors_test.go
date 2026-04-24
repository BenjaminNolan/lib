package errors

import (
	goerrors "errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	ErrExpectedNonNilResult Error = "expected non-nil result"
)

func TestError(t *testing.T) {
	const msg = "error"

	err := New(msg)

	require.NotNil(t, err, ErrExpectedNonNilResult.Error())
	assert.Equalf(t, msg, err.Error(), "expected '%s', got '%s'", msg, err)
}

func TestJoinSimple(t *testing.T) {
	err1Msg := "error 1"
	err2Msg := "error 2"

	err1 := New(err1Msg)
	err2 := New(err2Msg)

	require.NotNil(t, err1)
	require.NotNil(t, err2)

	err := Join(err1, err2)

	require.NotNil(t, err)

	errSlice, ok := err.(ErrorSlice)

	require.True(t, ok)
	require.NotNil(t, errSlice)

	errs := errSlice.Unwrap()

	require.NotNil(t, errs)
	require.Len(t, errs, 2)
	assert.Equal(t, errs[0].Error(), err1Msg)
	assert.Equal(t, errs[1].Error(), err2Msg)
}

func TestJoinWithNil(t *testing.T) {
	err1Msg := "error 1"
	err2Msg := "error 2"

	err1 := New(err1Msg)
	err2 := New(err2Msg)
	var err3 error

	require.NotNil(t, err1)
	require.NotNil(t, err2)
	require.Nil(t, err3)

	err := Join(err1, err2, err3)

	require.NotNil(t, err)

	errSlice, ok := err.(ErrorSlice)

	require.True(t, ok)
	require.NotNil(t, errSlice)

	errs := errSlice.Unwrap()

	require.NotNil(t, errs)
	require.Len(t, errs, 2)
	assert.Equal(t, errs[0].Error(), err1Msg)
	assert.Equal(t, errs[1].Error(), err2Msg)
}

func TestJoinNested(t *testing.T) {
	err1Msg := "error 1"
	err2Msg := "error 2"
	err3Msg := "error 3"
	err4Msg := "error 4"
	err5Msg := "error 5"

	err1 := New(err1Msg)
	err2 := New(err2Msg)
	err3 := New(err3Msg)
	err4 := New(err4Msg)
	err5 := New(err5Msg)

	require.NotNil(t, err1)
	require.NotNil(t, err2)
	require.NotNil(t, err3)
	require.NotNil(t, err4)
	require.NotNil(t, err5)

	err := Join(Join(err1, err2), err3, Join(err4, err5))

	require.NotNil(t, err)

	errSlice, ok := err.(ErrorSlice)

	require.True(t, ok)
	require.NotNil(t, errSlice)

	errs := errSlice.Unwrap()

	require.NotNil(t, errs)
	require.Len(t, errs, 5)
	assert.Equal(t, errs[0].Error(), err1Msg)
	assert.Equal(t, errs[1].Error(), err2Msg)
	assert.Equal(t, errs[2].Error(), err3Msg)
	assert.Equal(t, errs[3].Error(), err4Msg)
	assert.Equal(t, errs[4].Error(), err5Msg)
}

func TestIs(t *testing.T) {
	const err1 Error = "error 1"
	require.NotNil(t, err1)

	err2 := goerrors.New("error 2")
	require.NotNil(t, err2)

	errJoined := Join(err1, err2)
	require.NotNil(t, errJoined)

	assert.ErrorIs(t, errJoined, err1)
	assert.ErrorIs(t, errJoined, err2)

	assert.True(t, Is(errJoined, err1))
	assert.True(t, Is(errJoined, err2))
}

func TestIsNested(t *testing.T) {
	const err1 Error = "error 1"
	require.NotNil(t, err1)

	err2 := goerrors.New("error 2")
	require.NotNil(t, err2)

	err3 := goerrors.New("error 3")
	require.NotNil(t, err3)

	err4 := goerrors.New("error 4")
	require.NotNil(t, err4)

	err5 := goerrors.New("error 5")
	require.NotNil(t, err5)

	errJoined := Join(err1, Join(Join(err2, err3), err4), err5)
	require.NotNil(t, errJoined)

	assert.ErrorIs(t, errJoined, err1)
	assert.ErrorIs(t, errJoined, err2)
	assert.ErrorIs(t, errJoined, err3)
	assert.ErrorIs(t, errJoined, err4)
	assert.ErrorIs(t, errJoined, err5)

	assert.True(t, Is(errJoined, err1))
	assert.True(t, Is(errJoined, err2))
	assert.True(t, Is(errJoined, err3))
	assert.True(t, Is(errJoined, err4))
	assert.True(t, Is(errJoined, err5))
}

func TestAs(t *testing.T) {
	var err1type Error
	require.NotNil(t, err1type)

	err1 := New("error 1")
	require.NotNil(t, err1)

	err2 := goerrors.New("error 2")
	require.NotNil(t, err2)

	errJoined := Join(err1, err2)
	require.NotNil(t, errJoined)

	assert.ErrorAs(t, errJoined, &err1type)
	assert.True(t, As(errJoined, &err1type))
}

func TestAsNested(t *testing.T) {
	var err3type Error
	require.NotNil(t, err3type)

	err1 := goerrors.New("error 1")
	require.NotNil(t, err1)

	err2 := goerrors.New("error 2")
	require.NotNil(t, err2)

	const err3 Error = "error 3"
	require.NotNil(t, err3)

	err4 := goerrors.New("error 4")
	require.NotNil(t, err4)

	err5 := goerrors.New("error 5")
	require.NotNil(t, err5)

	errJoined := Join(err1, Join(Join(err2, err3), err4), err5)
	require.NotNil(t, errJoined)

	assert.ErrorAs(t, errJoined, &err3type)
	assert.True(t, As(errJoined, &err3type))
}

func TestErrorUnwrap(t *testing.T) {
	const err1 Error = "error 1"
	require.NotNil(t, err1)

	const err2 Error = "error 2"
	require.NotNil(t, err2)

	fmtErr := fmt.Errorf("std error: %w", fmt.Errorf("wrapped error: %w", err1))
	require.NotNil(t, fmtErr)

	joinedErrs := Join(err2, fmtErr)
	require.NotNil(t, joinedErrs)

	joinedErrsSlice, ok := joinedErrs.(ErrorSlice)
	require.True(t, ok)
	require.NotNil(t, joinedErrsSlice)

	unwrappedErr := joinedErrsSlice.Unwrap()
	require.NotNil(t, unwrappedErr)
	require.Len(t, unwrappedErr, 2)
	assert.Equal(t, unwrappedErr[0].Error(), err2.Error())
	assert.Equal(t, unwrappedErr[1].Error(), fmtErr.Error())
}

func TestUnwrap(t *testing.T) {
	const err1 Error = "error 1"
	require.NotNil(t, err1)

	wrappedErr := fmt.Errorf("wrapped err: %w", err1)
	require.NotNil(t, wrappedErr)

	unwrappedErr := Unwrap(wrappedErr)
	require.NotNil(t, unwrappedErr)
	assert.ErrorIs(t, unwrappedErr, err1)
	assert.True(t, Is(unwrappedErr, err1))
}
