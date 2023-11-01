package cerror

type BatchedErrors interface {
	// Satisfy the base Error interface.
	Error

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

func newBatchError(code, message string, extra string, errs []error) BatchedErrors {
	return newBaseError(code, message, extra, errs)
}
