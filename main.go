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

func log(ltype Level, details *runtime.Func, format string, a ...interface{}) {
	tTime = time.Since(lastLog)
	lastLog = time.Now()
	dts := ""
	if details != nil {
		dtss := strings.Split(details.Name(), "/")
		dts = fmt.Sprintf("[ %s ]", dtss[len(dtss)-1])
	}
	var lg = new(LogComposition)
	lg.Time = fmt.Sprintf("[ %s ]", padRight(lastLog.Format("15:04:05.000"), " ", 12))
	lg.Level = ltype
	lg.Details = dts
	if len(a) != 0 {
		lg.Format = fmt.Sprintf(format, a...)
	} else {
		lg.Format = format
	}
	lg.Interval = fmt.Sprintf(" +%s", tTime)
	formatter(lg)
}

func formatter(log *LogComposition) {
	str := log.Time
	if forceFullColor {
		str = appendType(str, log.Level)
		str += log.Details + " > " + log.Format + log.Interval
		str = fullColor(str, log.Level)
	} else {
		str = appendTypeWithColor(str, log.Level)
		str += log.Details + " > " + log.Format + log.Interval
	}
	str += "\n"
	fmt.Print(str)
}
