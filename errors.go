/*
Copyright (C) 2025 github.com/go-schwab

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see
<https://www.gnu.org/licenses/>.
*/

package trader

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// credit for most of this goes to @jazzboME

var ErrNeedReAuthorization = errors.New("need to reinitalize or account not available to caller")
var ErrValidation = errors.New("validation error - non fatal from Schwab")
var ErrForbidden = errors.New("url is forbidden to client")
var ErrNotFound = errors.New("url not found")
var ErrUnexpectedServer = errors.New("server is freaking out")
var ErrTemporaryServer = errors.New("server is taking a tylenol, brb")

// Custom Error Struct
type TraderError struct {
	Inner    error // initial error
	Response *http.Response
}

// Needs Error() to satisfy error interface
func (e *TraderError) Error() string {
	return fmt.Sprintf("%v: %v", e.Inner, e.Inner.Error())
}

// Unwrap is needed to support working with errors.Is & errors.As.
func (e *TraderError) Unwrap() error {
	// Return the inner error.
	return e.Inner
}

// WrapTraderError to easily create a new error which wraps the given error.
func WrapTraderError(err error, resp *http.Response) error {
	return &TraderError{
		Response: resp,
		Inner:    err,
	}
}

func GetMessage(e interface{}) string {
	body, err := io.ReadAll(e.(*TraderError).Response.Body)
	isErrNil(err)
	return string(body)
}

func GetStatusCode(e interface{}) int {
	return e.(*TraderError).Response.StatusCode
}

// is the err nil?
func isErrNil(err error) {
	if err != nil {
		log.Fatalf("[fatal] %s", err.Error())
	}
}
