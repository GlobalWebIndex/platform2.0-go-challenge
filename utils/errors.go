package utils

import "errors"

var (
	// ErrAlreadyExists custom error message for already existent resources
	ErrAlreadyExists = errors.New("resource already exists")
	// ErrNotFound custom error message for non existent resources
	ErrNotFound = errors.New("resource not found")
	// ErrInvalidAsset custom error message for Assets, with more than one of Chart, Insight and Audience
	ErrInvalidAsset = errors.New("invalid asset")
)

// NewError is a helper function for specific Echo error responses
func NewError(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}
