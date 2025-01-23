package errs

import "errors"

var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrOCITypeMismatch = errors.New("mismatch on OCI artifact type")
)
