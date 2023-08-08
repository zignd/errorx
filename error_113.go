//go:build go1.13
// +build go1.13

package errorx

import "errors"

func isOfType(err error, t *Type) bool {
	e := burrowForTyped(err)
	return e != nil && e.IsOfType(t)
}

func isOfTypeIgnoreTransparent(err error, t *Type) bool {
	e := burrowForTyped(err)
	return e != nil && e.IsOfTypeIgnoreTransparent(t)
}

func (e *Error) isOfType(t *Type) bool {
	cause := e
	for cause != nil {
		if !cause.transparent {
			return cause.errorType.IsOfType(t)
		}

		cause = burrowForTyped(cause.Cause())
	}

	return false
}

func (e *Error) isOfTypeIgnoreTransparent(t *Type) bool {
	return e.errorType.IsOfType(t)
}

// burrowForTyped returns either the first *Error in unwrap chain or nil
func burrowForTyped(err error) *Error {
	raw := err
	for raw != nil {
		typed := Cast(raw)
		if typed != nil {
			return typed
		}

		raw = errors.Unwrap(raw)
	}

	return nil
}
