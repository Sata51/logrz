package logrz

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

var (
	lastLog    time.Time
	tTime      time.Duration
	printLevel Level
)

func init() {
	lastLog = time.Now()
	printLevel = DebugLevel
	disableColor = false
}

//F generate a fatal log format
func F(format string, a ...interface{}) {
	if printLevel >= FatalLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(FatalLevel, details, format, a...)
		} else {
			log(FatalLevel, nil, format, a...)
		}
	}
}

//P generate a panic log format
func P(format string, a ...interface{}) {
	if printLevel >= PanicLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(PanicLevel, details, format, a...)
		} else {
			log(PanicLevel, nil, format, a...)
		}
	}
}

//E generate an error log format
func E(format string, a ...interface{}) {
	if printLevel >= ErrorLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(ErrorLevel, details, format, a...)
		} else {
			log(ErrorLevel, nil, format, a...)
		}
	}
}

//W generate a warn log format
func W(format string, a ...interface{}) {
	if printLevel >= WarnLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(WarnLevel, details, format, a...)
		} else {
			log(WarnLevel, nil, format, a...)
		}
	}
}

//I generate an info log format
func I(format string, a ...interface{}) {
	if printLevel >= InfoLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(InfoLevel, details, format, a...)
		} else {
			log(InfoLevel, nil, format, a...)
		}
	}
}

//D generate a debug log format
func D(format string, a ...interface{}) {
	if printLevel >= DebugLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(DebugLevel, details, format, a...)
		} else {
			log(DebugLevel, nil, format, a...)
		}
	}
}

//T generate a trace log format
func T(format string, a ...interface{}) {
	if printLevel >= TraceLevel {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(TraceLevel, details, format, a...)
		} else {
			log(TraceLevel, nil, format, a...)
		}
	}
}

func log(ltype Level, details *runtime.Func, format string, a ...interface{}) {
	tTime = time.Since(lastLog)
	lastLog = time.Now()
	var str string
	str = fmt.Sprintf("[ %s ]", padRight(lastLog.Format("15:04:05.999"), " ", 12))
	str = appendTypeWithColor(str, ltype)
	if details != nil {
		dts := strings.Split(details.Name(), "/")
		str += fmt.Sprintf("[ %s ]", dts[len(dts)-1])
	}
	str += " > "
	if len(a) != 0 {
		str += fmt.Sprintf(format, a...)
	} else {
		str += format
	}
	str += fmt.Sprintf(" +%s", tTime)
	str += "\n"
	fmt.Print(str)
}
