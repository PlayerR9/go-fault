package faults

import (
	flt "github.com/PlayerR9/go-fault/OLD"
)

// StdFaultCode is the type of a standard fault code.
type StdFaultCode int

const (
	Unknown    StdFaultCode = iota - 1 // UNKNOWN
	FatalErr                           // FATAL
	AssertFail                         // ASSERT

	OperationFail // Operation Failed
	BadParameter  // Bad Parameter
)

// NewNilReceiver creates a new fault with the code BadParameter and the
// message "Receiver must not be nil".
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func NewNilReceiver() flt.Fault {
	return flt.New(BadParameter, "Receiver must not be nil")
}

// FromString creates a new fault with the code Unknown and the message
// given by the given string.
//
// Parameters:
//   - s: The string.
//
// Returns:
//   - Fault: The new fault. Nil if the given string is empty.
func FromString(s string) flt.Fault {
	if s == "" {
		return nil
	}

	return flt.New(Unknown, s)
}

// FromError creates a new fault with the code Unknown and the message
// given by the given error.
//
// Parameters:
//   - err: The error.
//
// Returns:
//   - Fault: The new fault. Nil if the given error is nil.
func FromError(err error) flt.Fault {
	if err == nil {
		return nil
	}

	return flt.New(Unknown, err.Error())
}

// NewErrPanic creates a new fault with the code FatalErr and the message
// "A panic has occurred".
//
// Returns:
//   - Fault: The new fault.
func NewErrPanic(value any) flt.Fault {
	fault := flt.New(FatalErr, "A panic has occurred")

	return fault
}
