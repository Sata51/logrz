package logrz

import "runtime"

// F generate a fatal log format
func F(format string, a ...interface{}) {
	if printLevel >= FatalLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(FatalLevel, out, details, format, a...)
		} else {
			log(FatalLevel, out, nil, format, a...)
		}
	}
}

// P generate a panic log format
func P(format string, a ...interface{}) {
	if printLevel >= PanicLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(PanicLevel, out, details, format, a...)
		} else {
			log(PanicLevel, out, nil, format, a...)
		}
	}
}

// E generate an error log format
func E(format string, a ...interface{}) {
	if printLevel >= ErrorLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(ErrorLevel, out, details, format, a...)
		} else {
			log(ErrorLevel, out, nil, format, a...)
		}
	}
}

// W generate a warn log format
func W(format string, a ...interface{}) {
	if printLevel >= WarnLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(WarnLevel, out, details, format, a...)
		} else {
			log(WarnLevel, out, nil, format, a...)
		}
	}
}

// I generate an info log format
func I(format string, a ...interface{}) {
	if printLevel >= InfoLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(InfoLevel, out, details, format, a...)
		} else {
			log(InfoLevel, out, nil, format, a...)
		}
	}
}

// D generate a debug log format
func D(format string, a ...interface{}) {
	if printLevel >= DebugLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(DebugLevel, out, details, format, a...)
		} else {
			log(DebugLevel, out, nil, format, a...)
		}
	}
}

// T generate a trace log format
func T(format string, a ...interface{}) {
	if printLevel >= TraceLevel {
		out := true
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log(TraceLevel, out, details, format, a...)
		} else {
			log(TraceLevel, out, nil, format, a...)
		}
	}
}
