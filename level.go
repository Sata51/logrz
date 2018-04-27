package logrz

import (
	"fmt"
	"strings"
)

// Level type
type Level uint32

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case PanicLevel:
		return "PANIC"
	case TraceLevel:
		return "TRACE"
	}

	return "unknown"
}

// ParseLevel takes a string level and returns the logrz log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToUpper(lvl) {
	case "PANIC":
		return PanicLevel, nil
	case "FATAL":
		return FatalLevel, nil
	case "ERROR":
		return ErrorLevel, nil
	case "WARN", "WARNING":
		return WarnLevel, nil
	case "INFO":
		return InfoLevel, nil
	case "DEBUG":
		return DebugLevel, nil
	case "TRACE":
		return TraceLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid logrz Level: %q", lvl)
}

//AllLevels A constant exposing all logging levels
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrz.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota //Handle auto incrementation for interger comparison
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	TraceLevel
)
