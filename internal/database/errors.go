package database

// Level is a logging priority. Higher levels are more important.
type Level uint32

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel = iota
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

// ErrNotFound is returned when no rows are found.
type ErrNotFound string

// Error method to implement error interface.
func (e ErrNotFound) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrNotFound) Level() Level {
	return WarnLevel
}

// ErrConflict is returned when a conflict or duplicate entry is found.
type ErrConflict string

// Error method to implement error interface.
func (e ErrConflict) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrConflict) Level() Level {
	return ErrorLevel
}
