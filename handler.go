package logrz

import "runtime"

// F generate a fatal log format
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

// P generate a panic log format
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

// E generate an error log format
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

// W generate a warn log format
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

// I generate an info log format
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

// D generate a debug log format
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

// T generate a trace log format
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
