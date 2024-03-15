package internal

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorBuilder wraps github.com/pkg/errors and lets the user build an error
// through successive writes, each one wrapping the builder's error written
// error message.
type ErrorBuilder struct {
	err error
	msg string
}

// Write wraps the error builder's error with the arguments.
func (b *ErrorBuilder) Write(format string, args ...any) {
	var msg string
	if b.msg != "" {
		msg += b.msg + ": "
	}
	msg += fmt.Sprintf(format, args...)
	if b.err == nil {
		b.err = errors.New(msg)
		return
	}
	b.err = errors.Wrap(b.err, msg)
}

// SetMessage sets a string that will be prepended to all successive
// calls to Write.
func (b *ErrorBuilder) SetMessage(format string, args ...any) {
	b.msg = fmt.Sprintf(format, args...)
}

// UnsetMessage resets the prepended string.
func (b *ErrorBuilder) UnsetMessage() {
	b.msg = ""
}

// Error return the error builder's internal error.
func (b *ErrorBuilder) Error() error {
	return b.err
}

// String returns the error builder's internal error string.
func (b *ErrorBuilder) String() string {
	return b.err.Error()
}
