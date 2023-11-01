package cerror

import "fmt"

type BaseError struct {
	// Classification of error
	code string

	// Detailed information about error
	message string

	// extra information about error
	extra string

	// Optional original error this error is based off of. Allows building
	// chained errors.
	errs []error
}

func newBaseError(code, message string, extra string, origErrs []error) *BaseError {
	b := &BaseError{
		code:    code,
		message: message,
		extra:   extra,
		errs:    origErrs,
	}

	return b
}

func New(code, message string) Error {
	var errs []error

	return newBaseError(code, message, "", errs)
}

// Satisfies the error interface.
func (self BaseError) Error() string {
	size := len(self.errs)
	if size > 0 {
		return SprintError(self.code, self.message, self.extra, errorList(self.errs))
	}

	return SprintError(self.code, self.message, self.extra, nil)
}

// String returns the string representation of the error.
// Alias for Error to satisfy the stringer interface.
func (self BaseError) String() string {
	return self.Error()
}

// Code returns the short phrase depicting the classification of the error.
func (self BaseError) Code() string {
	return self.code
}

// Message returns the error details message.
func (self BaseError) Message() string {
	return self.message
}

// Message returns the error details message.
func (self BaseError) ExtraMsg() string {
	return self.extra
}

func (self BaseError) AddOrigError(origError error) Error {
	errs := append(self.errs, origError)
	newError := newBaseError(self.code, self.message, self.extra, errs)
	return newError
}

func (self BaseError) AddExtraMsg(extraMsg string) Error {
	extra := self.extra + extraMsg + " "
	newError := newBaseError(self.code, self.message, extra, self.errs)
	return newError
}

// OrigErr returns the original error if one was set. Nil is returned if no
// error was set. This only returns the first element in the list. If the full
// list is needed, use BatchedErrors.
func (self BaseError) OrigErr() error {
	switch len(self.errs) {
	case 0:
		return nil
	case 1:
		return self.errs[0]
	default:
		if err, ok := self.errs[0].(Error); ok {
			return newBatchError(err.Code(), err.Message(), err.ExtraMsg(), self.errs[1:])
		}
		return newBatchError("BatchedErrors", "multiple errors occurred", self.extra, self.errs)
	}
}

// OrigErrs returns the original errors if one was set. An empty slice is
// returned if no error was set.
func (self BaseError) OrigErrs() []error {
	return self.errs
}

// SprintError returns a string of the formatted error code.
//
// Both extra and origErr are optional.  If they are included their lines
// will be added, but if they are not included their lines will be ignored.
func SprintError(code, message, extra string, origErr error) string {
	msg := fmt.Sprintf("[%s] %s", code, message)
	if extra != "" {
		msg = fmt.Sprintf("%s\t%s", msg, extra)
	}
	if origErr != nil {
		msg = fmt.Sprintf("%s\t %s", msg, origErr.Error())
	}
	return msg
}
