package faults

import flt "github.com/PlayerR9/go-fault"

type StdFaultCode int

const (
	Unknown  StdFaultCode = iota // UNKNOWN
	FatalErr                     // FATAL
)

func NewFatalErr(msg string) flt.Fault {
	return flt.New(FatalErr, msg)
}
