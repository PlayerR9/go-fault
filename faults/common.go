package faults

import (
	flt "github.com/PlayerR9/go-fault"
)

type StdFaultCode int

const (
	Unknown       StdFaultCode = iota - 1 // UNKNOWN
	FatalErr                              // FATAL
	OperationFail                         // Operation Failed
	BadParameter                          // Bad Parameter
)

// NewNilReceiver creates a new fault with the code BadParameter and the
// message "Receiver must not be nil".
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func NewNilReceiver() flt.Fault {
	return flt.New(BadParameter, "Receiver must not be nil")
}

func FromString(s string) flt.Fault {
	if s == "" {
		return nil
	}

	return flt.New(Unknown, s)
}

func FromError(err error) flt.Fault {
	if err == nil {
		return nil
	}

	return flt.New(Unknown, err.Error())
}

func NewErrPanic(value any) flt.Fault {
	fault := flt.New(FatalErr, "A panic has occurred")

	return fault
}