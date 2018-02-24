package gologrz

import "testing"

func TestLogger(t *testing.T) {
	SetLevel(TraceLevel)
	T("Start!")
	D("Start!")
	I("Start!")
	W("Start!")
	E("Start!")
	F("Start!")
	P("Start!")
}
