package handler

// ErrBind represents when a bind error occurred.
type ErrBind string

// Error method to implement error interface.
func (e ErrBind) Error() string {
	return string(e)
}
