package utility

// PatchError - a custom error declration
type PatchError struct {
	s   string
	err error
}

// NewPatchError returns the new path object
func NewPatchError(s string, err error) *PatchError {
	return &PatchError{s, err}
}

// Error returns the string format of error
func (e *PatchError) Error() (out string) {
	if e.err != nil {
		return e.s + " : " + e.err.Error()
	}
	return
	// return e.s + fmt.Sprintf(": %v", e.err)
}
