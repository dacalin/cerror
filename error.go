package cerror

type Error interface {
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Extra Message
	ExtraMsg() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error

	// Add origin error and return a new Error object
	AddOrigError(origError error) Error

	// Add and extra message to Error
	AddExtraMsg(extraMsg string) Error
}
