package logrz

import (
	"fmt"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 36
	gray    = 37
)

var (
	disableColor   bool
	forceFullColor bool
)

func appendType(str string, level Level) string {
	var levelText string
	levelText = "[ " + padRight(level.String(), " ", 5) + " ]"
	str += levelText
	return str
}

func appendTypeWithColor(str string, level Level) string {
	var levelText string
	levelText = "[ " + padRight(level.String(), " ", 5) + " ]"
	if disableColor {
		str += levelText
	} else {
		var levelColor int
		switch level {
		case DebugLevel:
			levelColor = gray
		case WarnLevel:
			levelColor = yellow
		case ErrorLevel, FatalLevel, PanicLevel:
			levelColor = red
		default:
			levelColor = blue
		}

		str += fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, levelText)
	}
	return str
}

func fullColor(str string, level Level) string {
	var levelColor int
	switch level {
	case DebugLevel:
		levelColor = gray
	case WarnLevel:
		levelColor = yellow
	case ErrorLevel, FatalLevel, PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, str)
}
