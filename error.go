/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import (
	"net/http"

	"github.com/an-repository/errors"
)

type Error struct {
	Status  int
	Message string
}

func NewError(status int, err error) *Error {
	return &Error{
		Status:  status,
		Message: err.Error(),
	}
}

func NewStatusError(status int) *Error {
	return &Error{
		Status:  status,
		Message: http.StatusText(status),
	}
}

func NewMessageError(status int, message string, kv ...any) *Error {
	return NewError(status, errors.New(message, kv...))
}

func NewErrorWithMessage(status int, err error, message string, kv ...any) *Error {
	return NewError(status, errors.WithMessage(err, message, kv...))
}

func (e *Error) Error() string {
	return errors.WithMessage(errors.New(e.Message), "response error", "status", e.Status).Error()
}

/*
######################################################################################################## @(^_^)@ #######
*/
