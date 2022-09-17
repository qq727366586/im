package retry_error

import "context"

type RetryError struct {
	err error
}
type RetryFunc func(ctx context.Context) error

// Unwrap implements error wrapping.
func (e *RetryError) Unwrap() error {
	return e.err
}

// Error returns the error string.
func (e *RetryError) Error() string {
	if e.err == nil {
		return "retryable: <nil>"
	}
	return "retryable: " + e.err.Error()
}

// retryError marks an error as retryable.
func RetryableError(err error) error {
	if err == nil {
		return nil
	}
	return &RetryError{err}
}
