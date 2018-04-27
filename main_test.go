package logrz

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	var now = time.Now()
	SetUseFile(true, true, "./file.example")
	SetLevel(TraceLevel)
	T("no output  %s", now.Format(time.RFC1123Z))
	TS("Start!")
	DS("Start!")
	IS("Start!")
	WS("Start!")
	ES("Start!")
	FS("Start!")
	PS("Start!")
	T("will output  %s", now.Format(time.RFC1123Z))
	DisableColor(true)
	T("Start! %s", now.Format(time.RFC1123Z))
	D("Start! %s", now.Format(time.RFC1123Z))
	I("Start! %s", now.Format(time.RFC1123Z))
	W("Start! %s", now.Format(time.RFC1123Z))
	E("Start! %s", now.Format(time.RFC1123Z))
	F("Start! %s", now.Format(time.RFC1123Z))
	P("Start! %s", now.Format(time.RFC1123Z))
	DisableColor(false)
	SetUseFile(true, false, "./file.example")
	T("Start! %s", now.Format(time.RFC1123Z))
	D("Start! %s", now.Format(time.RFC1123Z))
	I("Start! %s", now.Format(time.RFC1123Z))
	W("Start! %s", now.Format(time.RFC1123Z))
	E("Start! %s", now.Format(time.RFC1123Z))
	F("Start! %s", now.Format(time.RFC1123Z))
	P("Start! %s", now.Format(time.RFC1123Z))
	ForceFullColor(true)
	T("Start! %s", now.Format(time.RFC1123Z))
	D("Start! %s", now.Format(time.RFC1123Z))
	I("Start! %s", now.Format(time.RFC1123Z))
	W("Start! %s", now.Format(time.RFC1123Z))
	E("Start! %s", now.Format(time.RFC1123Z))
	F("Start! %s", now.Format(time.RFC1123Z))
	P("Start! %s", now.Format(time.RFC1123Z))
	DisableColor(true)
	ForceFullColor(true)
	T("Start! %s", now.Format(time.RFC1123Z))
	D("Start! %s", now.Format(time.RFC1123Z))
	I("Start! %s", now.Format(time.RFC1123Z))
	W("Start! %s", now.Format(time.RFC1123Z))
	E("Start! %s", now.Format(time.RFC1123Z))
	F("Start! %s", now.Format(time.RFC1123Z))
	P("Start! %s", now.Format(time.RFC1123Z))
	SetUseFile(false, false, "")
}
