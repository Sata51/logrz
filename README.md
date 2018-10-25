# logrz

An easy way to log data to terminal in go

# Install

```bash
go get github.com/Sata51/logrz
```

# Usage

```go
import (
	"github.com/Sata51/logrz"
)

func main() {
	logrz.T("Start!")
	//[ 14:57:34.183 ][ TRACE ][ logrz.TestLogger ] > Start! +52.391µs
	logrz.D("Start!")
	//[ 14:57:34.184 ][ DEBUG ][ logrz.TestLogger ] > Start! +228.062µs
	logrz.I("Start!")
	//[ 14:57:34.184 ][ INFO  ][ logrz.TestLogger ] > Start! +9.855µs
	logrz.W("Start!")
	//[ 14:57:34.184 ][ WARN  ][ logrz.TestLogger ] > Start! +13.144µs
	logrz.E("Start!")
	//[ 14:57:34.184 ][ ERROR ][ logrz.TestLogger ] > Start! +7.501µs
	logrz.F("Start!")
	//[ 14:57:34.184 ][ FATAL ][ logrz.TestLogger ] > Start! +7.206µs
	logrz.P("Start!")
	//[ 14:57:34.184 ][ PANIC ][ logrz.TestLogger ] > Start! +7.263µs

	//Set minimum log level to be print
	//Log level are sorted from Trace > Debug > Info > Warning > Error > Fatal > Panic
	logrz.SetLevel(logrz.ErrorLevel)

	//Disable log color
	logrs.DisableColor(bool)

	//Force the entire line to be colorized
	logrz.ForceFullColor(bool)
}
```
