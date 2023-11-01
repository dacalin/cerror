package cerror

// An error list that satisfies the golang interface
type errorList []error

// Satisfies the error interface.
func (self errorList) Error() string {
	msg := ""
	// How do we want to handle the array size being zero
	if size := len(self); size > 0 {
		for i := 0; i < size; i++ {
			msg += "caused by: " + self[i].Error()
			// We check the next index to see if it is within the slice.
			// If it is, then we append a newline. We do this, because unit tests
			// could be broken with the additional '\n'
			if i+1 < size {
				msg += "\t"
			}
		}
	}
	return msg
}
