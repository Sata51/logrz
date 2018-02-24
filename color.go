package gologrz

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
	disableColor bool
)

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
